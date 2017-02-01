package echotest

import (
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
)

// TODO Move test client into separate module
type TestClient struct {
	server   *echo.Echo
	recorder *httptest.ResponseRecorder
	suite    suite.Suite
}

// TODO keep the TestClient generic if possible
// TODO Use *testing.T for better re-use
func NewClient(suite suite.Suite, server *echo.Echo) *TestClient {
	return &TestClient{
		server: server,
		suite:  suite,
	}
}

func (c *TestClient) Get(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	c.suite.NoError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}

func (c *TestClient) Delete(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", url, nil)
	c.suite.NoError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}
