package api_test

import (
	"net/http/httptest"
	"testing"

	"github.com/diyan/assimilator/migrations"
	"github.com/diyan/assimilator/testutil/echotest"
	"github.com/diyan/assimilator/testutil/factory"
	"github.com/diyan/assimilator/web"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	*require.Assertions
	HttpRecorder *httptest.ResponseRecorder
	Client       *echotest.TestClient
	App          *echo.Echo
	Factory      factory.TestFactory
}

// SetT overrides assert.Assertions with require.Assertions.
func (suite *testSuite) SetT(t *testing.T) {
	suite.Suite.SetT(t)
	suite.Assertions = require.New(t)
}

func (t *testSuite) SetupSuite() {
	// TODO check what is faster - re-create db or drop all tables?
	// select 'drop table "' || tablename || '" cascade;'
	// from pg_tables where schemaname = 'sentry_ci';

	// TODO remove duplicated code
	conn, err := dbr.Open("postgres", "postgres://sentry:RucLUS8A@localhost/postgres?sslmode=disable", nil)
	t.NoError(errors.Wrap(err, "failed to init db connection"))
	// dbr.Open calls sql.Open which returns err == nil even if there is no db connection,
	//   so it is required to explicitly ping the database
	err = conn.Ping()
	t.NoError(errors.Wrap(err, "failed to ping db"))
	sess := conn.NewSession(nil)
	// Force drop db while others may be connected
	_, err = sess.Exec(`
		select pg_terminate_backend(pid) 
		from pg_stat_activity 
		where datname = 'sentry_ci';`)
	t.NoError(err)
	_, err = sess.Exec("drop database if exists sentry_ci;")
	t.NoError(err)
	_, err = sess.Exec("create database sentry_ci;")
	t.NoError(err)
	migrations.UpgradeDB()
}

func (t *testSuite) TearDownSuite() {
	//fmt.Print("TearDownSuite")
}

// testify's suite.Suite calls following hooks on each test method execution:
// SetT, SetupTest, TearDownTest, SetT
// Question is why SetT func called twice?
func (t *testSuite) SetupTest() {
	t.App = web.GetApp()
	t.Factory = factory.New(t.T(), t.App)
	t.Client = echotest.NewClient(t.T(), t.App)
}

func (t *testSuite) TearDownTest() {
	t.Factory.Reset()
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
func (t *testSuite) TestProjectTags_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveTags(t.Factory.MakeTags()...)

	rr := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/tags/")
	t.Equal(200, rr.Code)
	// TODO Investigate why GoConvey crashing if t.JSONEq is false
	t.JSONEq(`[{
			"id": "1",
			"key": "server_name",
			"uniqueValues": 0,
			"name": null
		},
		{
			"id": "2",
			"key": "level",
			"uniqueValues": 0,
			"name": null
		}]`,
		rr.Body.String())
}

func (t *testSuite) TestProjectTags_Post() {

}
