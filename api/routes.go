package api

import (
	api "github.com/diyan/assimilator/api/endpoints"
	"github.com/diyan/assimilator/context"
	"github.com/labstack/echo"
)

// RegisterAPIRoutes adds API routes to the Echo's route group
func RegisterAPIRoutes(g *echo.Group, bind context.Binder) {
	// API tokens
	// TODO implement /api-tokens/

	// Auth
	// TODO implement /auth/

	// Broadcasts
	g.GET("/broadcasts/", api.BroadcastIndexGetEndpoint)
	g.PUT("/broadcasts/", api.BroadcastIndexPutEndpoint)

	// Users
	// TODO implement
	// /users/
	// /users/:user_id/
	// /users/:user_id/avatar/
	// /users/:user_id/authenticators/:auth_id/
	// /users/:user_id/identities/:identity_id/
	// /users/:user_id/organizations/

	// Organizations
	g.GET("/organizations/", bind.Base(api.OrganizationIndexGetEndpoint))
	g.GET("/organizations/:organization_slug/", bind.Organization(api.OrganizationDetailsGetEndpoint))

	// Teams
	// TODO implement

	// Projects
	p := g.Group("/projects/:organization_slug/:project_slug")
	//p.Use(mw.RequireUser)
	p.GET("/environments/", bind.Project(api.ProjectEnvironmentsGetEndpoint))
	p.GET("/issues/", bind.Project(api.ProjectGroupIndexGetEndpoint))
	p.GET("/groups/", bind.Project(api.ProjectGroupIndexGetEndpoint))
	p.GET("/searches/", bind.Project(api.ProjectSearchesGetEndpoint))
	p.GET("/members/", bind.Project(api.ProjectMemberIndexGetEndpoint))
	p.GET("/tags/", bind.Project(api.ProjectTagsGetEndpoint))

	// Groups
	g.GET("/issues/:issue_id/", api.GroupDetailsGetEndpoint)
	// TODO implement GroupEventsGetEndpoint
	//g.GET("/issues/:issue_id/events/", api.GroupEventsGetEndpoint)
	g.GET("/issues/:issue_id/events/latest/", bind.Base(api.GroupEventsLatestGetEndpoint))
	// ...
	g.GET("/issues/:issue_id/environments/:environment/", api.GroupEnvironmentDetailsGetEndpoint)

	// Events

	// Internal
	g.GET("/internal/health/", api.SystemHealthGetEndpoint)
}
