package store

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type ProjectStore struct {
	c echo.Context
}

func NewProjectStore(c echo.Context) ProjectStore {
	return ProjectStore{c: c}
}

func (s ProjectStore) GetProjectID(orgSlug, projectSlug string) (int64, error) {
	db, err := db.GetTx(s.c)
	if err != nil {
		return 0, errors.Wrapf(err, "can not get project")
	}
	projectID, err := db.SelectBySql(`
            select p.id
                from sentry_project p
                    join sentry_organization o on p.organization_id = o.id
            where o.slug = ? and p.slug = ?`,
		orgSlug, projectSlug).
		ReturnInt64()
	if err != nil {
		return 0, errors.Wrap(err, "can not get project")
	}
	return projectID, nil
}

func (s ProjectStore) GetEnvironments(projectID int64) ([]models.Environment, error) {
	db, err := db.GetTx(s.c)
	if err != nil {
		return nil, errors.Wrapf(err, "can not read project environments")
	}
	environments := []models.Environment{}
	_, err = db.SelectBySql(`
		select se.*
			from sentry_environment se
		where se.project_id = ?
        order by se.name`,
		projectID).
		LoadStructs(&environments)
	if err != nil {
		return nil, errors.Wrapf(err, "can not read project environments")
	}
	return environments, nil
}

func (s ProjectStore) GetTags(projectID int64) ([]*models.TagKey, error) {
	db, err := db.GetTx(s.c)
	if err != nil {
		return nil, errors.Wrapf(err, "can not read project tags")
	}
	tags := []*models.TagKey{}
	_, err = db.SelectBySql(`
		select fk.*
			from sentry_filterkey fk
		where fk.project_id = ? and fk.status = ?`,
		projectID, models.TagKeyStatusVisible).
		LoadStructs(&tags)
	if err != nil {
		return nil, errors.Wrapf(err, "can not read project tags")
	}
	// TODO tag.Key must be processed -> TagKey.get_standardized_key(tag_key.key)
	for _, tag := range tags {
		tag.PostGet()
	}
	return tags, nil
}

func (s ProjectStore) SaveProject(project models.Project) error {
	db, err := db.GetTx(s.c)
	if err != nil {
		return errors.Wrapf(err, "failed to save project")
	}
	_, err = db.InsertInto("sentry_project").
		Columns("id", "team_id", "organization_id", "name", "slug",
			"public", "status", "first_event", "date_added").
		Record(project).
		Exec()
	return errors.Wrapf(err, "failed to save project")
}

func (s ProjectStore) SaveTags(tags ...*models.TagKey) error {
	db, err := db.GetTx(s.c)
	if err != nil {
		return errors.Wrapf(err, "failed to save project tags")
	}
	query := db.InsertInto("sentry_filterkey").
		Columns("id", "project_id", "key", "values_seen", "label", "status")
	for _, tag := range tags {
		query = query.Record(tag)
	}
	// TODO can we just ignore rv / sql.Result?
	_, err = query.Exec()
	return errors.Wrapf(err, "failed to save project tags")
}
