package web

import (
	"strings"

	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/log"
	"github.com/diyan/assimilator/web/recover"
	"github.com/diyan/assimilator/web/renderer"

	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func NewApp(config conf.Config) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger = log.NewEchoLogger(config)
	e.HTTPErrorHandler = recover.NewEchoErrorHandler(config)
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		panic(errors.Wrap(err, "can not find template box"))
	}
	e.Renderer = renderer.New(templateBox)
	e.Use(conf.NewMiddleware(config))
	e.Use(log.NewAccessLogMiddleware(config))
	e.Use(recover.NewMiddleware(config))
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
