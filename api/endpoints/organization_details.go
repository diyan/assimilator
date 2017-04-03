package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
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

func OrganizationDetailsGetEndpoint(c context.Organization) error {
	// TODO fill in pendingAccessRequest -> select count(*) from sentry_organizationaccessrequest ...
	// TODO fill in org.features
	// if env.request:
	//     context['access'] = access.from_request(env.request, obj).scopes
	// else:
	//     context['access'] = access.from_user(user, obj).scopes
	// TODO fill in onboardingTasks
	// context['onboardingTasks'] = serialize(onboarding_tasks, user, OnboardingTasksSerializer())
	org := OrganizationDetails{Organization: c.Organization}
	defaultProjectLimit := 100
	err := c.Tx.SelectBySql(`
		select cast(oo.value as int) as project_limit
		from sentry_organizationoptions oo
		where oo.organization_id = ? and oo.key = 'sentry:project-rate-limit'`,
		c.Organization.ID).
		LoadStruct(&org.Quota)
	if err == dbr.ErrNotFound {
		org.Quota.ProjectLimit = defaultProjectLimit
	} else if err != nil {
		return errors.Wrap(err, "can not read project limit from organization options")
	}
	// TODO fill in quota.maxRate -> 'maxRate': quotas.get_organization_quota(obj),
	// TODO fill has_access field
	_, err = c.Tx.SelectBySql(`
		select
			t.*,
			om.id is not null as is_member,
			oar.id is not null as is_pending,
			1 as has_access -- TODO workaround
		from sentry_team t
			left join sentry_organizationmember_teams omt on t.id = omt.team_id
			left join sentry_organizationmember om on omt.organizationmember_id = om.id
			left join sentry_organizationaccessrequest oar
				on t.id = oar.team_id and om.id = oar.member_id
		where t.organization_id = ? and om.user_id = ? and t.status = ?`,
		c.Organization.ID, c.User.ID, models.TeamStatusVisible).
		LoadStructs(&org.Teams)
	/*
		_, err = db.Select("sentry_team.*").From("sentry_team").
			Join("sentry_organization", "sentry_team.organization_id = sentry_organization.id").
			Where("sentry_organization.slug = ?", orgSlug).
			Where("sentry_team.status = ?", 0).
			LoadStructs(&teams)
	*/
	if err != nil {
		return errors.Wrap(err, "can not read organization membership")
	}

	// TODO fill project.features array
	projects := []Project{}
	if len(org.Teams) > 0 {
		teamIDs := []int{}
		for _, team := range org.Teams {
			// TODO has_access is true if one of following are true:
			// is_member is true
			// team.organization.flags.allow_joinleave
			// request.is_superuser()
			teamIDs = append(teamIDs, team.ID)
		}
		_, err = c.Tx.SelectBySql(`
		select p.*
			from sentry_project p
		where p.team_id in ? and p.status = ?
		order by p.name, p.slug`,
			teamIDs, models.ProjectStatusVisible).
			LoadStructs(&projects)
		if err != nil {
			return err
		}
	}

	for _, project := range projects {
		for i := 0; i < len(org.Teams); i++ {
			team := org.Teams[i]
			if project.TeamID == team.ID {
				team.Projects = append(team.Projects, project)
				org.Teams[i] = team
			}
		}
	}
	org.Features = []string{"sso", "open-membership"}
	org.Access = []string{"org:write", "member:write", "project:read", "org:delete", "org:read", "event:read", "team:delete", "member:delete", "event:delete", "member:read", "project:write", "event:write", "project:delete", "team:read", "team:write"}
	return c.JSON(http.StatusOK, org)
}
