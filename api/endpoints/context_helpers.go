package api

import "github.com/labstack/echo"

// GetProjectID returns projectID for current HTTP request
func GetProjectID(c echo.Context) int64 {
	return c.Get("projectID").(int64)
}
