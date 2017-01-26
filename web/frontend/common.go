package frontend

import (
	"net/http"

	"github.com/labstack/echo"
)

func renderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"Ooops! This page has not ported from Sentry yet")
}

func getStaticMedia(c echo.Context) error {
	return renderNotImplemented(c)
}
