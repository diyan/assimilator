package api

import (
	"net/http"

	"github.com/diyan/assimilator/db/store"
	"github.com/labstack/echo"
)

func ProjectEnvironmentsGetEndpoint(c echo.Context) error {
	projectID := GetProjectID(c)
	projectStore := store.NewProjectStore(c)
	environments, err := projectStore.GetEnvironments(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, environments)
}
