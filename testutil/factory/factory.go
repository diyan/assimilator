package factory

import (
	"testing"

	"github.com/diyan/assimilator/db"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
)

type TestFactory struct {
	t   *testing.T
	tx  *dbr.Tx
	ctx echo.Context
}

func New(t *testing.T, server *echo.Echo) TestFactory {
	tx, err := db.New(MakeAppConfig())
	require.NoError(t, err)
	MockDB(server, tx)
	ctx := server.NewContext(nil, nil)
	db.ToE(ctx, tx)
	tf := TestFactory{
		t:   t,
		tx:  tx,
		ctx: ctx,
	}
	return tf
}

func (tf TestFactory) noError(err error, msgAndArgs ...interface{}) {
	require.New(tf.t).NoError(err, msgAndArgs)
}

func (tf TestFactory) Reset() {
	err := tf.tx.Rollback()
	tf.noError(err)
}

// MockDB adds early middleware that mock DB transaction to the test Echo instance
// TODO consider move this to the db or db_test package
func MockDB(server *echo.Echo, tx *dbr.Tx) {
	server.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			db.ToE(c, tx)
			return next(c)
		}
	})
}
