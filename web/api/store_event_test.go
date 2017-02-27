package api_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
)

func TestStoreEvent_Post(t *testing.T) {
	_, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	//res, bodyStr, errs := client.Get("http://example.com/api/0/issues/1/events/latest/").End()
	//assert.Nil(t, errs)
	//assert.Equal(t, 200, res.StatusCode)
	//assert.NotEmpty(t, bodyStr)
	//assert.JSONEq(t, `{
	//    }`,
	//	bodyStr)
}
