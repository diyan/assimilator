package migrations

import (

	// Enable PostgreSQL driver for migration tool
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	"github.com/pkg/errors"
)

func UpgradeDB(databaseURL string) error {
	// TODO Who is responsible to run `create database sentry_ci` statement?
	// TODO consider use async version and optimize app's startup time
	allErrors, ok := migrate.UpSync(databaseURL, "/home/alexey/go/src/github.com/diyan/assimilator/migrations/postgres")
	if !ok {
		// TODO handle error, do not just return the first one
		// pp.Print(allErrors)
		return errors.Wrapf(allErrors[0], "failed to upgrade database schema")
	}
	return nil
}
