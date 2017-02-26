package api

import (
	api "github.com/diyan/assimilator/api/endpoints"
	mw "github.com/diyan/assimilator/api/middleware"
	"github.com/labstack/echo"
)

// RegisterAPIRoutes adds API routes to the Echo's route group
func RegisterAPIRoutes(g *echo.Group) {
	// API tokens
	// TODO implement /api-tokens/

	// Auth
	// TODO implement /auth/

	// Broadcasts
	g.GET("/broadcasts/", api.BroadcastIndexGetEndpoint)
	g.GET("/broadcasts/", api.BroadcastIndexPutEndpoint)

	// Users
	// TODO implement
	// /users/
	// /users/:user_id/
	// /users/:user_id/avatar/
	// /users/:user_id/authenticators/:auth_id/
	// /users/:user_id/identities/:identity_id/
	// /users/:user_id/organizations/

	// Organizations
	g.GET("/organizations/", api.OrganizationIndexGetEndpoint)
	g.GET("/organizations/:organization_slug/", api.OrganizationDetailsGetEndpoint)

	// Teams
	// TODO implement

	// Projects
	p := g.Group("/projects/:organization_slug/:project_slug")
	//p.Use(mw.RequireUser)
	p.Use(mw.RequireOrganization)
	p.Use(mw.RequireProject)
	p.GET("/environments/", api.ProjectEnvironmentsGetEndpoint)
	p.GET("/issues/", api.ProjectGroupIndexGetEndpoint)
	p.GET("/groups/", api.ProjectGroupIndexGetEndpoint)
	p.GET("/searches/", api.ProjectSearchesGetEndpoint)
	p.GET("/members/", api.ProjectMemberIndexGetEndpoint)
	p.GET("/tags/", api.ProjectTagsGetEndpoint)

	// Groups
	g.GET("/issues/:issue_id/", api.GroupDetailsGetEndpoint)
	// TODO implement GroupEventsGetEndpoint
	//g.GET("/issues/:issue_id/events/", api.GroupEventsGetEndpoint)
	g.GET("/issues/:issue_id/events/latest/", api.GroupEventsLatestGetEndpoint)
	// ...
	g.GET("/issues/:issue_id/environments/:environment/", api.GroupEnvironmentDetailsGetEndpoint)

	// Events

	// Internal
	g.GET("/internal/health/", api.SystemHealthGetEndpoint)
}
