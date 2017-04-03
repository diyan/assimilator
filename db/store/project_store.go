package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type ProjectStore struct {
}

func NewProjectStore() ProjectStore {
	return ProjectStore{}
}

func (s ProjectStore) GetProject(tx *dbr.Tx, orgSlug, projectSlug string) (models.Project, error) {
	project := models.Project{}
	_, err := tx.SelectBySql(`
            select p.*
                from sentry_project p
                    join sentry_organization o on p.organization_id = o.id
            where o.slug = ? and p.slug = ?`,
		orgSlug, projectSlug).
		LoadStructs(&project)
	// TODO err will still be equal to nil if project not found
	if err != nil {
		return project, errors.Wrap(err, "can not get project")
	}
	return project, nil
}

func (s ProjectStore) GetEnvironments(tx *dbr.Tx, projectID int) ([]models.Environment, error) {
	environments := []models.Environment{}
	_, err := tx.SelectBySql(`
		select se.*
			from sentry_environment se
		where se.project_id = ?
        order by se.name`,
		projectID).
		LoadStructs(&environments)
	if err != nil {
		return nil, errors.Wrap(err, "can not read project environments")
	}
	return environments, nil
}

func (s ProjectStore) GetTags(tx *dbr.Tx, projectID int) ([]*models.TagKey, error) {
	tags := []*models.TagKey{}
	_, err := tx.SelectBySql(`
		select fk.*
			from sentry_filterkey fk
		where fk.project_id = ? and fk.status = ?`,
		projectID, models.TagKeyStatusVisible).
		LoadStructs(&tags)
	if err != nil {
		return nil, errors.Wrap(err, "can not read project tags")
	}
	// TODO tag.Key must be processed -> TagKey.get_standardized_key(tag_key.key)
	for _, tag := range tags {
		tag.PostGet()
	}
	return tags, nil
}

func (s ProjectStore) SaveProject(tx *dbr.Tx, project models.Project) error {
	_, err := tx.InsertInto("sentry_project").
		Columns("id", "team_id", "organization_id", "name", "slug",
			"public", "status", "first_event", "date_added").
		Record(project).
		Exec()
	return errors.Wrap(err, "failed to save project")
}

func (s ProjectStore) SaveEnvironment(tx *dbr.Tx, environment models.Environment) error {
	_, err := tx.InsertInto("sentry_environment").
		Columns("id", "project_id", "name", "date_added").
		Record(environment).
		Exec()
	return errors.Wrap(err, "failed to save project environment")
}

func (s ProjectStore) SaveTags(tx *dbr.Tx, tags ...*models.TagKey) error {
	query := tx.InsertInto("sentry_filterkey").
		Columns("id", "project_id", "key", "values_seen", "label", "status")
	for _, tag := range tags {
		query = query.Record(tag)
	}
	// TODO can we just ignore rv / sql.Result?
	_, err := query.Exec()
	return errors.Wrap(err, "failed to save project tags")
}

func (s ProjectStore) SaveSearches(tx *dbr.Tx, searches ...models.SavedSearch) error {
	query := tx.InsertInto("sentry_savedsearch").
		Columns("id", "project_id", "name", "query", "date_added", "is_default")
	for _, search := range searches {
		query = query.Record(search)
	}
	_, err := query.Exec()
	return errors.Wrap(err, "failed to save project searches")
}

func (s ProjectStore) SaveEventGroup(tx *dbr.Tx, group models.Group) error {
	_, err := tx.InsertInto("sentry_groupedmessage").
		Columns(
			"id", "logger", "level", "message", "view", "status", "times_seen",
			"last_seen", "first_seen", "data", "score", "project_id",
			"time_spent_total", "time_spent_count", "resolved_at", "active_at",
			"is_public", "platform", "num_comments", "first_release_id",
			"short_id").
		Record(group).
		Exec()
	return errors.Wrap(err, "failed to save groups of project issues")
}
