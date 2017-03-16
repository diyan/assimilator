package web_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestStaticJavaScript_Get(t *testing.T) {
	client, _ := fixture.Setup(t)
	defer fixture.TearDown(t)

	res, bodyStr, errs := client.Get("http://example.com/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.Contains(t, bodyStr, "Welcome to Sentry")
	assert.Contains(t, bodyStr, "https://github.com/getsentry/sentry")
	assert.Equal(t, res.Header.Get("Content-Type"), "application/x-javascript")
	assert.InDelta(t, 4000000, res.ContentLength, 1000000, "app.js should be ~4MB in size")
}

func TestStaticCSS_Get(t *testing.T) {
	client, _ := fixture.Setup(t)
	defer fixture.TearDown(t)

	res, bodyStr, errs := client.Get("http://example.com/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/sentry.css").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.Contains(t, bodyStr, "icon-sentry")
	assert.Contains(t, bodyStr, "sentry-loader")
	assert.Contains(t, res.Header.Get("Content-Type"), "text/css")
	assert.InDelta(t, 300000, res.ContentLength, 100000, "sentry.css should be ~300KB in size")
}
