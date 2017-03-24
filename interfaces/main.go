package interfaces

import (
	"fmt"

	"github.com/diyan/assimilator/models"
	"github.com/pkg/errors"
)

type EventInterfaces struct {
	// Built-in interfaces, sorted by name
	AppleCrashReport models.Marshaler
	Breadcrumbs      models.Marshaler
	Contexts         models.Marshaler
	CSP              models.Marshaler
	DebugMeta        models.Marshaler
	Device           models.Marshaler
	Exception        models.Marshaler
	LogEntry         models.Marshaler
	Query            models.Marshaler
	Repos            models.Marshaler
	Request          models.Marshaler
	SDK              models.Marshaler
	Stacktrace       models.Marshaler
	Template         models.Marshaler
	Threads          models.Marshaler
	User             models.Marshaler
}

func DecodeRecord(keyAlias, keyCanonical string, record interface{}, target interface{}) error {
	recordMap, ok := record.(map[interface{}]interface{})
	if !ok {
		return nil
	}
	value, ok := recordMap[keyAlias]
	if !ok {
		value, ok = recordMap[keyCanonical]
		if !ok {
			return nil
		}
	}
	err := models.DecodeRecord(value, target)
	return errors.Wrapf(err, fmt.Sprintf("can not decode node record to %s", keyCanonical))
}

func DecodeRequest(keyAlias, keyCanonical string, request map[string]interface{}, target interface{}) error {
	value, ok := request[keyAlias].(map[string]interface{})
	if !ok {
		value, ok = request[keyCanonical].(map[string]interface{})
		if !ok {
			return nil
		}
	}
	err := models.DecodeRequest(value, target)
	return errors.Wrapf(err, fmt.Sprintf("can not decode request to %s", keyCanonical))
}

func (event *EventInterfaces) DecodeRecord(record interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	event.AppleCrashReport = &AppleCrashReport{}
	event.AppleCrashReport.DecodeRecord(record)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.DecodeRecord(record)
	event.Contexts = &Contexts{}
	event.Contexts.DecodeRecord(record)
	event.CSP = &CSP{}
	event.CSP.DecodeRecord(record)
	event.DebugMeta = &DebugMeta{}
	event.DebugMeta.DecodeRecord(record)
	event.Device = &Device{}
	event.Device.DecodeRecord(record)
	event.Exception = &Exception{}
	event.Exception.DecodeRecord(record)
	event.Request = &HTTP{}
	event.Request.DecodeRecord(record)
	event.LogEntry = &LogEntry{}
	event.LogEntry.DecodeRecord(record)
	event.Query = &Query{}
	event.Query.DecodeRecord(record)
	event.Repos = &Repos{}
	event.Repos.DecodeRecord(record)
	event.SDK = &SDK{}
	event.SDK.DecodeRecord(record)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.DecodeRecord(record)
	event.Template = &Template{}
	event.Template.DecodeRecord(record)
	event.Threads = &Threads{}
	event.Threads.DecodeRecord(record)
	event.User = &User{}
	event.User.DecodeRecord(record)
	return nil
}

func (event *EventInterfaces) DecodeRequest(request map[string]interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	event.AppleCrashReport = &AppleCrashReport{}
	event.AppleCrashReport.DecodeRequest(request)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.DecodeRequest(request)
	event.Contexts = &Contexts{}
	event.Contexts.DecodeRequest(request)
	event.CSP = &CSP{}
	event.CSP.DecodeRequest(request)
	event.DebugMeta = &DebugMeta{}
	event.DebugMeta.DecodeRequest(request)
	event.Device = &Device{}
	event.Device.DecodeRequest(request)
	event.Exception = &Exception{}
	event.Exception.DecodeRequest(request)
	event.Request = &HTTP{}
	event.Request.DecodeRequest(request)
	event.LogEntry = &LogEntry{}
	event.LogEntry.DecodeRequest(request)
	event.Query = &Query{}
	event.Query.DecodeRequest(request)
	event.Repos = &Repos{}
	event.Repos.DecodeRequest(request)
	event.SDK = &SDK{}
	event.SDK.DecodeRequest(request)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.DecodeRequest(request)
	event.Template = &Template{}
	event.Template.DecodeRequest(request)
	event.Threads = &Threads{}
	event.Threads.DecodeRequest(request)
	event.User = &User{}
	event.User.DecodeRequest(request)
	return nil
}
