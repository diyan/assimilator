package api_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationIndex_Get(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())

	res, bodyStr, errs := client.Get("http://example.com/api/0/organizations/").End()
	assert.Nil(t, errs)
	// TODO `isEarlyAdopter: false` is expected in the response
	assert.JSONEq(t, `[{
            "id": "1",
            "name": "ACME-Team",
            "slug": "acme-team",
            "dateCreated": "2999-01-01T00:00:00Z"
        }]`,
		bodyStr)
	assert.Equal(t, 200, res.StatusCode)
}

func TestOrganizationIndex_Get_MemberOnly(t *testing.T) {
	t.Skip("Not yet implemented")
}
