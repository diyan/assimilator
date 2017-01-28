package api

import (
	"github.com/diyan/assimilator/db"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// MustGetProjectID is like GetProjectID but panics if `organization_slug`
//  or `project_slug` was not passed in HTTP request
// TODO context_helpers a bad name, later we need to split this module into the
//  several modules with good names.
func MustGetProjectID(c echo.Context) int64 {
	projectID, err := GetProjectID(c)
	if err != nil {
		panic(err)
	}
	return projectID
}

// GetProjectID returns projectID for current HTTP request
// TODO add argument, so we can return all projects or
//   projects that only visible to current user
func GetProjectID(c echo.Context) (int64, error) {
	orgSlug := c.Param("organization_slug")
	projectSlug := c.Param("project_slug")
	db, err := db.GetTx(c)
	if err != nil {
		return 0, err
	}
	// TODO validate that orgSlug is exists
	projectID, err := db.SelectBySql(`
		select p.id
			from sentry_project p
				join sentry_organization o on p.organization_id = o.id
		where o.slug = ? and p.slug = ?`,
		orgSlug, projectSlug).
		ReturnInt64()
	return projectID, errors.Wrap(err, "can not get project")
	// TODO return ResourceDoesNotExist if record was not found
	// TODO check project permissions -> self.check_object_permissions(request, project)
}
