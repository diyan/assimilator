package api_test

func (t *testSuite) TestProjectEnvironment_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveEnvironment(t.Factory.MakeEnvironment())

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/environments/").End()
	t.Nil(errs)
	t.Equal(200, res.StatusCode)
	t.JSONEq(`[{
			"id": "1",
			"name": ""
		}]`,
		bodyStr)
}
