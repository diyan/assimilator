package api

import (
	"errors"

	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
)

// GetProject returns project for current HTTP request
func GetProject(c echo.Context) models.Project {
	if project, ok := c.Get("project").(models.Project); ok {
		return project
	}
	panic(errors.New("failed to get project, is mw.RequireProject call missing?"))
}
