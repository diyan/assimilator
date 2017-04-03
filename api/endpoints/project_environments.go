package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/db/store"
)

func ProjectEnvironmentsGetEndpoint(c context.Project) error {
	projectStore := store.NewProjectStore()
	environments, err := projectStore.GetEnvironments(c.Tx, c.Project.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, environments)
}
