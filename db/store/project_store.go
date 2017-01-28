package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type ProjectStore struct {
	tx *dbr.Tx
}

func NewProjectStore(tx *dbr.Tx) ProjectStore {
	return ProjectStore{tx: tx}
}

func (s ProjectStore) SaveProject(project models.Project) error {
	_, err := s.tx.InsertInto("sentry_project").
		Columns("id", "team_id", "organization_id", "name", "slug",
			"public", "status", "first_event", "date_added").
		Record(project).
		Exec()
	return errors.Wrapf(err, "failed to save project")
}

func (s ProjectStore) SaveTags(tags ...*models.TagKey) error {
	q := s.tx.InsertInto("sentry_filterkey").
		Columns("id", "project_id", "key", "values_seen", "label", "status")
	for _, tag := range tags {
		q = q.Record(tag)
	}
	// TODO can we just ignore rv / sql.Result?
	_, err := q.Exec()
	return errors.Wrapf(err, "failed to save project tags")
}
