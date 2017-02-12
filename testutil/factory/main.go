package factory

import (
	"testing"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
)

type TestFactory struct {
	t                *testing.T
	tx               *dbr.Tx
	Reset            func()
	SaveOrganization func(org models.Organization)
	SaveProject      func(project models.Project)
	SaveTags         func(tags ...*models.TagKey)
}

func New(t *testing.T, server *echo.Echo) TestFactory {
	noError := require.New(t).NoError
	ctx := server.NewContext(nil, nil)
	// TODO remove hack
	ctx.Set("conf.Config", MakeAppConfig())
	tx, err := db.GetTx(ctx)
	noError(err)
	tf := TestFactory{
		t:  t,
		tx: tx,
	}
	tf.Reset = func() {
		err := tf.tx.Rollback()
		noError(err)
	}

	// TODO Tricky implementation. Mock *dbr.Tx in the test Echo instance
	server.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dbr.Tx", tx)
			return next(c)
		}
	})

	orgStore := store.NewOrganizationStore(ctx)
	projectStore := store.NewProjectStore(ctx)
	tf.SaveOrganization = func(org models.Organization) {
		noError(orgStore.SaveOrganization(org))
	}
	tf.SaveProject = func(project models.Project) {
		noError(projectStore.SaveProject(project))
	}
	tf.SaveTags = func(tags ...*models.TagKey) {
		noError(projectStore.SaveTags(tags...))
	}
	return tf
}
