package factory

import (
	"testing"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
)

type TestFactory struct {
	t  *testing.T
	tx *dbr.Tx
}

func New(t *testing.T, server *echo.Echo, tx *dbr.Tx) TestFactory {
	tf := TestFactory{
		t:  t,
		tx: tx,
	}
	return tf
}

func (tf TestFactory) noError(err error, msgAndArgs ...interface{}) {
	require.New(tf.t).NoError(err, msgAndArgs...)
}

func (tf TestFactory) Reset() {
	err := tf.tx.Rollback()
	tf.noError(err)
}
