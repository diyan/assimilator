package api_test

func (t *testSuite) TestProjectSearches_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveProjectSearches(t.Factory.MakeProjectSearches()...)

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/searches/").End()
	t.Nil(errs)
	t.Equal(200, res.StatusCode)
	t.JSONEq(`[
		    {
				"id": "1",
				"name": "Unresolved Issues",
				"query": "is:unresolved",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": true
			},
			{
				"id": "2",
				"name": "Needs Triage",
				"query": "is:unresolved is:unassigned",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			},
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
				"id": "5",
				"name": "New Today",
				"query": "is:unresolved age:-24h",
				"dateCreated": "2999-01-01T00:00:00Z",
				"isDefault": false
			}
		]`,
		bodyStr)
}
