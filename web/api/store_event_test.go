package api_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestStoreEvent_Post(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Post("http://example.com/api/1/store/").
		Send(`{
			"project": "1",
			"event_id": "44444444333322221111000000000000",			
			"logger": "javascript",
			"platform": "javascript",
			"release": "8.12.0",			
			"request": {
				"headers": {
					"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.76 Safari/537.36"
				},
				"url": "http://localhost:3000/acme-team/acme/issues/1/"
			},
			"user": {
				"id": 1,
				"email": "alexey.diyan@gmail.com",
				"ip_address": "192.169.100.1"
			},
			"culprit": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js:49:13510), <anonymous>",
			"extra": {
				"session:duration": 1145
			},
			"breadcrumbs": { "values": [
				{
					"timestamp": 1488129099.912, "type": "http", "category": "xhr",
					"data": { "method": "GET", "status_code": 200, "url": "/api/0/organizations/?member=1" }
				},
				{
					"timestamp": 1488129099.924, "type": "http", "category": "xhr",
					"data": { "method": "GET", "status_code": 200, "url": "/api/0/internal/health/" }
				},
				{
					"timestamp": 1488129099.929, "type": "http", "category": "xhr",
					"data": { "method": "GET", "status_code": 200, "url": "/api/0/organizations/?member=1" }
				},
				{
					"timestamp": 1488129099.931, "type": "http", "category": "xhr",
					"data": { "method": "GET", "status_code": 200, "url": "/api/0/organizations/acme-team/" }
				},
				{
					"timestamp": 1488129100.135, "type": "http", "category": "xhr",
					"data": { "method": "GET", "status_code": 200, "url": "/api/0/broadcasts/" }
				}
			]},
			"exception": { "values": [
				{
					"type": "TypeError",
					"value": "this.state.broadcasts.filter is not a function",
					"stacktrace": {
						"frames": [
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:67:4413), <anonymous>",
								"function": "XMLHttpRequest.wrapped",
								"lineno": 278, "colno": 29, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:1:12931), <anonymous>",
								"function": "XMLHttpRequest.eval",
								"lineno": 8605, "colno": 9, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:1:12931), <anonymous>",
								"function": "done",
								"lineno": 8264, "colno": 14, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:1:12931), <anonymous>",
								"function": "Object.fireWith [as resolveWith]",
								"lineno": 3211, "colno": 7, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:1:12931), <anonymous>",
								"function": "fire",
								"lineno": 3099, "colno": 30, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js:8:3111), <anonymous>",
								"function": "Object.eval [as success]",
								"lineno": 102, "colno": 23, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js:49:13510), <anonymous>",
								"function": "Request.success",
								"lineno": 82, "colno": 15, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:27:18077), <anonymous>",
								"function": "Constructor.ReactComponent.setState",
								"lineno": 64, "colno": 16, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:28:19604), <anonymous>",
								"function": "Object.enqueueSetState",
								"lineno": 210, "colno": 5, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:28:19604), <anonymous>",
								"function": "enqueueUpdate",
								"lineno": 25, "colno": 16, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:8:7050), <anonymous>",
								"function": "Object.enqueueUpdate",
								"lineno": 201, "colno": 22, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:84:48), <anonymous>",
								"function": "Object.batchedUpdates",
								"lineno": 63, "colno": 19, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:16:4797), <anonymous>",
								"function": "ReactDefaultBatchingStrategyTransaction.perform",
								"lineno": 151, "colno": 16, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:16:4797), <anonymous>",
								"function": "ReactDefaultBatchingStrategyTransaction.closeAll",
								"lineno": 204, "colno": 25, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:8:7050), <anonymous>",
								"function": "Object.flushBatchedUpdates",
								"lineno": 173, "colno": 19, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:8:7050), <anonymous>",
								"function": "ReactUpdatesFlushTransaction.perform",
								"lineno": 90, "colno": 38, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:16:4797), <anonymous>",
								"function": "ReactUpdatesFlushTransaction.perform",
								"lineno": 138, "colno": 20, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:16:4797), <anonymous>",
								"function": "ReactReconcileTransaction.perform",
								"lineno": 138, "colno": 20, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:8:7050), <anonymous>",
								"function": "runBatchedUpdates",
								"lineno": 151, "colno": 21, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:14:3867), <anonymous>",
								"function": "Object.performUpdateIfNecessary",
								"lineno": 158, "colno": 22, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper.performUpdateIfNecessary",
								"lineno": 558, "colno": 12, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper.updateComponent",
								"lineno": 642, "colno": 12, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper._performComponentUpdate",
								"lineno": 721, "colno": 10, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper._updateRenderedComponent",
								"lineno": 743, "colno": 36, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper._renderValidatedComponent",
								"lineno": 819, "colno": 34, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "ReactCompositeComponentWrapper._renderValidatedComponentWithoutOwnerOrContext",
								"lineno": 792, "colno": 27, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "measureLifeCyclePerf",
								"lineno": 74, "colno": 12, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/vendor.js:78:24141), <anonymous>",
								"function": "eval",
								"lineno": 793, "colno": 21, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js:49:13510), <anonymous>",
								"function": "Constructor.render",
								"lineno": 146, "colno": 14, "in_app": true
							},
							{
								"filename": "eval at <anonymous> (http://localhost:3000/_static/7a9894050589851ad3c1e1a2d1adb54ac08b8832/sentry/dist/app.js:49:13510), <anonymous>",
								"function": "Constructor.getUnseenIds",
								"lineno": 104, "colno": 34, "in_app": true
							}
						]
					}
				}
			]}
		}`).
		End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	assert.JSONEq(t,
		`{"id": "44444444333322221111000000000000"}`,
		bodyStr)
}

func TestStoreEvent_Post_MinimalBody(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Post("http://example.com/api/1/store/").
		Send(`{"event_id":"44444444333322221111000000000000"}`).
		End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	assert.JSONEq(t,
		`{"id": "44444444333322221111000000000000"}`,
		bodyStr)
}
