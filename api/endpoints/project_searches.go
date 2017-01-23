package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func ProjectSearchesGetEndpoint(c echo.Context) error {
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
	searches := []models.SavedSearch{}
	_, err = db.SelectBySql(`
		select ss.*
			from sentry_savedsearch ss
		where ss.project_id = ?
        order by ss.name`,
		projectID).
		LoadStructs(&searches)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, searches)
}

func ProjectSearchesPostEndpoint(c echo.Context) error {
	return RenderNotImplemented(c)
}
