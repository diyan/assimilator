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
	Values []Breadcrumb `json:"values"`
}

type Breadcrumb struct {
	Type      string      `json:"type"`
	Timestamp time.Time   `json:"timestamp"` // TODO add JSONTime custom type alias
	Level     interface{} `json:"level,omitempty"`
	Message   string      `json:"message,omitempty"`
	Category  string      `json:"category,omitempty"`
	Data      interface{} `json:"data,omitempty"` // TODO add custom type with MarshalJSON/UnmarshalJSON
	EventID   interface{} `json:"event_id,omitempty"`
}

func (breadcrumbs *Breadcrumbs) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (breadcrumbs *Breadcrumbs) UnmarshalAPI(rawEvent map[string]interface{}) error {
	// TODO Try to unmarshal each value in the array; if failed - skip one record, not all of them
	rawBreadcrumbs, ok := rawEvent["breadcrumbs"].(map[string]interface{})
	if !ok {
		rawBreadcrumbs, ok = rawEvent["sentry.interfaces.Breadcrumbs"].(map[string]interface{})
		if !ok {
			return nil
		}
	}
	for _, rawBreadcrumb := range rawBreadcrumbs["values"].([]interface{}) {
		breadcrumbMap := rawBreadcrumb.(map[string]interface{})
		breadcrumb := Breadcrumb{
			Type:      breadcrumbMap["type"].(string),
			Timestamp: time.Unix(int64(breadcrumbMap["timestamp"].(float64)), 0).UTC(),
			Level:     breadcrumbMap["level"],
			Message:   TrimLength(anyTypeToString(breadcrumbMap["message"]), 4096),
			Category:  TrimLength(anyTypeToString(breadcrumbMap["category"]), 256),
			EventID:   breadcrumbMap["event_id"],
		}
		if breadcrumb.Type == "" {
			breadcrumb.Type = "default"
		}
		if breadcrumb.Level == "info" {
			breadcrumb.Level = nil
		}
		if rawDataMap, ok := breadcrumbMap["data"].(map[string]interface{}); ok {
			dataMap := map[string]string{}
			breadcrumb.Data = dataMap
			for key, value := range rawDataMap {
				if stringValue, ok := value.(string); ok {
					dataMap[key] = stringValue
				} else {
					jsonValue, _ := json.Marshal(value)
					dataMap[key] = string(jsonValue)
				}
			}
		} else {
			// TODO seems like original `trim` function does not cast everything to string
			breadcrumb.Data = TrimLength(anyTypeToString(breadcrumbMap["data"]), 4096)
		}

		breadcrumbs.Values = append(breadcrumbs.Values, breadcrumb)
	}
	return nil
}

// canonical path: sentry.interfaces.Breadcrumbs
// path alias: breadcrumbs

// TrimLength truncates string up to length characters.
// If given string is shorter than length if will be returned without any changes.
// TODO move to the strings package
func TrimLength(v string, length int) string {
	if len(v) > length {
		return v[:length]
	}
	return v
}
