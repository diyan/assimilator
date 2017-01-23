package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/labstack/echo"
)

func ProjectTagsGetEndpoint(c echo.Context) error {
	projectID := MustGetProjectID(c)
	db, err := db.GetSession(c)
	if err != nil {
		return err
	}
	tags := []*models.TagKey{}
	_, err = db.SelectBySql(`
		select fk.*
			from sentry_filterkey fk
		where fk.project_id = ? and fk.status = ?`,
		projectID, models.TagKeyStatusVisible).
		LoadStructs(&tags)
	if err != nil {
		return err
	}
	// TODO tag.Key must be processed -> TagKey.get_standardized_key(tag_key.key)
	for _, tag := range tags {
		tag.PostGet()
	}
	return c.JSON(http.StatusOK, tags)
}

/* EXPECTED RESPONSE
curl -X GET http://localhost:9001/api/0/projects/acme/api/tags/
[
    {
        "uniqueValues": 1,
        "id": "2",
        "key": "server_name",
        "name": "Server"
    },
    {
        "uniqueValues": 1,
        "id": "1",
        "key": "level",
        "name": "Level"
    }
]
*/
