package web

import (
	"net/http"

	apiV0 "github.com/diyan/assimilator/api"
	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/web/api"
	"github.com/diyan/assimilator/web/frontend"
	"github.com/diyan/assimilator/web/frontend/debug"

	"github.com/labstack/echo"
)

/*
APIView base class has HTTP OPTIONS handler that respond with actual list of supported methods


*/

// RegisterRoutes ..
func RegisterRoutes(e *echo.Echo, bind context.Binder) {
	// TODO call registerDebugViews only if getattr(settings, 'DEBUG_VIEWS', settings.DEBUG)
	g := e.Group("/debug")
	debug.RegisterDebugViews(g)
	// The static version is either a 10 digit timestamp, a sha1, or md5 hash
	// :version \d{10}|[a-f0-9]{32,40}
	// TODO Use general-purpose static middleware or custom implementation?
	//e.GET("/_static/:module/*", frontend.GetStaticMedia)
	//e.GET("/_static/:version/:module/*", frontend.GetStaticMedia)

	// API
	g = e.Group("/api")
	api.RegisterAPIRoutes(g, bind)
	// API version 0
	g = e.Group("/api/0")
	apiV0.RegisterAPIRoutes(g, bind)
	frontend.RegisterFrontendRoutes(e, bind)

	// Legacy Redirects
	e.GET("/docs/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://docs.sentry.io/hosted/")
	})
	e.GET("/docs/api/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://docs.sentry.io/hosted/api/")
	})

	e.GET("/robots.txt", api.GetRobotsTxt)

	// Force a 404 of favicon.ico.
	// This url is commonly requested by browsers, and without
	// blocking this, it was treated as a 200 OK for a react page view.
	// A side effect of this is it may cause a bad redirect when logging in
	// since this gets stored in session as the last viewed page.
	// See: https://github.com/getsentry/sentry/issues/2195
	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.NoContent(404)
	})
	// crossdomain.xml
	e.GET("/crossdomain.xml", api.GetCrossdomainXMLIndex)
	e.GET("/api/:project_id/crossdomain.xml", api.GetCrossdomainXML)

	// plugins
	// TODO mount plugin handlers under /plugins/* prefix
	// e.GET("/plugins/*", include('sentry.plugins.base.urls'))

	// Legacy
	// TODO check original implementation
	// url(r'/", react_page_view),
}
