package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationDetails_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)

	factory.SaveUser(factory.MakeUser())
	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveOrganizationMember(factory.MakeOrganizationMember())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveTeamMember(factory.MakeTeamMember())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Get("http://example.com/api/0/organizations/acme-team").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	// TODO in response teams.1.project.1.features should be equal to ["quotas"]
	assert.JSONEq(t, `{
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
}
