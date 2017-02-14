package api_test

import (
	"testing"

	"github.com/diyan/assimilator/migrations"
	"github.com/diyan/assimilator/testutil/factory"
	"github.com/diyan/assimilator/testutil/testclient"
	"github.com/diyan/assimilator/web"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	Client  *gorequest.SuperAgent
	App     *echo.Echo
	Factory factory.TestFactory
}

func (t *testSuite) SetupSuite() {
	// TODO check what is faster - re-create db or drop all tables?
	// select 'drop table "' || tablename || '" cascade;'
	// from pg_tables where schemaname = 'sentry_ci';

	// TODO remove duplicated code
	noError := require.New(t.T()).NoError
	conn, err := dbr.Open("postgres", "postgres://postgres@localhost/postgres?sslmode=disable", nil)
	noError(errors.Wrap(err, "failed to init db connection"))
	// dbr.Open calls sql.Open which returns err == nil even if there is no db connection,
	//   so it is required to explicitly ping the database
	err = conn.Ping()
	noError(errors.Wrap(err, "failed to ping db"))
	sess := conn.NewSession(nil)
	// Force drop db while others may be connected
	_, err = sess.Exec(`
		select pg_terminate_backend(pid) 
		from pg_stat_activity 
		where datname = 'sentry_ci';`)
	noError(err)
	_, err = sess.Exec("drop database if exists sentry_ci;")
	noError(err)
	_, err = sess.Exec("create database sentry_ci;")
	noError(err)
	noError(migrations.UpgradeDB(factory.MakeAppConfig().DatabaseURL))
}

func (t *testSuite) SetupTest() {
	t.App = web.GetApp(factory.MakeAppConfig())
	// TODO TestFactory does a side effect which is used by TestClient
	//  make this code more explicit
	t.Factory = factory.New(t.T(), t.App)
	t.Client = testclient.New(t.App)
}

func (t *testSuite) TearDownTest() {
	t.Factory.Reset()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
