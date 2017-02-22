package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectTags_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)
	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveProject(factory.MakeProject())
	factory.SaveTags(factory.MakeTags()...)

	res, bodyStr, _ := client.Get("http://example.com/api/0/projects/acme-team/acme/tags/").End()
	assert.Equal(t, 200, res.StatusCode)
	// TODO in the response uniqueValues should be equal to 1
	assert.JSONEq(t, `[{
			"id": "1",
			"key": "server_name",
			"uniqueValues": 0,
			"name": "Server"
		},
		{
			"id": "2",
			"key": "level",
			"uniqueValues": 0,
			"name": "Level"
		}]`,
		bodyStr)
}

func TestProjectTags_Post(t *testing.T) {
	t.Skip("Not yet implemented")
}
