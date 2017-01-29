package middleware

import (
	"github.com/diyan/assimilator/db/store"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func RequireOrganization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

func RequireProject(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		orgSlug := c.Param("organization_slug")
		// TODO check with regex pattern, validate if that orgSlug exists in db
		if orgSlug == "" {
			return errors.New("'organization_slug' was not provided")
		}
		projectSlug := c.Param("project_slug")
		if projectSlug == "" {
			return errors.New("'project_slug' was not provided")
		}
		projectStore := store.NewProjectStore(c)
		projectID, err := projectStore.GetProjectID(orgSlug, projectSlug)
		if err != nil {
			return err
		}
		c.Set("projectID", projectID)
		return next(c)
		// TODO return ResourceDoesNotExist if record was not found
		// TODO check project permissions -> self.check_object_permissions(request, project)
	}
}
