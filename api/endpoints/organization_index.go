package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/models"
)

func OrganizationIndexGetEndpoint(c context.Base) error {
	// TODO implement memberOnly flag
	//memberOnly := c.Param("member")
	orgs := []models.Organization{}
	_, err := c.Tx.SelectBySql(`
		select o.* from sentry_organization o`).
		LoadStructs(&orgs)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orgs)
}
