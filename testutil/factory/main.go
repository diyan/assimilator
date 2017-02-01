package factory

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"

	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

type TestFactory struct {
	suite            suite.Suite
	tx               *dbr.Tx
	SaveOrganization func(org models.Organization)
	SaveProject      func(project models.Project)
	SaveTags         func(tags ...*models.TagKey)
}

// TODO Use *testing.T for better re-use
func New(suite suite.Suite, server *echo.Echo) TestFactory {
	noError := suite.Require().NoError
	ctx := server.NewContext(nil, nil)
	tx, err := db.GetTx(ctx)
	noError(err)
	tf := TestFactory{
		suite: suite,
		tx:    tx,
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

func (tf TestFactory) Reset() {
	err := tf.tx.Rollback()
	tf.suite.Require().NoError(err)
}
