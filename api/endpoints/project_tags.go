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
