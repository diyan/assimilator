package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectGroupIndex_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)
	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveProject(factory.MakeProject())
	factory.SaveEventGroup(factory.MakeEventGroup())

	res, bodyStr, errs := client.Get("http://example.com:80/api/0/projects/acme-team/acme/issues/?query=is%3Aunresolved&limit=25&sort=date&statsPeriod=24h&shortIdLookup=1").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	// TODO current implementation is good enough to show event group on the http://localhost:3000/acme-team/acme/ page
	// but following attribs are still missing:
	/*  "count": "1",
	    "isSubscribed": true,
		"stats": {
	            "24h": [
	                [1487005200, 0],
	                [1487008800, 0],
	                [1487012400, 0],
	                [1487016000, 0],
	                [1487019600, 0],
	                [1487023200, 0],
	                [1487026800, 0],
	                [1487030400, 0],
	                [1487034000, 0],
	                [1487037600, 0],
	                [1487041200, 0],
	                [1487044800, 0],
	                [1487048400, 0],
	                [1487052000, 0],
	                [1487055600, 0],
	                [1487059200, 0],
	                [1487062800, 0],
	                [1487066400, 0],
	                [1487070000, 0],
	                [1487073600, 0],
	                [1487077200, 0],
	                [1487080800, 0],
	                [1487084400, 0],
	                [1487088000, 0]
	            ]
	        }
	*/
	assert.JSONEq(t, `[
	    {
	        "id": "1",
	        "shareId": "312e31",
	        "shortId": "ACME-1",
	        "count": 0,
	        "userCount": 0,
	        "title": "This is a test message generated using `+"``raven test``"+` __main__ in <module>",
	        "culprit": "__main__ in <module>",
	        "permalink": "http://example.com:80/acme-team/acme/issues/1/",
	        "firstSeen": "2999-01-01T00:00:00Z",
			"lastSeen": "2999-01-01T00:00:00Z",
	        "logger": null,
	        "level": "info",
	        "status": "unresolved",
	        "statusDetails": {},
	        "isPublic": false,
	        "project": {
	            "name": "ACME",
	            "slug": "acme"
	        },
	        "type": "default",
			"metadata": {
	            "title": "This is a test message generated using `+"``raven test``"+` __main__ in <module>"
	        },
			"numComments": 0,
	        "assignedTo": null,
	        "isBookmarked": false,			
	        "isSubscribed": false,
	        "subscriptionDetails": null,
	        "hasSeen": false,			
	        "annotations": [],
	        "stats": {
				"24h": null
			}
	    }
	]`, bodyStr)
}
