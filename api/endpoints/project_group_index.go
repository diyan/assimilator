package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	base32num "github.com/Dancapistan/gobase32"
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/diyan/assimilator/tsdb"
	"github.com/diyan/assimilator/web/frontend"
	"github.com/gocraft/dbr"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

const errInvalidStatsPeriod = "Invalid stats_period. Valid choices are '', '24h', and '14d'"

/*
// OrganizationDetails ...
type OrganizationDetails struct {
	models.Organization
	PendingAccessRequests int      `json:"pendingAccessRequests"`
	Features              []string `json:"features"`
	Quota                 Quota    `json:"quota"`
	Access                []string `json:"access"`
	Teams                 []Team   `json:"teams"`
}
*/

// TODO check that TimeSpent is not missing
type Group struct {
	models.Group
	ShortID             string            `json:"shortId"`
	ShareID             string            `json:"shareId"`
	Status              string            `json:"status"`
	Logger              *string           `json:"logger"`
	Level               string            `json:"level"`
	Type                string            `json:"type"`
	Annotations         []string          `json:"annotations"`
	AssignedTo          *string           `json:"assignedTo"`
	Count               int               `json:"count"`
	UserCount           int               `json:"userCount"`
	HasSeen             bool              `json:"hasSeen"`
	Project             GroupProjectInfo  `json:"project"`
	IsBookmarked        bool              `json:"isBookmarked"`
	IsSubscribed        bool              `json:"isSubscribed"`
	Permalink           string            `json:"permalink"`
	Metadata            map[string]string `json:"metadata"`      // TODO check type
	StatusDetails       map[string]string `json:"statusDetails"` // TODO check type
	SubscriptionDetails *string           `json:"subscriptionDetails"`
	Stats               GroupStatistic    `json:"stats"`
}

// GroupStatistic ...
type GroupStatistic struct {
	For24h []string `json:"24h"`
}

