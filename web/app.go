package web

import (
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/web/renderer"
	"github.com/diyan/echox/log"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	mwx "github.com/diyan/echox/middleware"
	mw "github.com/labstack/echo/middleware"
)

func GetApp(config conf.Config) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger = log.Logrus()
	// Register default error handler
	e.HTTPErrorHandler = HTTPErrorHandler
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		panic(errors.Wrap(err, "can not find template box"))
	}
	e.Renderer = renderer.New(templateBox)
	e.Use(conf.NewMiddleware(config))
	// TOOD add configuration flag to enable/disable access logs
	// Register access log logger
	e.Use(mwx.LogrusLogger(nil))
	e.Use(RecoverMiddleware())
	// TODO setup route to serve static files
	e.Static("/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry", "ui")
	e.Pre(mw.AddTrailingSlashWithConfig(mw.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			uri := c.Request().URL.String()
			// TODO consider check suffixes .xml .ico .txt .json .js .svg .png .gif .html .eot .woff .css .ttf .woff2
			if strings.HasPrefix(uri, "/_static/") ||
				strings.HasSuffix(uri, "/robots.txt") ||
				strings.HasSuffix(uri, "/favicon.ico") ||
				strings.HasSuffix(uri, "/crossdomain.xml") {
				return true
			}
			return false
		}}))

	RegisterRoutes(e)

	return e
}
