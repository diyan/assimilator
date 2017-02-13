package api_test

func (t *testSuite) TestProjectTags_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveTags(t.Factory.MakeTags()...)

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/tags/").End()
	t.Nil(errs)
	t.JSONEq(`[{
			"id": "1",
			"key": "server_name",
			"uniqueValues": 0,
			"name": null
		},
		{
			"id": "2",
			"key": "level",
			"uniqueValues": 0,
			"name": null
		}]`,
		bodyStr)
	t.Equal(200, res.StatusCode)
}

func (t *testSuite) TestProjectTags_Post() {
	t.T().Skip("Not yet implemented")
}
