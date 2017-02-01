package echotest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestClient struct {
	server   http.Handler
	recorder *httptest.ResponseRecorder
	t        *testing.T
}

func NewClient(t *testing.T, server http.Handler) *TestClient {
	return &TestClient{
		server: server,
		t:      t,
	}
}

func (c TestClient) noError(err error, msgAndArgs ...interface{}) {
	require.NoError(c.t, err, msgAndArgs)
}

func (c *TestClient) Get(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	c.noError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}

func (c *TestClient) Delete(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", url, nil)
	c.noError(err)
	c.server.ServeHTTP(recorder, req)
	return recorder
}