// GroupProjectInfo ...
type GroupProjectInfo struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func ProjectGroupIndexGetEndpoint(c echo.Context) error {
	orgSlug := c.Param("organization_slug")
	project := GetProject(c)
	statsPeriod := c.QueryParam("statsPeriod")
	shortIDLookup, _ := strconv.ParseBool(c.QueryParam("shortIdLookup"))
	// TODO return HTTP 400 if shortIdLookup has invalid format
	if !(statsPeriod == "" || statsPeriod == "24h" || statsPeriod == "14d") {
		// TODO introduce better error handling -> return err.InvalidStatsPeriod
		return c.JSON(400, map[string]string{"detail": errInvalidStatsPeriod})
	}
	query := strings.TrimSpace(c.QueryParam("query"))
	db, err := db.FromE(c)
	if err != nil {
		return err
	}
	if query != "" {
		var matchingGroupID int64
		if len(query) == 32 {
			// Check to see if we've got an event ID
			matchingGroupID, _ = db.SelectBySql(`
				select em.group_id
					from sentry_eventmapping em
				where em.project_id = ? and em.event_id = ?`,
				project.ID, query).
				ReturnInt64()
		} else if shortIDLookup && models.LooksLikeShortID(query) {
			// If the query looks like a short id, we want to provide some
			// information about where that is.  Note that this can return
			// results for another project.  The UI deals with this.
			// TODO Get Group by Short ID
			return renderNotImplemented(c)
		}
		if matchingGroupID != 0 {
			//  TODO Get by Group ID
			// response['X-Sentry-Direct-Hit'] = '1'
			return renderNotImplemented(c)
		}
	}

	/* TODO is this obsolete code?
	groups := []models.Group{}
	_, err = db.SelectBySql(`
		select gm.* from sentry_groupedmessage gm`).
		LoadStructs(&groups)
	if err != nil {
		return err
	}*/
	// TODO Implement `func GetSearchBackend()` that returns SearchBackend interface
	// NOTE Sentry v8.10 implements only single DjangoSearchBackend
	// TODO CONTINUE !!!
	queryDto, err := buildQueryDto(c, project.ID)
	if err != nil {
		// TODO If validation error return JSON -> Response({'detail': six.text_type(exc)}, status=400)
		return err
	}
	sb, err := buildSelectBuilder(queryDto, db)
	if err != nil {
		return err
	}
	//sql, args := sb.ToSql()
	//pp.Print(sql)
	//pp.Print(args)
	groups := []*Group{}
	_, err = sb.LoadStructs(&groups)
	if err != nil {
		return errors.Wrap(err, "can not read groups of project issues")
	}
	if statsPeriod != "" {
		segments := int64(0)
		interval := time.Duration(0)
		// TODO add stats for the list of group/issue IDs
		if statsPeriod == "14d" {
			segments = 14
			interval = time.Duration(24 * time.Hour)
		} else if statsPeriod == "24h" {
			segments = 24
			interval = time.Duration(1 * time.Hour)
		} else {
			return errors.New(errInvalidStatsPeriod)
		}
		now := time.Now().UTC()
		groupIDs := []int{}
		for _, group := range groups {
			if group.ProjectID != nil && *group.ProjectID == project.ID {
				group.Project = GroupProjectInfo{
					Name: project.Name,
					Slug: project.Slug,
				}
			} else {
				return errors.Errorf(
					"Not implemented. Event group does not belong to the requested project. Group.ID=%d, Group.ProjectID=%d, Request.ProjectSlug=%s, Request.ProjectID=%d",
					group.ID, *group.ProjectID, project.Slug, project.ID)
			}
			group.Logger = &group.Group.Logger
			if *group.Logger == "" {
				group.Logger = nil
			}

			statusCode := group.Group.Status
			group.StatusDetails = map[string]string{}
			//if attrs['ignore_duration']:
			//	if attrs['ignore_duration'] < timezone.now() and status == GroupStatus.IGNORED:
			//		status = GroupStatus.UNRESOLVED
			//	else:
			//		status_details['ignoreUntil'] = attrs['ignore_duration']
			//elif status == GroupStatus.UNRESOLVED and obj.is_over_resolve_age():
			//	status = GroupStatus.RESOLVED
			//	status_details['autoResolved'] = True
			if statusCode == models.GroupStatusResolved {
				group.Status = "resolved"
				//if attrs['pending_resolution']:
				//    group.StatusDetails["inNextRelease"] = True
			} else if statusCode == models.GroupStatusIgnored {
				group.Status = "ignored"
			} else if statusCode == models.GroupStatusPendingDeletion || statusCode == models.GroupStatusDeletionInProgress {
				group.Status = "pending_deletion"
			} else if statusCode == models.GroupStatusPendingMerge {
				group.Status = "pending_merge"
			} else {
				group.Status = "unresolved"
			}

			if group.Group.ShortID != nil {
				group.ShortID = fmt.Sprintf(
					"%s-%s",
					strings.ToUpper(project.Slug),
					base32num.Encode(uint32(*group.Group.ShortID)))
			}

			group.Level = getLogLevelString(group.Group.Level)
			group.Permalink = fmt.Sprintf("%s://%s%s",
				c.Request().URL.Scheme,
				c.Request().URL.Host,
				c.Echo().URI(frontend.GetSentryGroupView, orgSlug, group.Project.Slug, group.ID))
			groupIDs = append(groupIDs, group.ID)
		}
		start := now.Add(-time.Duration(((segments - 1) * interval.Nanoseconds())))
		stats := tsdb.New().GetRange(tsdb.Group, groupIDs, start, now, int(interval.Seconds()))
		pp.Print(stats)
		/*
		   for item in item_list:
		       attrs[item].update({
		           'stats': stats[item.id],
		       })

		*/
	}

	/*issues := []Issue{Issue{
		TimeSpent:   nil,
		LastSeen:    "2016-10-31T15:33:51Z",
		NumComments: 0,
		UserCount:   0,
		Stats: IssueStatistic{
			For24h: []string{},
		},
		Culprit:      "__main__ in <module>",
		Title:        "This is a test message generated using ``raven test``",
		ID:           "2",
		AssignedTo:   nil,
		Logger:       nil,
		Annotations:  []string{},
		Status:       "unresolved",
		IsPublic:     false,
		HasSeen:      true,
		ShareID:      "322e32",
		FirstSeen:    "2016-10-31T15:33:51Z",
		Count:        "1",
		Permalink:    "http://localhost:9000/acme/api/issues/2/",
		Level:        "info",
		IsBookmarked: false,
		Project: ProjectRef{
			Name: "API",
			Slug: "api",
		},
		StatusDetails: map[string]string{},
	}}*/
	return c.JSON(http.StatusOK, groups)
}

type QueryDto struct {
	ProjectID         int
	Query             string // ?
	Status            *int
	Tags              string
	BookmarkedBy      string
	AssignedTo        string
	FirstRelease      *string
	SortBy            string // date (default) | priority | new | freq
	Unassigned        *bool
	AgeFrom           time.Time
	AgeFromInclusive  bool
	AgeTo             time.Time
	AgeToInclusive    bool
	DateFrom          time.Time
	DateFromInclusive bool
	DateTo            time.Time
	DateToInclusive   bool
	Cursor            string
	Limit             string
}

// TODO use const modifier
var GroupStatusChoices = map[string]int{
	"resolved":              models.GroupStatusResolved,
	"unresolved":            models.GroupStatusUnresolved,
	"ignored":               models.GroupStatusIgnored,
	"resolvedInNextRelease": models.GroupStatusUnresolved,

	// GroupStatusMuted status will be removed in Sentry 9.0
	"muted": models.GroupStatusIgnored,
}

