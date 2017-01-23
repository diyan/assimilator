package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/labstack/echo"
)

func OrganizationIndexGetEndpoint(c echo.Context) error {
	//memberOnly := c.Param("member")
	db, err := db.GetSession(c)
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

/* EXPECTED RESPONSE
curl -X GET http://localhost:9001/api/0/organizations                                                       ✹ ✭
[
  {
    "id": "1",
    "name": "Sentry",
    "slug": "sentry",
    "dateCreated": "2016-11-10T11:26:23.079809Z"
  },
  {
    "id": "2",
    "name": "ACME",
    "slug": "acme",
    "dateCreated": "2016-11-10T11:27:51.50939Z"
  }
]
*/
