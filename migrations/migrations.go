package migrations

import (
	"github.com/GeertJohan/go.rice"
	"github.com/mattes/migrate"
	"github.com/pkg/errors"
	// Enable PostgreSQL driver for migration tool
	_ "github.com/mattes/migrate/database/postgres"
	// Enable File source for migration tool
	_ "github.com/mattes/migrate/source/file"
	// Enable go.rice source for migration tool
	"github.com/diyan/assimilator/migrations/source"
)

func UpgradeDB(databaseURL string) error {
	// TODO Who is responsible to run `create database sentry_ci` statement?
	box, err := rice.FindBox("postgres")
	if err != nil {
		return errors.Wrap(err, "can not find db migrations")
	}
	sourceDriver, err := source.WithInstance(box)
	if err != nil {
		return errors.Wrap(err, "can not init source driver for db migrations")
	}
	m, err := migrate.NewWithSourceInstance("go.rice", sourceDriver, databaseURL)
	if err != nil {
		return errors.Wrap(err, "failed to upgrade database schema")
	}
	if err := m.Up(); err != nil {
		return errors.Wrap(err, "failed to upgrade database schema")
	}
	return nil
}
