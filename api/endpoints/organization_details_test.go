package api_test

func (t *testSuite) TestOrganizationDetails_Get() {
	t.Factory.SaveOrganization(t.Factory.MakeOrganization())
	t.Factory.SaveOrganizationMember(t.Factory.MakeOrganizationMember())
	t.Factory.SaveTeam(t.Factory.MakeTeam())
	t.Factory.SaveTeamMember(t.Factory.MakeTeamMember())
	t.Factory.SaveProject(t.Factory.MakeProject())

	res, bodyStr, errs := t.Client.Get("http://example.com/api/0/organizations/acme-team").End()
	t.Nil(errs)
	// TODO in response teams.1.project.1.features should be equal to ["quotas"]
	t.JSONEq(`{
            "id": "1",        
            "slug": "acme-team",
            "name": "ACME-Team",
            "dateCreated": "2999-01-01T00:00:00Z",
            "pendingAccessRequests": 0,
            "teams": [{
                "id": "1",
                "slug": "acme-team",
                "name": "ACME-Team",
                "hasAccess": true,
                "isPending": false,
                "dateCreated": "2999-01-01T00:00:00Z",
                "isMember": true,
                "projects": [{
                    "id": "1",
                    "slug": "acme",
                    "name": "ACME",
                    "isPublic": false,
                    "dateCreated": "2999-01-01T00:00:00Z",
                    "firstEvent": "2999-01-01T00:00:00Z",
                    "features": null
                 }]
            }],
            "features": ["sso", "open-membership"],            
            "quota": {
                "projectLimit": 100,
                "maxRate": 0
            },
            "access": [
                "org:write",
                "member:write",
                "project:read",
                "org:delete",
                "org:read",
                "event:read",
                "team:delete",
                "member:delete",
                "event:delete",
                "member:read",
                "project:write",
                "event:write",
                "project:delete",
                "team:read",
                "team:write"
            ]
        }`,
		bodyStr)
	t.Equal(200, res.StatusCode)
}
