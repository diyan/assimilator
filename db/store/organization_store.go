package store

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// TODO consider move store sources from ./db/store to ./store
type OrganizationStore struct {
	c echo.Context
}

func NewOrganizationStore(c echo.Context) OrganizationStore {
	return OrganizationStore{c: c}
}

func (s OrganizationStore) SaveOrganization(org models.Organization) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save organization")
	}
	_, err = db.InsertInto("sentry_organization").
		Columns("id", "name", "slug", "status", "flags", "default_role", "date_added").
		Record(org).
		Exec()
	return errors.Wrap(err, "failed to save organization")
}

func (s OrganizationStore) SaveOrganizationMember(
	orgMember models.OrganizationMember) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save organization member")
	}
	_, err = db.InsertInto("sentry_organizationmember").
		Columns("id", "organization_id", "user_id", "type", "date_added", "email", "has_global_access", "flags", "role", "token").
		Record(orgMember).
		Exec()
	return errors.Wrap(err, "failed to save organization member")
}
