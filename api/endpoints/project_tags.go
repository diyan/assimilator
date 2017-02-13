package api

import (
	"net/http"

	"github.com/diyan/assimilator/db/store"
	"github.com/labstack/echo"
)

func ProjectTagsGetEndpoint(c echo.Context) error {
	projectID := GetProjectID(c)
	projectStore := store.NewProjectStore(c)
	tags, err := projectStore.GetTags(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tags)
}
