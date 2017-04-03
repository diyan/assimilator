package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/db/store"
)

func ProjectTagsGetEndpoint(c context.Project) error {
	projectStore := store.NewProjectStore()
	tags, err := projectStore.GetTags(c.Tx, c.Project.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, tags)
}
