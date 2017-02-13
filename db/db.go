package db

import (
	"github.com/diyan/assimilator/conf"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// FromE returns DB transaction associated with echo's Context
func FromE(c echo.Context) (*dbr.Tx, error) {
	if tx, ok := c.Get("dbr.Tx").(*dbr.Tx); ok {
		return tx, nil
	}
	tx, err := New(conf.FromE(c))
	if err != nil {
		return nil, err
	}
	ToE(c, tx)
	return tx, nil
}

// ToE save DB transaction into provided echo's Context
func ToE(c echo.Context, tx *dbr.Tx) {
	c.Set("dbr.Tx", tx)
}

// New starts new DB transactions
func New(c conf.Config) (*dbr.Tx, error) {
	conn, err := dbr.Open("postgres", c.DatabaseURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init db connection")
	}
	// dbr.Open calls sql.Open which returns err == nil even if there is no db connection,
	//   so it is required to explicitly ping the database
	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}
	sess := conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "can not start db transaction")
	}
	return tx, nil
}
