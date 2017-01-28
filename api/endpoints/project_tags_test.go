package api_test

import (
	"fmt"
	"net/http"
	"time"

	"net/http/httptest"
	"testing"

	"github.com/bluele/factory-go/factory"
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
	Tx           *dbr.Tx
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
	tx, err := db.GetTx(t.App.NewContext(nil, nil))
	t.NoError(err)
	t.Tx = tx
	// Mock *dbr.Tx in the test App instance
	t.App.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("dbr.Tx", t.Tx)
			return next(c)
		}
	})
	t.Client = NewEchoTestClient(t.Suite, t.App)
}

func (t *testSuite) TearDownTest() {
	err := t.Tx.Rollback()
	t.NoError(err)
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

var time_of_2999_01_01__00_00_00 = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.UTC)

var OrganizationFactory = factory.NewFactory(
	// TODO Seems like we have to return a pointer. go-factory does not work with return by value
	&models.Organization{
		ID:          1,
		Name:        "ACME-Team",
		Slug:        "acme-team",
		Status:      models.OrganizationStatusVisible,
		Flags:       1, // What does this mean?
		DefaultRole: "member",
		DateCreated: time_of_2999_01_01__00_00_00,
	},
)

var ProjectFactory = factory.NewFactory(
	&models.Project{
		ID:             1,
		TeamID:         1,
		OrganizationID: 1,
		Name:           "ACME",
		Slug:           "acme",
		Public:         false,
		Status:         models.ProjectStatusVisible,
		FirstEvent:     time_of_2999_01_01__00_00_00,
		DateCreated:    time_of_2999_01_01__00_00_00,
	},
)

// TODO If wrong name will be passed to SeqInt the test will be not visible for GoConvey!
var TagKeyFactory = factory.NewFactory(
	&models.TagKey{
		ProjectID: 1,
	},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).SeqInt("Key", func(n int) (interface{}, error) {
	return fmt.Sprintf("key-%d", n), nil
})

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

// TODO setup project, organization, etc using text fixtures
func (t *testSuite) TestProjectTags_Get() {
	org := OrganizationFactory.MustCreate().(*models.Organization)
	project := ProjectFactory.MustCreate().(*models.Project)
	tagKey1 := TagKeyFactory.MustCreate().(*models.TagKey)
	tagKey2 := TagKeyFactory.MustCreate().(*models.TagKey)
	orgStore := store.NewOrganizationStore(t.Tx)
	err := orgStore.SaveOrganization(*org)
	t.NoError(err)
	projectStore := store.NewProjectStore(t.Tx)
	err = projectStore.SaveProject(*project)
	t.NoError(err)
	err = projectStore.SaveTags(tagKey1, tagKey2)
	t.NoError(err)

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
