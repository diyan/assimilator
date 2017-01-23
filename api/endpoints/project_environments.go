package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func ProjectEnvironmentsGetEndpoint(c echo.Context) error {
	projectID := MustGetProjectID(c)
	db, err := db.GetSession(c)
	if err != nil {
		return errors.Wrap(err, "can not get db session")
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
