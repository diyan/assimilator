package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type OrganizationStore struct {
	tx *dbr.Tx
}

func NewOrganizationStore(tx *dbr.Tx) OrganizationStore {
	return OrganizationStore{tx: tx}
}

func (s OrganizationStore) SaveOrganization(org models.Organization) error {
	_, err := s.tx.InsertInto("sentry_organization").
		Columns("id", "name", "slug", "status", "flags", "default_role", "date_added").
		Record(org).
		Exec()
	return errors.Wrapf(err, "failed to save organization")
}
