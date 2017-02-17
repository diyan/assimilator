package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectSearches_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)
	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveProject(factory.MakeProject())
	factory.SaveProjectSearches(factory.MakeProjectSearches()...)

	res, bodyStr, errs := client.Get("http://example.com/api/0/projects/acme-team/acme/searches/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, `[
		    {
				"id": "3",
				"name": "Assigned To Me",
				"query": "is:unresolved assigned:me",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			},
			{
				"id": "4",
				"name": "My Bookmarks",
				"query": "is:unresolved bookmarks:me",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			},
			{
				"id": "2",
				"name": "Needs Triage",
				"query": "is:unresolved is:unassigned",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			},
			{
				"id": "5",
				"name": "New Today",
				"query": "is:unresolved age:-24h",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			},
			{
				"id": "1",
				"name": "Unresolved Issues",
				"query": "is:unresolved",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": true
			}
		]`,
		bodyStr)
}
