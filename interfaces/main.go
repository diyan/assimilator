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
	event.Request = &HTTP{}
	event.Request.DecodeRecord(record)
	event.User = &User{}
	event.User.DecodeRecord(record)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.DecodeRecord(record)
	event.SDK = &SDK{}
	event.SDK.DecodeRecord(record)
	event.Exception = &Exception{}
	event.Exception.DecodeRecord(record)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.DecodeRecord(record)
	return nil
}

func (event *EventInterfaces) DecodeRequest(request map[string]interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	event.Request = &HTTP{}
	event.Request.DecodeRequest(request)
	event.User = &User{}
	event.User.DecodeRequest(request)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.DecodeRequest(request)
	event.SDK = &SDK{}
	event.SDK.DecodeRequest(request)
	event.Exception = &Exception{}
	event.Exception.DecodeRequest(request)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.DecodeRequest(request)
	return nil
}
