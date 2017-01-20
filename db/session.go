package db

import (
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

func GetSession() (*dbr.Session, error) {
	conn, err := dbr.Open("postgres", "postgres://sentry:RucLUS8A@localhost/sentry?sslmode=disable", nil)
	if err != nil {
		return nil, errors.Wrap(err, "db connection failed")
	}
	return conn.NewSession(nil), nil
}
