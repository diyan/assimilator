package web

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/template"
	"github.com/diyan/echox/log"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"

	mwx "github.com/diyan/echox/middleware"
	mw "github.com/labstack/echo/middleware"
	logrusfmt "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	template.RegisterTags()
	template.RegisterFilters()
}

func GetApp(config conf.Config) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger = log.Logrus()
	// Register default error handler
	e.HTTPErrorHandler = HTTPErrorHandler
	e.Renderer = pongor.GetRenderer(pongor.PongorOption{Reload: true})
	e.Use(conf.NewMiddleware(config))
	// TODO ForceColors only if codegangsta/gin detected
	logrus.SetFormatter(&logrusfmt.TextFormatter{ShortTimestamp: true, ForceColors: true})
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
