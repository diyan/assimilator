package frontend

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"Ooops! This page has not ported from Sentry yet")
}

func GetStaticMedia(c echo.Context) error {
	return RenderNotImplemented(c)
}
