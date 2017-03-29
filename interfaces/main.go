package interfaces

import (
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/models"
)

var keyCanonicalToAlias = map[string]string{}

// keyExpected contains keys that are expected for the store event API and key-value node storage
// TODO is project, timestamp, event_id, message, time_spent is expected?
var keyExpected = map[string]bool{
	// keys from store event API
	"server_name": true,
	"logger":      true,
	"level":       true,
	"culprit":     true,
	"platform":    true,
	"release":     true,
	"tags":        true,
	"environment": true,
	"fingerprint": true,
	"modules":     true,
	"extra":       true,
	"project":     true,
	"timestamp":   true,
	"event_id":    true,
	"message":     true,
	"time_spent":  true,
	// keys from key/value node storage
	"_ref":         true,
	"_ref_version": true,
	"version":      true,
	"type":         true,
	"metadata":     true,
	"received":     true,
	"errors":       true,
}

type EventEntry struct {
	Type  string            `kv:"-" in:"-" json:"type"`
	Value models.Identifier `kv:"-" in:"-" json:"data"`
}

type EventInterfaces struct {
	// Built-in interfaces, sorted by name
	AppleCrashReport *AppleCrashReport `kv:"applecrashreport" in:"applecrashreport" json:"-"`

	Breadcrumbs *Breadcrumbs `kv:"breadcrumbs" in:"breadcrumbs" json:"-"`
	Contexts    *Contexts    `kv:"contexts"    in:"contexts"    json:"contexts,omitempty"`
	CSP         *CSP         `kv:"csp"         in:"csp"         json:"-"`
	DebugMeta   *DebugMeta   `kv:"debug_meta"  in:"debug_meta"  json:"-"`
	Device      *Device      `kv:"devce"       in:"devce"       json:"device,omitempty"`
	Exception   *Exception   `kv:"exception"   in:"exception"   json:"-"`
	LogEntry    *LogEntry    `kv:"logentry"    in:"logentry"    json:"-"`
	Query       *Query       `kv:"query"       in:"query"       json:"-"`
	Repos       *Repos       `kv:"repos"       in:"repos"       json:"-"`
	Request     *HTTP        `kv:"request"     in:"request"     json:"-"`
	SDK         *SDK         `kv:"sdk"         in:"sdk"         json:"sdk,omitempty"`
	Stacktrace  *Stacktrace  `kv:"stacktrace"  in:"stacktrace"  json:"-"`
	Template    *Template    `kv:"template"    in:"template"    json:"-"`
	Threads     *Threads     `kv:"threads"     in:"threads"     json:"-"`
	User        *User        `kv:"user"        in:"user"        json:"user"` // TODO omitempty?

	Message string       `in:"-" json:"message"`
	Entries []EventEntry `in:"-" json:"entries"` // TODO omitempty?
}

func Register(value models.Identifier) {
	keyExpected[value.KeyCanonical()] = true
	keyExpected[value.KeyAlias()] = true
	keyCanonicalToAlias[value.KeyCanonical()] = value.KeyAlias()
}

func (event *EventInterfaces) DecodeRecord(record map[string]interface{}) error {
	// TODO too many known types here, use interfaces instead
	// TODO add error handling
	// TODO remove hardcode
	event.Contexts = &Contexts{}
	event.iterate(func(value models.Identifier) {
		if decoder, ok := value.(models.RecordDecoder); ok {
			decoder.DecodeRecord(record)
		}
	})
	fillEntries(event)
	return nil
}

func (event *EventInterfaces) DecodeRequest(request map[string]interface{}) error {
	event.iterate(func(value models.Identifier) {
		if decoder, ok := value.(models.RequestDecoder); ok {
			decoder.DecodeRequest(request)
		}
	})
	return nil
}

func (event *EventInterfaces) iterate(hookFn func(models.Identifier)) {
	interfaces := []models.Identifier{
		event.AppleCrashReport,
		event.Breadcrumbs,
		event.Contexts,
		event.CSP,
		event.DebugMeta,
		event.Device,
		event.Exception,
		event.LogEntry,
		event.Query,
		event.Repos,
		event.Request,
		event.SDK,
		event.Stacktrace,
		event.Template,
		event.Threads,
		event.User,
	}
	for _, value := range interfaces {
		hookFn(value)
	}
}

func fillEntries(event *EventInterfaces) {
	// TODO remove hardcode
	event.Message = event.LogEntry.Message
	// Reserved interfaces which remains in the root: User, SDK, Device, Contexts
	interfaces := []models.Identifier{
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
		if value == nil || reflect.ValueOf(value).IsNil() {
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

// TODO return error that contains all unexpected keys
func ToAliasKeys(source map[string]interface{}) map[string]interface{} {
	rv := map[string]interface{}{}
	for key, value := range source {
		if alias, ok := keyCanonicalToAlias[key]; ok {
			rv[alias] = value
		} else if keyExpected[key] {
			rv[key] = value
		} else {
			logrus.WithField("keyName", key).Warn("unexpected key in the event map")
		}
	}
	return rv
}
