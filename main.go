package main

import (
	"os"

	"github.com/diyan/assimilator/template"
	"github.com/diyan/assimilator/web"

	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/echox/log"
	mwx "github.com/diyan/echox/middleware"
	"github.com/echo-contrib/pongor"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	logrusfmt "github.com/x-cray/logrus-prefixed-formatter"
)

// TODO keep main.go small, move everything to the web/server.go
func init() {
	template.RegisterTags()
	template.RegisterFilters()
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Logger = log.Logrus()
	// Register default error handler
	e.HTTPErrorHandler = web.HTTPErrorHandler
	e.Renderer = pongor.GetRenderer(pongor.PongorOption{Reload: true})
	// TODO ForceColors only if codegangsta/gin detected
	logrus.SetFormatter(&logrusfmt.TextFormatter{ShortTimestamp: true, ForceColors: true})
	// Register access log logger
	e.Use(mwx.LogrusLogger(nil))
	e.Use(web.RecoverMiddleware())
	// TODO setup route to serve static files
	e.Static("/_static", "_static")
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

	web.RegisterRoutes(e)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}
	e.Logger.Fatal(e.Start(port))
}
