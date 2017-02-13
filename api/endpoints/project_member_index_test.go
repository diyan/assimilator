package api_test

func (t *testSuite) TestProjectMemberIndex_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveOrganizationMember(t.Factory.MakeOrganizationMember())
	t.Factory.SaveProject(t.Factory.MakeProject())
	t.Factory.SaveUser(t.Factory.MakeUser())

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/projects/acme-team/acme/members/").End()
	t.Nil(errs)
	// TODO in the response avatarUrl should be not empty, for ex https://secure.gravatar.com/avatar/01bce7702975191fdc402565bd1045a8?s=32&d=mm
	t.JSONEq(`[{
            "id": "1",
            "username": "admin",
            "name": "admin@example.com",
            "email": "admin@example.com",            
            "avatarUrl": "",
            "options": {
                "timezone": "UTC",
                "stacktraceOrder": "default",
                "language": "en",
                "clock24Hours": false
            }
        }]`,
		bodyStr)
	t.Equal(200, res.StatusCode)
}
