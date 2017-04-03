package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/models"
)

func ProjectSearchesGetEndpoint(c context.Project) error {
	searches := []models.SavedSearch{}
	_, err := c.Tx.SelectBySql(`
		select ss.*
			from sentry_savedsearch ss
		where ss.project_id = ?
        order by ss.name`,
		c.Project.ID).
		LoadStructs(&searches)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, searches)
}

func ProjectSearchesPostEndpoint(c context.Project) error {
	return renderNotImplemented(c)
}
