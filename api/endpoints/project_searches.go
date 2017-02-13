package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
)

func ProjectSearchesGetEndpoint(c echo.Context) error {
	projectID := GetProjectID(c)
	db, err := db.FromE(c)
	if err != nil {
		return err
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
	return renderNotImplemented(c)
}
