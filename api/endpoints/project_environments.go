package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func ProjectEnvironmentsGetEndpoint(c echo.Context) error {
	orgSlug := c.Param("organization_slug")
	projectSlug := c.Param("project_slug")
	db, err := db.GetSession()
	if err != nil {
		return errors.Wrap(err, "can not get db session")
	}
	projectID, err := db.SelectBySql(`
		select p.id
			from sentry_project p
				join sentry_organization o on p.organization_id = o.id
		where o.slug = ? and p.slug = ?`,
		orgSlug, projectSlug).
		ReturnInt64()
	if err != nil {
		return errors.Wrap(err, "can not read project")
	}
	environments := []models.Environment{}
	_, err = db.SelectBySql(`
		select se.*
			from sentry_environment se
		where se.project_id = ?
        order by se.name`,
		projectID).
		LoadStructs(&environments)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, environments)
}
