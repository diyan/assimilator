package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/labstack/echo"
)

// OrganizationDetails ...
type OrganizationDetails struct {
	models.Organization
	PendingAccessRequests int      `json:"pendingAccessRequests"`
	Features              []string `json:"features"`
	Quota                 Quota    `json:"quota"`
	Access                []string `json:"access"`
	Teams                 []Team   `json:"teams"`
}

// Quota ..
type Quota struct {
	ProjectLimit int `db:"project_limit" json:"projectLimit"`
	MaxRate      int `json:"maxRate"`
}

// Project ..
type Project struct {
	models.Project
	IsPublic bool     `json:"isPublic"`
	Features []string `json:"features"`
}

// Team ...
type Team struct {
	models.Team
	HasAccess bool      `db:"has_access" json:"hasAccess"`
	IsPending bool      `db:"is_pending" json:"isPending"`
	IsMember  bool      `db:"is_member" json:"isMember"`
	Projects  []Project `json:"projects"`
}

func OrganizationDetailsGetEndpoint(c echo.Context) error {
	orgSlug := c.Param("organization_slug")
	userID := 1 // TODO get ID from context.request.user.id
	db, err := db.GetTx(c)
	if err != nil {
		return err
	}

	// TODO fill in pendingAccessRequest -> select count(*) from sentry_organizationaccessrequest ...
	// TODO fill in org.features
	// if env.request:
	//     context['access'] = access.from_request(env.request, obj).scopes
	// else:
	//     context['access'] = access.from_user(user, obj).scopes
	// TODO fill in onboardingTasks
	// context['onboardingTasks'] = serialize(onboarding_tasks, user, OnboardingTasksSerializer())
	org := OrganizationDetails{}
	err = db.SelectBySql(`
		select
			o.*
		from sentry_organization o
			join sentry_organizationmember om on o.id = om.organization_id
		where o.slug = ? and om.user_id = ? and o.status = ?`,
		orgSlug, userID, models.OrganizationStatusVisible).
		LoadStruct(&org)
	if err != nil {
		return err
	}
	quota := Quota{}
	defaultProjectLimit := 100
	err = db.SelectBySql(`
		select coalesce(cast(oo.value as int), ?) as project_limit
		from sentry_organization o
			left join sentry_organizationoptions oo
				on o.id = oo.organization_id and key = 'sentry:project-rate-limit'
		where o.slug = ?`,
		defaultProjectLimit, orgSlug).
		LoadStruct(&quota)
	if err != nil {
		return err
	}
	// TODO fill in quota.maxRate -> 'maxRate': quotas.get_organization_quota(obj),
	org.Quota = quota

	teams := []Team{}
	// TODO fill has_access field
	_, err = db.SelectBySql(`
		select
			t.*,
			om.id is not null as is_member,
			oar.id is not null as is_pending,
			1 as has_access -- TODO workaround
		from sentry_team t
			join sentry_organization o on t.organization_id = o.id
			left join sentry_organizationmember_teams omt on t.id = omt.team_id
			left join sentry_organizationmember om on omt.organizationmember_id = om.id
			left join sentry_organizationaccessrequest oar
				on t.id = oar.team_id and om.id = oar.member_id
		where o.slug = ? and om.user_id = ? and t.status = ?`,
		orgSlug, userID, models.TeamStatusVisible).
		LoadStructs(&teams)
	/*
		_, err = db.Select("sentry_team.*").From("sentry_team").
			Join("sentry_organization", "sentry_team.organization_id = sentry_organization.id").
			Where("sentry_organization.slug = ?", orgSlug).
			Where("sentry_team.status = ?", 0).
			LoadStructs(&teams)
	*/
	if err != nil {
		return err
	}
	teamIDs := []int{}
	for _, team := range teams {
		// TODO has_access is true if one of following are true:
		// is_member is true
		// team.organization.flags.allow_joinleave
		// request.is_superuser()
		teamIDs = append(teamIDs, team.ID)
	}

	// TODO fill project.features array
	projects := []Project{}
	_, err = db.SelectBySql(`
		select p.*
			from sentry_project p
		where p.team_id in ? and p.status = ?
		order by p.name, p.slug`,
		teamIDs, models.ProjectStatusVisible).
		LoadStructs(&projects)
	if err != nil {
		return err
	}
	for _, project := range projects {
		for i := 0; i < len(teams); i++ {
			team := teams[i]
			if project.TeamID == team.ID {
				team.Projects = append(team.Projects, project)
				teams[i] = team
			}
		}
	}
	org.Teams = teams
	org.Features = []string{"sso", "open-membership"}
	org.Access = []string{"org:write", "member:write", "project:read", "org:delete", "org:read", "event:read", "team:delete", "member:delete", "event:delete", "member:read", "project:write", "event:write", "project:delete", "team:read", "team:write"}
	return c.JSON(http.StatusOK, org)
}

/* EXPECTED RESPONSE
curl -X GET http://localhost:9001/api/0/organizations/acme
{
    "pendingAccessRequests": 0,
    "slug": "acme",
    "name": "ACME",
    "quota": {
        "projectLimit": 100,
        "maxRate": 0
    },
    "dateCreated": "2016-11-10T11:27:51.509Z",
    "access": [
        "org:write",
        "member:write",
        "project:read",
        "org:delete",
        "org:read",
        "event:read",
        "team:delete",
        "member:delete",
        "event:delete",
        "member:read",
        "project:write",
        "event:write",
        "project:delete",
        "team:read",
        "team:write"
    ],
    "teams": [
        {
            "slug": "acme",
            "name": "ACME",
            "hasAccess": true,
            "isPending": false,
            "dateCreated": "2016-11-10T11:27:51.522Z",
            "isMember": true,
            "id": "2",
            "projects": [
                {
                    "slug": "api",
                    "name": "API",
                    "isPublic": false,
                    "dateCreated": "2016-11-10T11:27:52.646Z",
                    "firstEvent": "2016-11-10T11:31:27Z",
                    "id": "2",
                    "features": [
                        "quotas"
                    ]
                }
            ]
        }
    ],
    "id": "2",
    "features": [
        "sso",
        "open-membership"
    ]
}
*/
