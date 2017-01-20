package api

import (
	"net/http"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"

	"github.com/labstack/echo"
)

func ProjectTagsGetEndpoint(c echo.Context) error {
	orgSlug := c.Param("organization_slug")
	projectSlug := c.Param("project_slug")
	db, err := db.GetSession()
	if err != nil {
		return err
	}
	projectID, err := db.SelectBySql(`
		select p.ID
			from sentry_project p
				join sentry_team t on p.team_id = t.id
				join sentry_organization o on t.organization_id = o.id
		where o.slug = ? and p.slug = ? and p.status = ?`,
		orgSlug, projectSlug, models.ProjectStatusVisible).
		ReturnInt64()
	if err != nil {
		// TODO return ResourceDoesNotExist if record was not found
		return err
	}
	// TODO check project permissions -> self.check_object_permissions(request, project)

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