func buildQueryDto(c echo.Context, projectID int) (QueryDto, error) {
	// TODO Develop domain-specific context that embeds echo.Context
	// TODO Do not ask for projectID, ask only assimilator.Context that includes this info
	query := QueryDto{
		ProjectID:         projectID,
		SortBy:            "date",
		AgeFromInclusive:  true,
		AgeToInclusive:    true,
		DateFromInclusive: true,
		DateToInclusive:   true,
	}
	if c.QueryParam("status") != "" {
		if status, ok := GroupStatusChoices[c.QueryParam("status")]; ok {
			query.Status = &status
		} else {
			return query, errors.New("invalid status")
		}
	}
	return query, nil
}

func buildSelectBuilder(query QueryDto, db *dbr.Tx) (*dbr.SelectBuilder, error) {
	sb := db.Select("gm.*").From("sentry_groupedmessage gm").Where("gm.project_id = ?", query.ProjectID)
	if query.Query != "" {
		// TODO(dcramer): if we want to continue to support search on SQL
		// we should at least optimize this in Postgres so that it does
		// the query filter **after** the index filters, and restricts the
		// result set
		sb = sb.Where("lower(gm.message) like lower(?) or lower(gm.view) like lower(?)", query.Query)
	}
	if query.Status == nil {
		sb = sb.Where("gm.status not in ?", []int{
			models.GroupStatusPendingDeletion,
			models.GroupStatusDeletionInProgress,
			models.GroupStatusPendingMerge,
		})
	} else {
		sb = sb.Where("gm.status = ?", query.Status)
	}
	if query.BookmarkedBy != "" {
		// TODO filter by bookmark_set__project and bookmark_set__user
	}
	if query.AssignedTo != "" {
		// TODO filter by assignee_set__project and assignee_set__user
	} else if query.Unassigned != nil {
		// TODO filter by assignee_set__isnull
	}
	if query.FirstRelease != nil {
		//if query.FirstRelease == Empty {
		//	return sb, nil
		//}
		// TODO filter by first_release__project and first_release__version
	}
	if query.Tags != "" {
		// TODO implement filter by tags
		// matches, err := tagToFilters(db, projectID, query.Tags)
		//if matches == nil {
		//	return sb, nil
		//}
		//sb = sb.Where("g.id in ?", matches)
	}
	if !query.AgeFrom.IsZero() {
		if query.AgeFromInclusive {
			sb = sb.Where("gm.first_seen >= ?", query.AgeFrom)
		} else {
			sb = sb.Where("gm.first_seen > ?", query.AgeFrom)
		}
	}
	if !query.AgeTo.IsZero() {
		if query.AgeToInclusive {
			sb = sb.Where("gm.first_seen <= ?", query.AgeTo)
		} else {
			sb = sb.Where("gm.first_seen < ?", query.AgeTo)
		}
	}
	if !query.DateFrom.IsZero() || !query.DateTo.IsZero() {
		eventQuery := db.Select("distinct m.group_id").
			From("sentry_message m").
			Where("m.project_id = ?", query.ProjectID)
		if !query.DateFrom.IsZero() {
			if !query.DateFromInclusive {
				eventQuery = eventQuery.Where("m.datetime >= ?", query.DateFrom)
			} else {
				eventQuery = eventQuery.Where("m.datetime > ?", query.DateFrom)
			}
		}
		if !query.DateTo.IsZero() {
			if !query.DateToInclusive {
				eventQuery = eventQuery.Where("m.datetime <= ?", query.DateTo)
			} else {
				eventQuery = eventQuery.Where("m.datetime < ?", query.DateTo)
			}
		}
		// Limit to the first 1000 results
		groupIDs, err := eventQuery.Limit(1000).ReturnInt64s()
		pp.Print(err)
		pp.Print(groupIDs)
		if err != nil {
			return nil, err
		}
		// if Event is not on the primary database remove Django's
		// implicit subquery by coercing to a list
		// TODO CONTINUE !! Add to SelectBuilder `where id in groupIDs`
	}
	// TODO CONTINUE !! Add order by clause for MySQL, Oracle, MS SQL, PostgreSQL, SQLite
	if query.SortBy == "date" {
		// TODO Is order by Desc or Asc
		sb = sb.OrderBy("gm.last_seen")
	} else if query.SortBy == "priority" {
		sb = sb.OrderBy("gm.score")
	} else if query.SortBy == "new" {
		sb = sb.OrderBy("gm.first_seen")
	} else if query.SortBy == "freq" {
		sb = sb.OrderBy("gm.times_seen")
	} else {
		// TODO What if SortBy is empty
	}
	return sb, nil
}

func getLogLevelString(logLevelCode int) string {
	logLevels := map[int]string{
		10: "debug",
		20: "info",
		30: "warning",
		40: "error",
		50: "fatal",
	}
	if level, ok := logLevels[logLevelCode]; ok {
		return level
	}
	return "unknown"
}
