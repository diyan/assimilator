package db

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func GetSession(c echo.Context) (*dbr.Session, error) {
	conn, err := dbr.Open("postgres", "postgres://sentry:RucLUS8A@localhost/sentry?sslmode=disable", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open db connection")
	}
	return conn.NewSession(nil), nil
}
