package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupEventsLatests_Get(t *testing.T) {
	client, factory := Setup(t)
	defer TearDown(t)

	factory.SaveOrganization(factory.MakeOrganization())
	factory.SaveTeam(factory.MakeTeam())
	factory.SaveProject(factory.MakeProject())
	factory.SaveEventGroup(factory.MakeEventGroup())
	factory.SaveEvent(factory.MakeEvent())
	factory.SaveEventNodeBlob(factory.MakeEventNodeBlobV2())

	res, bodyStr, errs := client.Get("http://example.com/api/0/issues/1/events/latest/").End()
	assert.Nil(t, errs)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, bodyStr)
	// TODO dateReceived 2017-02-24T09:25:39Z -> 2999-01-01T00:00:00Z
	// TODO eventID dcf8c1d1cd284d3fbfeffb43ddb7c0f7 or 8b76c9e141134b38a1e5f67d8709618d -> 44444444333322221111000000000000
	assert.JSONEq(t, `{
            "id": "1",
            "groupID": "1",
            "release": null,
            "platform": "python",
            "message": "This is a test message generated using `+"``raven test``"+`",
            "eventID": "dcf8c1d1cd284d3fbfeffb43ddb7c0f7",
            "userReport": null,
            "nextEventID": null,
            "previousEventID": null,
            "size": 6597,
            "errors": [],
            "type": "default",
            "metadata": {
                "title": "This is a test message generated using `+"``raven test``"+`"
            },
            "tags": [
                { "key": "server_name", "value": "f23dd84359d9" },
                { "key": "level", "value": "info" }
            ],
            "dateCreated": "2999-01-01T00:00:00Z",
            "dateReceived": "2017-02-24T09:25:39Z",
            "user": null,
            "entries": [
                {
                    "type": "message",
                    "data": {
                        "message": "This is a test message generated using `+"``raven test``"+`"
                    }
                },
                {
                    "type": "stacktrace",
                    "data": {
                        "frames": [
                            {
                                "colNo": null,
                                "lineNo": 391,                                
                                "instructionOffset": null,
                                "instructionAddr": null,
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/lib/python2.7/site-packages/raven/base.py",                                
                                "module": "raven.base",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "raven/base.py",
                                "function": "build_msg",
                                "context": [
                                    [ 386, "                frames = stack" ],
                                    [ 387, "" ],
                                    [ 388, "            stack_info = get_stack_info(" ],
                                    [ 389, "                frames," ],
                                    [ 390, "                transformer=self.transform," ],
                                    [ 391, "                capture_locals=self.capture_locals," ],
                                    [ 392, "            )" ],
                                    [ 393, "            data.update({" ],
                                    [ 394, "                'stacktrace': stack_info," ],
                                    [ 395, "            })" ],
                                    [ 396, "" ]
                                ],
                                "vars": {
                                    "public_key": null,
                                    "v": {
                                        "'message'": "u'This is a test message generated using `+"``raven test``"+`'",
                                        "'params'": [],
                                        "'formatted'": null
                                    },
                                    "event_type": "'raven.events.Message'",
                                    "culprit": null,
                                    "tags": null,
                                    "event_id": "'8b76c9e141134b38a1e5f67d8709618d'",
                                    "k": "'sentry.interfaces.Message'",
                                    "extra": {
                                        "'loadavg'": [ 0.11, 0.28, 0.38 ],
                                        "'user'": "'root'"
                                    },
                                    "frames": "<generator object iter_stack_frames at 0x7f026dd8cb40>",
                                    "stack": true,
                                    "time_spent": null,
                                    "fingerprint": null,
                                    "handler": "<raven.events.Message object at 0x7f026ddb5dd0>",
                                    "result": {
                                        "'sentry.interfaces.Message'": {
                                            "'message'": "u'This is a test message generated using `+"``raven test``"+`'",
                                            "'params'": [],
                                            "'formatted'": null
                                        },
                                        "'message'": "u'This is a test message generated using `+"``raven test``"+`'"
                                    },
                                    "kwargs": {
                                        "'message'": "'This is a test message generated using `+"``raven test``"+`'",
                                        "'level'": 20
                                    },
                                    "date": null,
                                    "data": {
                                        "'extra'": {},
                                        "'message'": "u'This is a test message generated using `+"``raven test``"+`'",
                                        "'tags'": {},
                                        "'sentry.interfaces.Message'": {
                                            "'message'": "u'This is a test message generated using `+"``raven test``"+`'",
                                            "'params'": [],
                                            "'formatted'": null
                                        }
                                    },
                                    "self": "<raven.base.Client object at 0x7f0274921190>"
                                }
                            },
                            {                                                                
                                "colNo": null,
                                "lineNo": 608,
                                "instructionOffset": null,
                                "instructionAddr": null,
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/lib/python2.7/site-packages/raven/base.py",
                                "module": "raven.base",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "raven/base.py",
                                "function": "capture",
                                "context": [
                                    [ 603, "                return" ],
                                    [ 604, "            self.record_exception_seen(exc_info)" ],
                                    [ 605, "" ],
                                    [ 606, "        data = self.build_msg(" ],
                                    [ 607, "            event_type, data, date, time_spent, extra, stack, tags=tags," ],
                                    [ 608, "            **kwargs)" ],
                                    [ 609, "" ],
                                    [ 610, "        self.send(**data)" ],
                                    [ 611, "" ],
                                    [ 612, "        return data['event_id']" ],
                                    [ 613, "" ]
                                ],
                                "vars": {
                                    "event_type": "'raven.events.Message'",
                                    "tags": null,
                                    "self": "<raven.base.Client object at 0x7f0274921190>",
                                    "extra": {
                                        "'loadavg'": [ 0.11, 0.28, 0.38 ],
                                        "'user'": "'root'"
                                    },
                                    "time_spent": null,
                                    "kwargs": {
                                        "'message'": "'This is a test message generated using `+"``raven test``"+`'",
                                        "'level'": 20
                                    },
                                    "date": null,
                                    "exc_info": null,
                                    "data": null,
                                    "stack": true
                                }
                            },
                            {
                                "colNo": null,
                                "lineNo": 759,
                                "instructionOffset": null,
                                "instructionAddr": null,
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/lib/python2.7/site-packages/raven/base.py",
                                "module": "raven.base",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "raven/base.py",
                                "function": "captureMessage",
                                "context": [
                                    [ 754, "        \"\"\"" ],
                                    [ 755, "        Creates an event from `+"``message``"+`." ],
                                    [ 756, "" ],
                                    [ 757, "        >>> client.captureMessage('My event just happened!')" ],
                                    [ 758, "        \"\"\"" ],
                                    [ 759, "        return self.capture('raven.events.Message', message=message, **kwargs)" ],
                                    [ 760, "" ],
                                    [ 761, "    def captureException(self, exc_info=None, **kwargs):" ],
                                    [ 762, "        \"\"\"" ],
                                    [ 763, "        Creates an event from an exception." ],
                                    [ 764, "" ]
                                ],
                                "vars": {
                                    "message": "'This is a test message generated using `+"``raven test``"+`'",
                                    "self": "<raven.base.Client object at 0x7f0274921190>",
                                    "kwargs": {
                                        "'extra'": {
                                            "'loadavg'": [ 0.11, 0.28, 0.38 ],
                                            "'user'": "'root'"
                                        },
                                        "'stack'": true,
                                        "'data'": null,
                                        "'level'": 20,
                                        "'tags'": null
                                    }
                                }
                            },
                            {
                                "colNo": null,
                                "lineNo": 81,
                                "instructionOffset": null,
                                "instructionAddr": null,                               
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/lib/python2.7/site-packages/raven/scripts/runner.py",
                                "module": "raven.scripts.runner",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "raven/scripts/runner.py",
                                "function": "send_test_message",
                                "context": [
                                    [ 76, "        level=logging.INFO," ],
                                    [ 77, "        stack=True," ],
                                    [ 78, "        tags=options.get('tags', {})," ],
                                    [ 79, "        extra={" ],
                                    [ 80, "            'user': get_uid()," ],
                                    [ 81, "            'loadavg': get_loadavg()," ],
                                    [ 82, "        }," ],
                                    [ 83, "    )" ],
                                    [ 84, "" ],
                                    [ 85, "    sys.stdout.write('Event ID was %r\\n' % (ident,))" ],
                                    [ 86, "" ]
                                ],
                                "vars": {
                                    "k": "[Filtered]",
                                    "client": "<raven.base.Client object at 0x7f0274921190>",
                                    "data": null,
                                    "options": {
                                        "'tags'": null,
                                        "'data'": null
                                    },
                                    "remote_config": "<raven.conf.remote.RemoteConfig object at 0x7f026dd96150>"
                                }
                            },
                            {
                                "colNo": null,
                                "lineNo": 113,
                                "instructionOffset": null,
                                "instructionAddr": null,                                
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/lib/python2.7/site-packages/raven/scripts/runner.py",
                                "module": "raven.scripts.runner",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "raven/scripts/runner.py",
                                "function": "main",
                                "context": [
                                    [ 108, "    print(\" \", dsn)" ],
                                    [ 109, "    print()" ],
                                    [ 110, "" ],
                                    [ 111, "    client = Client(dsn, include_paths=['raven'])" ],
                                    [ 112, "" ],
                                    [ 113, "    send_test_message(client, opts.__dict__)" ],
                                    [ 114, "" ],
                                    [ 115, "    # TODO(dcramer): correctly support async models" ],
                                    [ 116, "    time.sleep(3)" ],
                                    [ 117, "    if client.state.did_fail():" ],
                                    [ 118, "        sys.stdout.write('error!\\n')" ]
                                ],
                                "vars": {
                                    "parser": "<optparse.OptionParser instance at 0x7f026dff2560>",
                                    "args": [
                                        "'test'",
                                        "'http://763a78a695424ed687cf8b7dc26d3161:[Filtered]@localhost:9000/2'"
                                    ],
                                    "dsn": "'http://763a78a695424ed687cf8b7dc26d3161:[Filtered]@localhost:9000/2'",
                                    "client": "<raven.base.Client object at 0x7f0274921190>",
                                    "root": "<logging.Logger object at 0x7f026dfc57d0>",
                                    "opts": "<Values at 0x7f026dd89320: {'data': None, 'tags': None}>"
                                }
                            },
                            {
                                "colNo": null,
                                "lineNo": 11,
                                "instructionOffset": null,
                                "instructionAddr": null,
                                "symbol": null,
                                "symbolAddr": null,
                                "absPath": "/usr/local/bin/raven",
                                "module": "__main__",
                                "package": null,
                                "platform": null,
                                "errors": null,
                                "inApp": false,
                                "filename": "bin/raven",
                                "function": "<module>",
                                "context": [
                                    [ 6, "" ],
                                    [ 7, "from raven.scripts.runner import main" ],
                                    [ 8, "" ],
                                    [ 9, "if __name__ == '__main__':" ],
                                    [ 10, "    sys.argv[0] = re.sub(r'(-script\\.pyw?|\\.exe)?$', '', sys.argv[0])" ],
                                    [ 11, "    sys.exit(main())" ]
                                ],
                                "vars": {
                                    "__builtins__": "<module '__builtin__' (built-in)>",
                                    "__file__": "'/usr/local/bin/raven'",
                                    "__package__": null,
                                    "sys": "<module 'sys' (built-in)>",
                                    "re": "<module 're' from '/usr/local/lib/python2.7/re.pyc'>",
                                    "__name__": "'__main__'",
                                    "main": "<function main from raven.scripts.runner at 0x7f026dd95aa0>",
                                    "__doc__": null
                                }
                            }
                        ],
                        "framesOmitted": null,
                        "hasSystemFrames": false
                    }
                }
            ],
            "packages": {
                "python": "2.7.13",
                "raven": "5.32.0"
            },
            "sdk": {
                "clientIP": "127.0.0.1",
                "version": "5.32.0",
                "name": "raven-python",
                "upstream": {
                    "url": "https://docs.sentry.io/clients/python/",
                    "isNewer": false,
                    "name": "raven-python"
                }
            },
            "contexts": {},
            "context": {
                "sys.argv": [
                    "'/usr/local/bin/raven'",
                    "'test'",
                    "'http://763a78a695424ed687cf8b7dc26d3161:[Filtered]@localhost:9000/2'"
                ],
                "loadavg": [ 0.11, 0.28, 0.38 ],
                "user": "'root'"
            }
        }`,
		bodyStr)
}
