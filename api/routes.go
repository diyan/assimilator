package api

import (
	api "github.com/diyan/assimilator/api/endpoints"
	mw "github.com/diyan/assimilator/api/middleware"
	"github.com/labstack/echo"
)

// RegisterAPIRoutes adds API routes to the Echo's route group
func RegisterAPIRoutes(g *echo.Group) {
	// Organizations
	g.GET("/organizations/", api.OrganizationIndexGetEndpoint)
	g.GET("/organizations/:organization_slug/", api.OrganizationDetailsGetEndpoint)

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

	// Internal
	g.GET("/internal/health/", api.SystemHealthGetEndpoint)
}
