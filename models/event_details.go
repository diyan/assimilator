package models

import (
	"time"

	"github.com/AlekSi/pointer"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/pkg/errors"
)

// TODO how about MarshalUnmarshaller name similar to ReadWriter?
type Marshaler interface {
	//MarshalAPI() ([]byte, error)
	//MarshalStore() (interface{}, error)
	UnmarshalAPI(map[string]interface{}) error
	UnmarshalRecord(interface{}) error
}

// TODO add fields release, message (getter of MessageInfo.Message),
//  size, dateReceived, entries (type message and type stacktrace),
//  context, contexts
type EventDetails struct {
	EventID      string
	ProjectID    int
	Logger       string
	Platform     string
	Culprit      string
	Ref          int
	RefVersion   int
	Version      string
	Release      *string
	DateCreated  time.Time
	Fingerprint  []string
	Modules      map[string]interface{} // TODO ensure type is not map[string]string
	Type         string
	Size         int
	Errors       []EventError
	Tags         []TagKeyValue
	ReceivedTime time.Time
	Packages     map[string]string
	Metadata     map[string]string
	Context      map[string]interface{} // TODO ensure type is not map[string]string
	// TODO move Entries to the API model
	//Entries         []interface{}               `json:"entries"`
	User       *string
	UserReport *string
}

type eventDetailsAPI struct {
	Release *string       `json:"release"`
	Type    string        `json:"type"`
	Size    int           `json:"size"`                   // TODO how to unpickle
	Errors  []EventError  `json:"errors" pickle:"errors"` // TODO type?
	Tags    []TagKeyValue `json:"tags"`
	//SDK          SDKInterface           `json:"sdk"`
	ReceivedTime time.Time              `json:"dateReceived"` // TODO check
	Packages     map[string]string      `json:"packages"`
	Metadata     map[string]string      `json:"metadata"`
	Context      map[string]interface{} `json:"context"`
	//Contexts     ContextsInterface      `json:"contexts"`
	Entries    []interface{} `json:"entries"`
	User       *string       `json:"user"`       // TODO type?
	UserReport *string       `json:"userReport"` // TODO type?
}

type eventDetailsRecord struct {
	Ref          int                         `pickle:"_ref"`
	RefVersion   int                         `pickle:"_ref_version"`
	Version      string                      `pickle:"version"`
	Release      *string                     `pickle:"release"`
	Type         string                      `pickle:"type"`
	Errors       []EventError                `pickle:"errors"` // TODO type?
	Tags         [][]string                  `pickle:"tags"`
	ReceivedTime float64                     `pickle:"received"`
	Packages     map[interface{}]interface{} `pickle:"modules"`
	Metadata     map[interface{}]interface{} `pickle:"metadata"`
	Extra        map[interface{}]interface{} `pickle:"extra"`
	Fingerprint  []string                    `pickle:"fingerprint"`
}

func (event *EventDetails) UnmarshalRecord(nodeBlob interface{}) error {
	record := eventDetailsRecord{}
	if err := pickle.UnpackInto(&record).From(nodeBlob, nil); err != nil {
		return errors.Wrapf(err, "can not convert node blob to event details")
	}
	event.Ref = record.Ref
	event.RefVersion = record.RefVersion
	event.Version = record.Version
	event.Release = record.Release
	event.Type = record.Type
	event.Errors = record.Errors
	event.Size = 6597 // TODO remove hardcode
	// TOOD Entries is a field of API object
	//event.Entries = append(event.Entries, map[string]interface{}{
	//	"type": "message",
	//	"data": map[string]string{"message": rv.Message.Message},
	//})
	for _, tag := range record.Tags {
		event.Tags = append(event.Tags, TagKeyValue{
			tag[0],
			tag[1],
		})
	}
	event.ReceivedTime = time.Unix(int64(record.ReceivedTime), 0).UTC()
	event.Packages = toStringMapString(record.Packages)
	event.Metadata = toStringMapString(record.Metadata)
	event.Context = toStringMap(record.Extra)
	event.Fingerprint = record.Fingerprint
	// TOOD Entries is a field of API object
	//rv.Entries = append(rv.Entries, map[string]interface{}{
	//	"type": "stacktrace",
	//	"data": rv.Stacktrace,
	//})
	return nil
}

func (eventDetails *EventDetails) MarshalAPI() ([]byte, error) {
	// TODO map eventDetails root struct to the eventDetailsAPI
	// Is []byte correct return type?
	return nil, nil
}

type TagKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// TODO toBool, toInt functions is copied between models and interfaces package
func toBool(value interface{}) bool {
	switch typedValue := value.(type) {
	case bool:
		return typedValue
	default:
		// TODO remove panic one all use-cases are checked
		panic(errors.Errorf("unexpected bool type %T", typedValue))
	}
}

func toInt(value interface{}) int {
	switch typedValue := value.(type) {
	case int64:
		return int(typedValue)
	case int:
		return typedValue
	default:
		// TODO remove panic one all use-cases are checked
		panic(errors.Errorf("unexpected int type %T", typedValue))
	}
}

func toIntPtr(value interface{}) *int {
	_, isPickleNone := value.(pickle.PickleNone)
	if value == nil || isPickleNone {
		return nil
	}
	return pointer.ToInt(toInt(value))
}

func toString(value interface{}) string {
	switch typedValue := value.(type) {
	case string:
		return typedValue
	default:
		// TODO remove panic one all use-cases are checked
		panic(errors.Errorf("unexpected string type %T", typedValue))
	}
}

func toStringPtr(value interface{}) *string {
	_, isPickleNone := value.(pickle.PickleNone)
	if value == nil || isPickleNone {
		return nil
	}
	return pointer.ToString(toString(value))
}

func toStringSlice(value interface{}) (rv []string) {
	if sliceValue, ok := value.([]interface{}); ok {
		for _, item := range sliceValue {
			rv = append(rv, toString(item))
		}
	}
	return
}

func toStringMapString(value interface{}) (rv map[string]string) {
	if mapValue, ok := value.(map[interface{}]interface{}); ok {
		rv = map[string]string{}
		for key, value := range mapValue {
			rv[toString(key)] = toString(value)
		}
	}
	return
}

func toStringMap(value interface{}) (rv map[string]interface{}) {
	if mapValue, ok := value.(map[interface{}]interface{}); ok {
		rv = map[string]interface{}{}
		for key, value := range mapValue {
			rv[toString(key)] = value
		}
	}
	return
}
