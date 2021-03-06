package api

import (
	"github.com/diyan/assimilator/context"
	"github.com/labstack/echo"
)

// RegisterAPIRoutes adds API routes to the Echo's route group
func RegisterAPIRoutes(g *echo.Group, bind context.Binder) {
	// Store endpoints first since they are the most active
	//e.GET("/api/store/", storeGetView)
	//e.POST("/api/store/", storePostView)

	// TODO Can not register same handler for two different routes
	//g = g.Group("/store")
	//g.GET("/", storeGetView)
	//g.POST("/", storePostView)
	// :project_id is [\w_-]+
	g = g.Group("/:project_id/store")
	g.GET("/", storeGetView)
	g.POST("/", bind.Base(storePostView))
	// :project_id is \d+
	g = g.Group("/:project_id/csp-report")
	// TODO is CspReportGetView needed?
	g.GET("/", cspReportGetView)
	g.POST("/", cspReportPostView)
}
