package web_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestServerSideTemplateRenderer_Get(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())

	res, bodyStr, errs := client.Get("http://example.com//acme-team/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.Contains(t, res.Header.Get("Content-Type"), "text/html")
	assert.Contains(t, bodyStr, "<title>Sentry</title>", "Title should be rendered from sentry/layout.html template")
	assert.Contains(t, bodyStr, "Sentry.routes", "React routes should be rendered from sentry/bases/react.html")
	assert.InDelta(t, 3000, res.ContentLength, 1000, "server-side rendered page should be ~3KB in size")
}
