package api

import (
	"errors"

	"github.com/labstack/echo"
)

// GetProjectID returns projectID for current HTTP request
func GetProjectID(c echo.Context) int64 {
	if projectID, ok := c.Get("projectID").(int64); ok {
		return projectID
	}
	panic(errors.New("failed to get projectID, is mw.RequireProject call missing?"))
}
