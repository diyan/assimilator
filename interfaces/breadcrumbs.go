package interfaces

import (
	"encoding/json"
	"time"
)

// Breadcrumbs interface stores information that leads up to an error.
//
// - ``message`` must be no more than 1000 characters in length.
//
// [{
//     "type": "message",
//     // timestamp can be ISO format or a unix timestamp (as float)
//     "timestamp": "2016-01-17T12:30:00",
//     "data": {
//         "message": "My raw message with interpreted strings like %s",
//     }
// ], ...}
type Breadcrumbs struct {
	Values []Breadcrumb `in:"values" json:"values"`
}

type Breadcrumb struct {
	Type      string                 `in:"type"      json:"type"`
	Timestamp time.Time              `in:"timestamp" json:"timestamp"`
	Level     interface{}            `in:"level"     json:"level,omitempty"`
	Message   string                 `in:"message"   json:"message,omitempty"`
	Category  string                 `in:"category"  json:"category,omitempty"`
	Data      map[string]interface{} `in:"data"      json:"data,omitempty"`
	EventID   interface{}            `in:"event_id"  json:"event_id,omitempty"`
}

func (*Breadcrumbs) KeyAlias() string {
	return "breadcrumbs"
}

func (*Breadcrumbs) KeyCanonical() string {
	return "sentry.interfaces.Breadcrumbs"
}

func (breadcrumbs *Breadcrumbs) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, breadcrumbs)
}

func (breadcrumbs *Breadcrumbs) DecodeRequest(request map[string]interface{}) error {
	// TODO Try to unmarshal each value in the array; if failed - skip one record, not all of them
	err := DecodeRequest(request, breadcrumbs)
	for i := 0; i < len(breadcrumbs.Values); i++ {
		breadcrumb := &breadcrumbs.Values[i]
		breadcrumb.Message = TrimLength(breadcrumb.Message, 4096)
		breadcrumb.Category = TrimLength(breadcrumb.Category, 256)
		if breadcrumb.Type == "" {
			breadcrumb.Type = "default"
		}
		if breadcrumb.Level == "info" {
			breadcrumb.Level = nil
		}
		for key, value := range breadcrumb.Data {
			if stringValue, ok := value.(string); ok {
				breadcrumb.Data[key] = stringValue
			} else {
				jsonValue, _ := json.Marshal(value)
				breadcrumb.Data[key] = string(jsonValue)
			}
		}
	}
	return err
}

// TrimLength truncates string up to length characters.
// If given string is shorter than length if will be returned without any changes.
// TODO move to the strings package
func TrimLength(v string, length int) string {
	if len(v) > length {
		return v[:length]
	}
	return v
}
