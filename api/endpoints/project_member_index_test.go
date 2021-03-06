package api_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestProjectMemberIndex_Get(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveUser(factory.MakeUser())
	factory.SaveOrganizationMember(factory.MakeOrganizationMember())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Get("http://example.com/api/0/projects/acme-team/acme/members/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	// TODO in the response avatarUrl should be not empty, for ex https://secure.gravatar.com/avatar/01bce7702975191fdc402565bd1045a8?s=32&d=mm
	assert.JSONEq(t, `[{
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
}
