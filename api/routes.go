package api

import (
	api "github.com/diyan/assimilator/api/endpoints"

	"github.com/labstack/echo"
)

// RegisterAPIRoutes adds API routes to the Echo's route group
func RegisterAPIRoutes(g *echo.Group) {
	// Organizations
	g.GET("/organizations/", api.OrganizationIndexGetEndpoint)
	g.GET("/organizations/:organization_slug/", api.OrganizationDetailsGetEndpoint)

	// Projects
	g.GET("/projects/:organization_slug/:project_slug/environments/", api.ProjectEnvironmentsGetEndpoint)
	g.GET("/projects/:organization_slug/:project_slug/issues/", api.ProjectGroupIndexGetEndpoint)
	g.GET("/projects/:organization_slug/:project_slug/groups/", api.ProjectGroupIndexGetEndpoint)
	g.GET("/projects/:organization_slug/:project_slug/searches/", api.ProjectSearchesGetEndpoint)

	g.GET("/projects/:organization_slug/:project_slug/members/", api.ProjectMemberIndexGetEndpoint)
	g.GET("/projects/:organization_slug/:project_slug/tags/", api.ProjectTagsGetEndpoint)
	// Internal
	g.GET("/internal/health/", api.SystemHealthGetEndpoint)
}
