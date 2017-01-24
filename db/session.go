package db

import (
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func GetSession(c echo.Context) (*dbr.Session, error) {
	conn, err := dbr.Open("postgres", "postgres://sentry:RucLUS8A@localhost/sentry?sslmode=disable", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init db connection")
	}
	// dbr.Open calls sql.Open which returns err == nil even if there is no db connection,
	//   so it is required to explicitly ping the database
	err = conn.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}
	return conn.NewSession(nil), nil
}
