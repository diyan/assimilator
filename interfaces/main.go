package interfaces

import (
	"fmt"

	"github.com/diyan/assimilator/models"
	"github.com/pkg/errors"
)

var errNoValue = errors.New("value was not provided")

type EventEntry struct {
	Type  string           `kv:"-" in:"-" json:"type"`
	Value models.Marshaler `kv:"-" in:"-" json:"data"`
}

type EventInterfaces struct {
	// Built-in interfaces, sorted by name
	AppleCrashReport models.Marshaler `json:"-"`
	Breadcrumbs      models.Marshaler `json:"-"`
	Contexts         models.Marshaler `json:"contexts,omitempty"`
	CSP              models.Marshaler `json:"-"`
	DebugMeta        models.Marshaler `json:"-"`
	Device           models.Marshaler `json:"device,omitempty"`
	Exception        models.Marshaler `json:"-"`
	LogEntry         models.Marshaler `json:"-"`
	Query            models.Marshaler `json:"-"`
	Repos            models.Marshaler `json:"-"`
	Request          models.Marshaler `json:"-"`
	SDK              models.Marshaler `json:"sdk,omitempty"`
	Stacktrace       models.Marshaler `json:"-"`
	Template         models.Marshaler `json:"-"`
	Threads          models.Marshaler `json:"-"`
	User             models.Marshaler `json:"user"` // TODO omitempty?

	Message string       `json:"message"`
	Entries []EventEntry `json:"entries"` // TODO omitempty?
}

func DecodeRecord(record interface{}, target models.Marshaler) error {
	recordMap, ok := record.(map[interface{}]interface{})
	if !ok {
		return nil
	}
	value, ok := recordMap[target.KeyAlias()]
	if !ok {
		value, ok = recordMap[target.KeyCanonical()]
		if !ok {
			return errNoValue
		}
	}
	err := models.DecodeRecord(value, target)
	return errors.Wrapf(err, fmt.Sprintf("can not decode node record to %s", target.KeyCanonical()))
}

func DecodeRequest(request map[string]interface{}, target models.Marshaler) error {
	value, ok := request[target.KeyAlias()].(map[string]interface{})
	if !ok {
		value, ok = request[target.KeyCanonical()].(map[string]interface{})
		if !ok {
			return errNoValue
		}
	}
	err := models.DecodeRequest(value, target)
	return errors.Wrapf(err, fmt.Sprintf("can not decode request to %s", target.KeyCanonical()))
}

func (event *EventInterfaces) DecodeRecord(record interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	var value models.Marshaler = &AppleCrashReport{}
	if err := value.DecodeRecord(record); err == nil {
		event.AppleCrashReport = value
	}
	value = &Breadcrumbs{}
	if err := value.DecodeRecord(record); err == nil {
		event.Breadcrumbs = value
	}
	value = &Contexts{}
	if err := value.DecodeRecord(record); err == nil {
		event.Contexts = value
	}
	// TODO remove hardcode
	event.Contexts = &Contexts{}
	value = &CSP{}
	if err := value.DecodeRecord(record); err == nil {
		event.CSP = value
	}
	value = &DebugMeta{}
	if err := value.DecodeRecord(record); err == nil {
		event.DebugMeta = value
	}
	value = &Device{}
	if err := value.DecodeRecord(record); err == nil {
		event.Device = value
	}
	value = &Exception{}
	if err := value.DecodeRecord(record); err == nil {
		event.Exception = value
	}
	value = &HTTP{}
	if err := value.DecodeRecord(record); err == nil {
		event.Request = value
	}
	value = &LogEntry{}
	if err := value.DecodeRecord(record); err == nil {
		event.LogEntry = value
	}
	value = &Query{}
	if err := value.DecodeRecord(record); err == nil {
		event.Query = value
	}
	value = &Repos{}
	if err := value.DecodeRecord(record); err == nil {
		event.Repos = value
	}
	value = &SDK{}
	if err := value.DecodeRecord(record); err == nil {
		event.SDK = value
	}
	value = &Stacktrace{}
	if err := value.DecodeRecord(record); err == nil {
		event.Stacktrace = value
	}
	value = &Template{}
	if err := value.DecodeRecord(record); err == nil {
		event.Template = value
	}
	value = &Threads{}
	if err := value.DecodeRecord(record); err == nil {
		event.Threads = value
	}
	value = &User{}
	if err := value.DecodeRecord(record); err == nil {
		event.User = value
	}

	fillEntries(event)
	return nil
}

func (event *EventInterfaces) DecodeRequest(request map[string]interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	var value models.Marshaler = &AppleCrashReport{}
	if err := value.DecodeRequest(request); err == nil {
		event.AppleCrashReport = value
	}
	value = &Breadcrumbs{}
	if err := value.DecodeRequest(request); err == nil {
		event.Breadcrumbs = value
	}
	value = &Contexts{}
	if err := value.DecodeRequest(request); err == nil {
		event.Contexts = value
	}
	value = &CSP{}
	if err := value.DecodeRequest(request); err == nil {
		event.CSP = value
	}
	value = &DebugMeta{}
	if err := value.DecodeRequest(request); err == nil {
		event.DebugMeta = value
	}
	value = &Device{}
	if err := value.DecodeRequest(request); err == nil {
		event.Device = value
	}
	value = &Exception{}
	if err := value.DecodeRequest(request); err == nil {
		event.Exception = value
	}
	value = &HTTP{}
	if err := value.DecodeRequest(request); err == nil {
		event.Request = value
	}
	value = &LogEntry{}
	if err := value.DecodeRequest(request); err == nil {
		event.LogEntry = value
	}
	value = &Query{}
	if err := value.DecodeRequest(request); err == nil {
		event.Query = value
	}
	value = &Repos{}
	if err := value.DecodeRequest(request); err == nil {
		event.Repos = value
	}
	value = &SDK{}
	if err := value.DecodeRequest(request); err == nil {
		event.SDK = value
	}
	value = &Stacktrace{}
	if err := value.DecodeRequest(request); err == nil {
		event.Stacktrace = value
	}
	value = &Template{}
	if err := value.DecodeRequest(request); err == nil {
		event.Template = value
	}
	value = &Threads{}
	if err := value.DecodeRequest(request); err == nil {
		event.Threads = value
	}
	value = &User{}
	if err := value.DecodeRequest(request); err == nil {
		event.User = value
	}

	return nil
}

func fillEntries(event *EventInterfaces) {
	// TODO remove hardcode
	event.Message = event.LogEntry.(*LogEntry).Message
	// Reserved interfaces which remains in the root: User, SDK, Device, Contexts
	interfaces := []models.Marshaler{
		event.AppleCrashReport,
		event.Breadcrumbs,
		event.CSP,
		event.DebugMeta,
		event.Exception,
		event.LogEntry,
		event.Query,
		event.Repos,
		event.Request,
		event.Stacktrace,
		event.Template,
		event.Threads,
	}
	for _, value := range interfaces {
		if value == nil {
			continue
		}
		// TODO remove hardcode
		typeName := value.KeyAlias()
		if typeName == "logentry" {
			typeName = "message"
		}
		event.Entries = append(event.Entries, EventEntry{
			Type:  typeName,
			Value: value,
		})
	}
}
