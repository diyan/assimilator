package api

import "github.com/labstack/echo"

func GroupDetailsGetEndpoint(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write([]byte(`{
    "seenBy": [],
    "pluginIssues": [],
    "lastSeen": "2017-01-25T12:51:01Z",
    "userReportCount": 0,
    "numComments": 0,
    "userCount": 0,
    "stats": {
        "30d": [
            [
                1484438400,
                0
            ],
            [
                1484524800,
                0
            ],
            [
                1484611200,
                0
            ],
            [
                1484697600,
                0
            ],
            [
                1484784000,
                0
            ],
            [
                1484870400,
                0
            ],
            [
                1484956800,
                0
            ],
            [
                1485043200,
                0
            ],
            [
                1485129600,
                0
            ],
            [
                1485216000,
                0
            ],
            [
                1485302400,
                0
            ],
            [
                1485388800,
                0
            ],
            [
                1485475200,
                0
            ],
            [
                1485561600,
                0
            ],
            [
                1485648000,
                0
            ],
            [
                1485734400,
                0
            ],
            [
                1485820800,
                0
            ],
            [
                1485907200,
                0
            ],
            [
                1485993600,
                0
            ],
            [
                1486080000,
                0
            ],
            [
                1486166400,
                0
            ],
            [
                1486252800,
                0
            ],
            [
                1486339200,
                0
            ],
            [
                1486425600,
                0
            ],
            [
                1486512000,
                0
            ],
            [
                1486598400,
                0
            ],
            [
                1486684800,
                0
            ],
            [
                1486771200,
                0
            ],
            [
                1486857600,
                0
            ],
            [
                1486944000,
                0
            ],
            [
                1487030400,
                0
            ]
        ],
        "24h": [
            [
                1487001600,
                0
            ],
            [
                1487005200,
                0
            ],
            [
                1487008800,
                0
            ],
            [
                1487012400,
                0
            ],
            [
                1487016000,
                0
            ],
            [
                1487019600,
                0
            ],
            [
                1487023200,
                0
            ],
            [
                1487026800,
                0
            ],
            [
                1487030400,
                0
            ],
            [
                1487034000,
                0
            ],
            [
                1487037600,
                0
            ],
            [
                1487041200,
                0
            ],
            [
                1487044800,
                0
            ],
            [
                1487048400,
                0
            ],
            [
                1487052000,
                0
            ],
            [
                1487055600,
                0
            ],
            [
                1487059200,
                0
            ],
            [
                1487062800,
                0
            ],
            [
                1487066400,
                0
            ],
            [
                1487070000,
                0
            ],
            [
                1487073600,
                0
            ],
            [
                1487077200,
                0
            ],
            [
                1487080800,
                0
            ],
            [
                1487084400,
                0
            ],
            [
                1487088000,
                0
            ]
        ]
    },
    "culprit": "__main__ in <module>",
    "title": "This is a test message generated using raven test",
    "id": "1",
    "assignedTo": null,
    "participants": [],
    "logger": null,
    "type": "default",
    "annotations": [],
    "metadata": {
        "title": "This is a test message generated using raven test"
    },
    "status": "unresolved",
    "pluginActions": [],
    "tags": [],
    "subscriptionDetails": null,
    "isPublic": false,
    "hasSeen": false,
    "firstRelease": null,
    "shortId": "ACME-1",
    "shareId": "322e31",
    "firstSeen": "2017-01-25T12:51:01Z",
    "count": "1",
    "permalink": "http://localhost:9000/acme-team/acme/issues/1/",
    "level": "info",
    "isSubscribed": true,
    "isBookmarked": false,
    "project": {
        "name": "ACME",
        "slug": "acme"
    },
    "lastRelease": null,
    "activity": [
        {
            "data": {},
            "dateCreated": "2017-01-25T12:51:01Z",
            "type": "first_seen",
            "id": "None",
            "user": null
        }
    ],
    "statusDetails": {}
    }`))
	return nil
}
