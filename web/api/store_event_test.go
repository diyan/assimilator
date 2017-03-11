package api_test

import (
	"testing"

	"github.com/diyan/assimilator/testutil/fixture"
	"github.com/stretchr/testify/assert"
)

func TestStoreEvent_Post_RavenJSPayload(t *testing.T) {
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

func TestStoreEvent_Post_RavenPythonPayload(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Post("http://example.com/api/1/store/").
		Send(`{
			"project": "1",			
			"event_id": "44444444333322221111000000000000",
			"platform": "python",
			"level": 20,
			"server_name": "falcon",
			"timestamp": "2017-03-11T09:58:41Z",
			"time_spent": null,			
			"message": "This is a test message generated using ` + "``raven test``" + `'",
			"sentry.interfaces.Message": {
				"message": "This is a test message generated using ` + "``raven test``" + `'",
				"params": [],
				"formatted": null
			},
			"repos": {},
			"tags": {},
			"modules": {
				"python": "2.7.13",
				"raven": "6.0.0"
			},
			"sdk": {
				"version": "6.0.0",
				"name": "raven-python"
			},
			"stacktrace": {
				"frames": [
					{
						"function": "build_msg",
						"abs_path": "/tmp/env2/lib/python2.7/site-packages/raven/base.py",
						"pre_context": [
							"                frames = stack",
							"",
							"            stack_info = get_stack_info(",
							"                frames,",
							"                transformer=self.transform,"
						],
						"post_context": [
							"            )",
							"            data.update({",
							"                'stacktrace': stack_info,",
							"            })",
							""
						],
						"vars": {
							"public_key": null,
							"v": {
								"'message'": "u'This is a test message generated using ` + "``raven test``" + `'",
								"'params'": [],
								"'formatted'": null
							},
							"event_type": "'raven.events.Message'",
							"culprit": null,
							"tags": null,
							"event_id": "'894e39cfdb3e4b8895a6e90cae206782'",
							"k": "'sentry.interfaces.Message'",
							"extra": {
								"'loadavg'": [
									0.24,
									0.57,
									0.49
								],
								"'user'": "'alexey'"
							},
							"frames": "<generator object iter_stack_frames at 0x7fe197ad4f50>",
							"stack": true,
							"time_spent": null,
							"fingerprint": null,
							"handler": "<raven.events.Message object at 0x7fe1977d8d90>",
							"result": {
								"'sentry.interfaces.Message'": {
									"'message'": "u'This is a test message generated using ` + "``raven test``" + `'",
									"'params'": [],
									"'formatted'": null
								},
								"'message'": "u'This is a test message generated using ` + "``raven test``" + `'"
							},
							"kwargs": {
								"'message'": "'This is a test message generated using ` + "``raven test``" + `'",
								"'level'": 20
							},
							"date": null,
							"data": {
								"'extra'": {},
								"'message'": "u'This is a test message generated using ` + "``raven test``" + `'",
								"'tags'": {},
								"'sentry.interfaces.Message'": {
									"'message'": "u'This is a test message generated using ` + "``raven test``" + `'",
									"'params'": [],
									"'formatted'": null
								}
							},
							"self": "<raven.base.Client object at 0x7fe197a45b10>"
						},
						"module": "raven.base",
						"filename": "raven/base.py",
						"lineno": 406,
						"in_app": false,
						"context_line": "                capture_locals=self.capture_locals,"
					},
					{
						"function": "capture",
						"abs_path": "/tmp/env2/lib/python2.7/site-packages/raven/base.py",
						"pre_context": [
							"                return",
							"            self.record_exception_seen(exc_info)",
							"",
							"        data = self.build_msg(",
							"            event_type, data, date, time_spent, extra, stack, tags=tags,"
						],
						"post_context": [
							"",
							"        self.send(**data)",
							"",
							"        self._local_state.last_event_id = data['event_id']",
							""
						],
						"vars": {
							"event_type": "'raven.events.Message'",
							"tags": null,
							"self": "<raven.base.Client object at 0x7fe197a45b10>",
							"extra": {
								"'loadavg'": [
									0.24,
									0.57,
									0.49
								],
								"'user'": "'alexey'"
							},
							"time_spent": null,
							"kwargs": {
								"'message'": "'This is a test message generated using ` + "``raven test``" + `'",
								"'level'": 20
							},
							"date": null,
							"exc_info": null,
							"data": null,
							"stack": true
						},
						"module": "raven.base",
						"filename": "raven/base.py",
						"lineno": 624,
						"in_app": false,
						"context_line": "            **kwargs)"
					},
					{
						"function": "captureMessage",
						"abs_path": "/tmp/env2/lib/python2.7/site-packages/raven/base.py",
						"pre_context": [
							"        \"\"\"",
							"        Creates an event from` + "``message``" + `.",
							"",
							"        >>> client.captureMessage('My event just happened!')",
							"        \"\"\""
						],
						"post_context": [
							"",
							"    def captureException(self, exc_info=None, **kwargs):",
							"        \"\"\"",
							"        Creates an event from an exception.",
							""
						],
						"vars": {
							"message": "'This is a test message generated using ` + "``raven test``" + `'",
							"self": "<raven.base.Client object at 0x7fe197a45b10>",
							"kwargs": {
								"'extra'": {
									"'loadavg'": [
										0.24,
										0.57,
										0.49
									],
									"'user'": "'alexey'"
								},
								"'stack'": true,
								"'data'": null,
								"'level'": 20,
								"'tags'": null
							}
						},
						"module": "raven.base",
						"filename": "raven/base.py",
						"lineno": 778,
						"in_app": false,
						"context_line": "        return self.capture('raven.events.Message', message=message, **kwargs)"
					},
					{
						"function": "send_test_message",
						"abs_path": "/tmp/env2/lib/python2.7/site-packages/raven/scripts/runner.py",
						"pre_context": [
							"        level=logging.INFO,",
							"        stack=True,",
							"        tags=options.get('tags', {}),",
							"        extra={",
							"            'user': get_uid(),"
						],
						"post_context": [
							"        },",
							"    )",
							"",
							"    sys.stdout.write('Event ID was %r\\n' % (ident,))",
							""
						],
						"vars": {
							"k": "'secret_key'",
							"client": "<raven.base.Client object at 0x7fe197a45b10>",
							"data": null,
							"options": {
								"'tags'": null,
								"'data'": null
							},
							"remote_config": "<raven.conf.remote.RemoteConfig object at 0x7fe197a4c110>"
						},
						"module": "raven.scripts.runner",
						"filename": "raven/scripts/runner.py",
						"lineno": 81,
						"in_app": false,
						"context_line": "            'loadavg': get_loadavg(),"
					},
					{
						"function": "main",
						"abs_path": "/tmp/env2/lib/python2.7/site-packages/raven/scripts/runner.py",
						"pre_context": [
							"    print(\" \", dsn)",
							"    print()",
							"",
							"    client = Client(dsn, include_paths=['raven'])",
							""
						],
						"post_context": [
							"",
							"    # TODO(dcramer): correctly support async models",
							"    time.sleep(3)",
							"    if client.state.did_fail():",
							"        sys.stdout.write('error!\\n')"
						],
						"vars": {
							"parser": "<optparse.OptionParser instance at 0x7fe197a19c68>",
							"args": [
								"'test'",
								"'http://111:111@localhost:3000/1'"
							],
							"dsn": "'http://111:111@localhost:3000/1'",
							"client": "<raven.base.Client object at 0x7fe197a45b10>",
							"root": "<logging.Logger object at 0x7fe197accdd0>",
							"opts": "<Values at 0x7fe197a19128: {'data': None, 'tags': None}>"
						},
						"module": "raven.scripts.runner",
						"filename": "raven/scripts/runner.py",
						"lineno": 113,
						"in_app": false,
						"context_line": "    send_test_message(client, opts.__dict__)"
					},
					{
						"function": "<module>",
						"abs_path": "/tmp/env2/bin/raven",
						"pre_context": [
							"",
							"from raven.scripts.runner import main",
							"",
							"if __name__ == '__main__':",
							"    sys.argv[0] = re.sub(r'(-script\\.pyw?|\\.exe)?$', '', sys.argv[0])"
						],
						"post_context": [],
						"vars": {
							"__builtins__": "<module '__builtin__' (built-in)>",
							"__file__": "'/tmp/env2/bin/raven'",
							"__package__": null,
							"sys": "<module 'sys' (built-in)>",
							"re": "<module 're' from '/tmp/env2/lib/python2.7/re.pyc'>",
							"__name__": "'__main__'",
							"main": "<function main from raven.scripts.runner at 0x7fe197ab8230>",
							"__doc__": null
						},
						"module": "__main__",
						"filename": "bin/raven",
						"lineno": 11,
						"in_app": false,
						"context_line": "    sys.exit(main())"
					}
				]
			},
			"extra": {
				"sys.argv": [
					"'/tmp/env2/bin/raven'",
					"'test'",
					"'http://111:111@localhost:3000/1'"
				],
				"loadavg": [
					0.24,
					0.57,
					0.49
				],
				"user": "'alexey'"
			}
		}`).
		End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	assert.JSONEq(t,
		`{"id": "44444444333322221111000000000000"}`,
		bodyStr)
}

func TestStoreEvent_Post_InvalidBody(t *testing.T) {
	client, factory := fixture.Setup(t)
	defer fixture.TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())

	res, bodyStr, errs := client.Post("http://example.com/api/1/store/").
		Send(`{
			"unexpectedObject": {
				"array": ["hello", 123.123, true, null],
				"string": "hello",
				"number": 123.123,
				"boolean": true,
				"null": null
			},
			"unexpectedArray": ["hello", 123.123, true, null],
			"unexpectedString": "hello",
			"unexpectedNumber": 123.123,
			"unexpectedBoolean": true,
			"unexpectedNull": null
		}`).
		End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	//assert.JSONEq(t,
	//	`{"id": "44444444333322221111000000000000"}`,
	//	bodyStr)
}
