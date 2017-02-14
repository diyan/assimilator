package api_test

func (t *testSuite) TestOrganizationIndex_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/organizations/").End()
	t.Nil(errs)
	// TODO `isEarlyAdopter: false` is expected in the response
	t.JSONEq(`[{
            "id": "1",
            "name": "ACME-Team",
            "slug": "acme-team",
            "dateCreated": "2999-01-01T00:00:00Z"
        }]`,
		bodyStr)
	t.Equal(200, res.StatusCode)
}

func (t *testSuite) TestOrganizationIndex_Get_MemberOnly() {
	t.T().Skip("Not yet implemented")
}
