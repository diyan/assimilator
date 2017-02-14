package testclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/parnurzeal/gorequest"
)

func New(handler http.Handler) *gorequest.SuperAgent {
	mockTransport := mockTransport{
		handler: handler,
	}
	// Don't replace httpClient's Transport with SuperAgent's Transport
	gorequest.DisableTransportSwap = true
	httpAgent := gorequest.New()
	httpAgent.Client = &http.Client{Transport: mockTransport}
	return httpAgent
}

type mockTransport struct {
	handler http.Handler
}

func (mt mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	rr := httptest.NewRecorder()
	//rr.Body = &bytes.Buffer{}
	mt.handler.ServeHTTP(rr, req)
	return &http.Response{
		StatusCode:    rr.Code,
		Status:        http.StatusText(rr.Code),
		Header:        rr.HeaderMap,
		Body:          ioutil.NopCloser(rr.Body),
		ContentLength: int64(rr.Body.Len()),
		Request:       req,
	}, nil
}
