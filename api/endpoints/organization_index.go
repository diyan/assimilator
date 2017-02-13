package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/labstack/echo"
)

func OrganizationIndexGetEndpoint(c echo.Context) error {
	//memberOnly := c.Param("member")
	db, err := db.FromE(c)
	if err != nil {
		return err
	}
	orgs := []models.Organization{}
	_, err = db.SelectBySql(`
		select o.* from sentry_organization o`).
		LoadStructs(&orgs)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orgs)
}
