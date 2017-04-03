package web

import (
	"net/http"
	"strings"

	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/log"
	"github.com/diyan/assimilator/web/recover"
	"github.com/diyan/assimilator/web/renderer"
	"github.com/gocraft/dbr"

	"github.com/GeertJohan/go.rice"
	"github.com/diyan/assimilator/db"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
)

func NewApp(config conf.Config) *echo.Echo {
	//dbTxMaker := db.TxMakerFunc{return db.New(config)}
	dbTxMaker := func() (*dbr.Tx, error) { return db.New(config) }
	return NewAppCustom(config, dbTxMaker)
}

// TODO use struct to pass settings and dependencies
func NewAppCustom(config conf.Config, dbTxMaker db.TxMakerFunc) *echo.Echo {
	e := echo.New()
	e.Debug = true
	e.Logger = log.NewEchoLogger(config)
	e.HTTPErrorHandler = recover.NewEchoErrorHandler(config)
	templateBox, err := rice.FindBox("templates")
	if err != nil {
		panic(errors.Wrap(err, "can not find template box"))
	}
	e.Renderer = renderer.New(templateBox)
	// TODO consider remove conf.FromE(...) and conf.NewMiddleware(...)
	//e.Use(conf.NewMiddleware(config))
	//e.Use(log.NewAccessLogMiddleware(config))
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

	bind := context.NewBinder(config, dbTxMaker)
	RegisterRoutes(e, bind)
	return e
}
