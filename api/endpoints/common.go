package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderNotImplemented(c echo.Context) error {
	return c.JSON(
		http.StatusNotImplemented,
		"Ooops! This page has not ported from Sentry yet")
}
