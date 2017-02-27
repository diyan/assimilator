package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func renderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"We are sorry, this page was not yet ported from Sentry")
}

func cspReportGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func cspReportPostView(c echo.Context) error {
	return renderNotImplemented(c)
}

func GetRobotsTxt(c echo.Context) error {
	return renderNotImplemented(c)
}

func GetCrossdomainXMLIndex(c echo.Context) error {
	return renderNotImplemented(c)
}

func GetCrossdomainXML(c echo.Context) error {
	return renderNotImplemented(c)
}
