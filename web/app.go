package web

import (
	"net/http"
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
	uiBox, err := rice.FindBox("../ui")
	if err != nil {
		panic(errors.Wrap(err, "can not find ui box"))
	}
	// TODO Extract into ui.RegisterRoutes(e) func
	uiPrefix := "/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/"
	e.GET(
		uiPrefix+"*",
		echo.WrapHandler(http.StripPrefix(uiPrefix, http.FileServer(uiBox.HTTPBox()))))
	e.Pre(mw.AddTrailingSlashWithConfig(mw.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			// TODO consider check suffixes .xml .ico .txt .json .js .svg .png .gif .html .eot .woff .css .ttf .woff2
			if strings.HasPrefix(path, "/_static/") ||
				strings.HasSuffix(path, "/robots.txt") ||
				strings.HasSuffix(path, "/favicon.ico") ||
				strings.HasSuffix(path, "/crossdomain.xml") {
				return true
			}
			return false
		}}))

	RegisterRoutes(e)

	return e
}
