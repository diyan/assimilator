package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

// TODO consider move store sources from ./db/store to ./store
type OrganizationStore struct {
}

func NewOrganizationStore() OrganizationStore {
	return OrganizationStore{}
}

func (s OrganizationStore) GetOrganization(tx *dbr.Tx, orgSlug string) (*models.Organization, error) {
	org := &models.Organization{}
	_, err := tx.SelectBySql(`
            select o.*
                from sentry_organization o
            where o.slug = ?`,
		orgSlug).
		LoadStructs(org)
	// TODO err will still be equal to nil if organization not found
	if err != nil {
		return org, errors.Wrap(err, "can not get organization")
	}
	return org, nil
}

func (s OrganizationStore) SaveOrganization(tx *dbr.Tx, org models.Organization) error {
	_, err := tx.InsertInto("sentry_organization").
		Columns("id", "name", "slug", "status", "flags", "default_role", "date_added").
		Record(org).
		Exec()
	return errors.Wrap(err, "failed to save organization")
}

func (s OrganizationStore) SaveOrganizationMember(
	tx *dbr.Tx, orgMember models.OrganizationMember) error {
	_, err := tx.InsertInto("sentry_organizationmember").
		Columns("id", "organization_id", "user_id", "type", "date_added", "email", "has_global_access", "flags", "role", "token").
		Record(orgMember).
		Exec()
	return errors.Wrap(err, "failed to save organization member")
}
