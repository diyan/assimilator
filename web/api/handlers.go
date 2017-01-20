package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"We are sorry, this page was not yet ported from Sentry")
}

func StoreGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func StorePostView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func CspReportGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func CspReportPostView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func GetRobotsTxt(c echo.Context) error {
	return RenderNotImplemented(c)
}

func GetCrossdomainXMLIndex(c echo.Context) error {
	return RenderNotImplemented(c)
}

func GetCrossdomainXml(c echo.Context) error {
	return RenderNotImplemented(c)
}
