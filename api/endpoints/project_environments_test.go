package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectEnvironment_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)
	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveProject(factory.MakeProject())
	factory.SaveEnvironment(factory.MakeEnvironment())

	res, bodyStr, errs := client.Get("http://example.com/api/0/projects/acme-team/acme/environments/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, `[{
			"id": "1",
			"name": ""
		}]`,
		bodyStr)
}
