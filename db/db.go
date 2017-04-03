package db

import (
	"github.com/diyan/assimilator/conf"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type TxMakerFunc func() (*dbr.Tx, error)

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
