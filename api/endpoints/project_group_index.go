package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/gocraft/dbr"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

const errInvalidStatsPeriod = "Invalid stats_period. Valid choices are '', '24h', and '14d'"

// Issue ...
type Issue struct {
	ID            string            `json:"id"`
	FirstSeen     string            `json:"firstSeen"`
	LastSeen      string            `json:"lastSeen"`
	TimeSpent     *string           `json:"timeSpent"` // TODO check type
	NumComments   int               `json:"numComments"`
	UserCount     int               `json:"userCount"`
	Stats         IssueStatistic    `json:"stats"`
	Culprit       string            `json:"culprit"`
	Title         string            `json:"title"`
	AssignedTo    *string           `json:"assignedTo"` // TODO check type
	Logger        *string           `json:"logger"`     // TODO check type
	Annotations   []string          `json:"annotations"`
	Status        string            `json:"status"`
	IsPublic      bool              `json:"isPublic"`
	HasSeen       bool              `json:"hasSeen"`
	ShareID       string            `json:"shareId"`
	Count         string            `json:"count"`
	Permalink     string            `json:"permalink"`
	Level         string            `json:"level"`
	IsBookmarked  bool              `json:"isBookmarked"`
	Project       ProjectRef        `json:"project"`
	StatusDetails map[string]string `json:"statusDetails"` // TODO check type
}

// IssueStatistic ...
type IssueStatistic struct {
	For24h []string `json:"24h"`
}

// ProjectRef ...
type ProjectRef struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func ProjectGroupIndexGetEndpoint(c echo.Context) error {
	projectID := MustGetProjectID(c)
	statsPeriod := c.QueryParam("statsPeriod")
	shortIDLookup, _ := strconv.ParseBool("shortIdLookup")
	if !(statsPeriod == "" || statsPeriod == "24h" || statsPeriod == "14d") {
		// TODO introduce better error handling -> return err.InvalidStatsPeriod
		return c.JSON(400, map[string]string{"detail": errInvalidStatsPeriod})
	}
	query := strings.TrimSpace(c.QueryParam("query"))
	db, err := db.GetTx(c)
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
				projectID, query).
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

	groups := []models.Group{}
	_, err = db.SelectBySql(`
		select gm.* from sentry_groupedmessage gm`).
		LoadStructs(&groups)
	if err != nil {
		return err
	}
	// TODO Implement `func GetSearchBackend()` that returns SearchBackend interface
	// NOTE Sentry v8.10 implements only single DjangoSearchBackend
	// TODO CONTINUE !!!
	queryDto, err := buildQueryDto(c, projectID)
	if err != nil {
		// TODO If validation error return JSON -> Response({'detail': six.text_type(exc)}, status=400)
		return err
	}
	sb, err := buildSelectBuilder(queryDto, db)
	if err != nil {
		return err
	}
	sql, args := sb.ToSql()
	pp.Print(sql)
	pp.Print(args)
	issues := []models.Group{}
	_, err = sb.LoadStructs(&issues)
	if err != nil {
		return err
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
	return c.JSON(http.StatusOK, issues)
}

type QueryDto struct {
	ProjectID         int64
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

func buildQueryDto(c echo.Context, projectID int64) (QueryDto, error) {
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

/* EXPECTED RESPONSE
curl 'http://localhost:9000/api/0/projects/acme/api/issues/?limit=25&statsPeriod=24h&query=is%3Aunresolved+'
[
    {
        "timeSpent": null,
        "lastSeen": "2016-11-10T11:31:35Z",
        "numComments": 0,
        "userCount": 0,
        "stats": {
            "24h": []
        },
        "culprit": "/tmp/raven/bin/raven test http://571a1ad9bc9245329b22af2731db79d0:7910613d12ce478483eb9da048e45bed@localhost:9000/2",
        "title": "This is a test message generated using ``raven test``",
        "id": "1",
        "assignedTo": null,
        "logger": null,
        "annotations": [],
        "status": "unresolved",
        "isPublic": false,
        "hasSeen": false,
        "shareId": "322e31",
        "firstSeen": "2016-11-10T11:31:27Z",
        "count": "3",
        "permalink": "http://localhost:9000/acme/api/issues/1/",
        "level": "info",
        "isBookmarked": false,
        "project": {
            "name": "API",
            "slug": "api"
        },
        "statusDetails": {}
    }
]
*/
