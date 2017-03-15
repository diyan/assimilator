package migrations

import (
	"github.com/diyan/assimilator/db/migrations/source"

	"github.com/GeertJohan/go.rice"
	"github.com/mattes/migrate"
	"github.com/pkg/errors"
	// Enable PostgreSQL driver for migration tool
	_ "github.com/mattes/migrate/database/postgres"
)

func UpgradeDB(databaseURL string) error {
	// TODO Who should run `create database sentry_ci` statement?
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
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to upgrade database schema")
	}
	return nil
}
