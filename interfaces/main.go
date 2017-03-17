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

func (event *EventInterfaces) UnmarshalRecord(nodeBlob interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	event.Request = &HTTP{}
	event.Request.UnmarshalRecord(nodeBlob)
	event.User = &User{}
	event.User.UnmarshalRecord(nodeBlob)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.UnmarshalRecord(nodeBlob)
	event.SDK = &SDK{}
	event.SDK.UnmarshalRecord(nodeBlob)
	event.Exception = &Exception{}
	event.Exception.UnmarshalRecord(nodeBlob)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.UnmarshalRecord(nodeBlob)
	return nil
}

func (event *EventInterfaces) UnmarshalAPI(rawEvent map[string]interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	event.Request = &HTTP{}
	event.Request.UnmarshalAPI(rawEvent)
	event.User = &User{}
	event.User.UnmarshalAPI(rawEvent)
	event.Breadcrumbs = &Breadcrumbs{}
	event.Breadcrumbs.UnmarshalAPI(rawEvent)
	event.SDK = &SDK{}
	event.SDK.UnmarshalAPI(rawEvent)
	event.Exception = &Exception{}
	event.Exception.UnmarshalAPI(rawEvent)
	event.Stacktrace = &Stacktrace{}
	event.Stacktrace.UnmarshalAPI(rawEvent)
	return nil
}
