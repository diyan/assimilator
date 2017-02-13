package migrations

import (
	"path/filepath"
	// Enable PostgreSQL driver for migration tool
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	"github.com/pkg/errors"
)

func UpgradeDB(databaseURL string) error {
	// TODO Who is responsible to run `create database sentry_ci` statement?
	// TODO consider use async version and optimize app's startup time
	// TODO this relative path works only for single test package, use go.rice instead
	migrationsPath, err := filepath.Abs("../../migrations/postgres")
	if err != nil {
		return errors.Wrap(err, "can not get absolute path for db migrations")
	}
	allErrors, ok := migrate.UpSync(databaseURL, migrationsPath)
	if !ok {
		// TODO handle error, do not just return the first one
		// pp.Print(allErrors)
		return errors.Wrap(allErrors[0], "failed to upgrade database schema")
	}
	return nil
}
