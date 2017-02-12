package db

import (
	"github.com/diyan/assimilator/conf"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// GetTx returns DB transaction associated with current HTTP request
func GetTx(c echo.Context) (*dbr.Tx, error) {
	if tx, ok := c.Get("dbr.Tx").(*dbr.Tx); ok {
		return tx, nil
	}
	return NewTx(c)
}

// NewTx starts new DB transactions and associate it with current HTTP request
func NewTx(c echo.Context) (*dbr.Tx, error) {
	if sess, ok := c.Get("dbr.Session").(*dbr.Session); ok {
		tx, err := sess.Begin()
		if err != nil {
			return nil, errors.Wrap(err, "can not start db transaction")
		}
		c.Set("dbr.Tx", tx)
		return tx, nil
	}
	conn, err := dbr.Open("postgres", conf.FromEC(c).DatabaseURL, nil)
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
	c.Set("dbr.Session", sess)
	return NewTx(c)
}
