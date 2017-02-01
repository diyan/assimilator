package api_test

import (
	"net/http"
	"time"

	"net/http/httptest"
	"testing"

	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/migrations"
	"github.com/diyan/assimilator/models"
	"github.com/diyan/assimilator/web"
	"github.com/labstack/echo"
	"github.com/pkg/errors"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/gocraft/dbr"
)

type testSuite struct {
	suite.Suite
	*require.Assertions
	HttpRecorder *httptest.ResponseRecorder
	Client       *EchoTestClient
	App          *echo.Echo
	Factory      TestFactory
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
	//t.HttpRecorder = httptest.NewRecorder()
	t.App = web.GetApp()
	t.Factory = NewTestFactory(t.Suite, t.App)
	t.Client = NewEchoTestClient(t.Suite, t.App)
}

func (t *testSuite) TearDownTest() {
	t.Factory.Reset()
}

type TestFactory struct {
	suite            suite.Suite
	tx               *dbr.Tx
	SaveOrganization func(org models.Organization)
	SaveProject      func(project models.Project)
	SaveTags         func(tags ...*models.TagKey)
}

func NewTestFactory(suite suite.Suite, server *echo.Echo) TestFactory {
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

var time_of_2999_01_01__00_00_00 = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.UTC)

func (tf TestFactory) MakeTags() []*models.TagKey {
	tag1 := models.TagKey{
		ID:        1,
		ProjectID: 1,
		Key:       "server_name",
	}
	tag2 := tag1
	tag2.ID = 2
	tag2.Key = "level"
	return []*models.TagKey{&tag1, &tag2}
}

func (tf TestFactory) MakeOrganization() models.Organization {
	return models.Organization{
		ID:          1,
		Name:        "ACME-Team",
		Slug:        "acme-team",
		Status:      models.OrganizationStatusVisible,
		Flags:       1, // TODO Introduce constants
		DefaultRole: "member",
		DateCreated: time_of_2999_01_01__00_00_00,
	}
}

func (tf TestFactory) MakeProject() models.Project {
	return models.Project{
		ID:             1,
		TeamID:         1,
		OrganizationID: 1,
		Name:           "ACME",
		Slug:           "acme",
		Public:         false,
		Status:         models.ProjectStatusVisible,
		FirstEvent:     time_of_2999_01_01__00_00_00,
		DateCreated:    time_of_2999_01_01__00_00_00,
	}
}

// TODO Move test client into separate module
type EchoTestClient struct {
	server   *echo.Echo
	recorder *httptest.ResponseRecorder
	suite    suite.Suite
}

// TODO keep the TestClient generic if possible
func NewEchoTestClient(suite suite.Suite, server *echo.Echo) *EchoTestClient {
	return &EchoTestClient{
		server: server,
		suite:  suite,
	}
}

func (c *EchoTestClient) Get(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	c.suite.NoError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}

func (c *EchoTestClient) Delete(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", url, nil)
	c.suite.NoError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

// TODO setup project, organization, etc using text fixtures
func (t *testSuite) TestProjectTags_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveTags(t.Factory.MakeTags()...)
	// TODO move TestFactory to the factory.go
	// TODO Group the API to improve code completion
	// org := OrganizationFactory.MustCreate().(*models.Organization)
	// t.Factory.Organization.MustCreate().(*models.Organization)
	// Q: code below is easy to implement, but how to make it type safe?

	rr := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/tags/")
	t.Equal(200, rr.Code)
	// TODO result below is from read db but we should use test db
	// TODO Investigate why GoConvey crashing if t.JSONEq is false
	t.JSONEq(`[{
			"id": "1",
			"key": "key-1",
			"uniqueValues": 0,
			"name": null
		},
		{
			"id": "2",
			"key": "key-2",
			"uniqueValues": 0,
			"name": null
		}]`,
		rr.Body.String())

	// TODO Can we pass t.Tx to the TagKeyFactory.MustCreateWithOption ?
	// TODO Try to develop API like this - t.Factory.TagKey.MustCreate()
}

func (t *testSuite) TestProjectTags_Post() {

}
