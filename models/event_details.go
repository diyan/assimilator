package models

import (
	"fmt"
	"reflect"
	"time"

	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// TODO pick a better name
type Marshaler interface {
	//EncodeResponse() ([]byte, error)
	//EncodeRecord() (interface{}, error)
	DecodeRequest(map[string]interface{}) error
	DecodeRecord(interface{}) error
}

// TODO add fields release, message (getter of MessageInfo.Message),
//  size, dateReceived, entries (type message and type stacktrace),
//  context, contexts
type EventDetails struct {
	EventID      string `kv:"event_id" in:"event_id"`
	ProjectID    int
	Logger       string
	Platform     string
	Culprit      string
	Ref          int `kv:"_ref"`
	RefVersion   int `kv:"_ref_version"`
	Version      string
	Release      *string
	DateCreated  time.Time
	Fingerprint  []string
	Modules      map[string]interface{} // TODO ensure type is not map[string]string
	Type         string
	Size         int
	Errors       []EventError
	Tags         []TagKeyValue
	ReceivedTime time.Time `kv:"received"`
	Packages     map[string]string
	Metadata     map[string]string
	Extra        map[string]interface{} // TODO ensure type is not map[string]string
	// TODO Entries belongs to the API model
	//Entries         []interface{}               `json:"entries"`
	UserReport *string
}

func TimeDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t != reflect.TypeOf(time.Time{}) {
		return data, nil
	}
	if timeFloat, ok := data.(float64); ok {
		return time.Unix(int64(timeFloat), 0).UTC(), nil
	} else if timeString, ok := data.(string); ok {
		time, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return nil, err
		}
		return time, nil
	}
	return nil, fmt.Errorf("type is neither float64 nor string")
}

func TagsDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if t != reflect.TypeOf([]TagKeyValue{}) {
		return data, nil
	}
	tags := []TagKeyValue{}
	// Valid tags are both {"tagKey": "tagValue"} and [["tagKey", "tagValue"]]
	if tagsMap, ok := data.(map[string]interface{}); ok {
		for k, v := range tagsMap {
			// TODO check length of tag key and tag value
			tags = append(tags, TagKeyValue{
				Key: anyTypeToString(k), Value: anyTypeToString(v),
			})
		}
	} else if tagsSlice, ok := data.([]interface{}); ok {
		for _, tagBlob := range tagsSlice {
			// TODO safe type assertion
			tag := tagBlob.([]interface{})
			// TODO check length of tag key and tag value
			tags = append(tags, TagKeyValue{
				Key: anyTypeToString(tag[0]), Value: anyTypeToString(tag[1]),
			})
		}
	} else {
		return nil, fmt.Errorf("type is neither map[string]interface{} nor []interface{}")
	}
	return tags, nil
}

// TODO Hook works but looks like we have to traverse maps and slices
func PickleNoneDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if f != reflect.TypeOf(pickle.PickleNone{}) {
		return data, nil
	}
	//fmt.Printf("PickleNoneDecodeHook, f = %v, t = %v\n", f, t)
	return nil, nil
}

func StringMapDecodeHook(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	if !(f == reflect.TypeOf(map[interface{}]interface{}{}) &&
		t == reflect.TypeOf(map[string]interface{}{})) {
		return data, nil
	}
	return nil, nil
}

func anyTypeToString(v interface{}) string {
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}

func DecodeRecord(record interface{}, target interface{}) error {
	metadata := mapstructure.Metadata{}
	decodeHook := mapstructure.ComposeDecodeHookFunc(TimeDecodeHook, TagsDecodeHook, PickleNoneDecodeHook)
	config := mapstructure.DecoderConfig{
		DecodeHook:       decodeHook,
		Metadata:         &metadata,
		WeaklyTypedInput: false,
		TagName:          "kv",
		Result:           target,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return errors.Wrapf(err, "can not decode record from key/value node store")
	}
	err = decoder.Decode(record)
	return errors.Wrapf(err, "can not decode record from key/value node store")
}

func (event *EventDetails) DecodeRecord(record interface{}) error {
	if err := DecodeRecord(record, event); err != nil {
		return err
	}
	// TODO iterate unused keys, convert them to canonical interface path;
	//   if it's not interface - trackError
	//pp.Print("metadata.Unused", metadata.Unused)

	event.Size = 6597 // TODO remove hardcode
	// TOOD Entries is a field of API object
	//event.Entries = append(event.Entries, map[string]interface{}{
	//	"type": "message",
	//	"data": map[string]string{"message": rv.Message.Message},
	//})
	// TOOD Entries is a field of API object
	//event.Entries = append(rv.Entries, map[string]interface{}{
	//	"type": "stacktrace",
	//	"data": rv.Stacktrace,
	//})
	return nil
}

// TODO Consider use interface{} instead of map[string]interface{}
func DecodeRequest(request map[string]interface{}, target interface{}) error {
	metadata := mapstructure.Metadata{}
	decodeHook := mapstructure.ComposeDecodeHookFunc(TimeDecodeHook, TagsDecodeHook)
	config := mapstructure.DecoderConfig{
		DecodeHook:       decodeHook,
		Metadata:         &metadata,
		WeaklyTypedInput: true,
		TagName:          "in",
		Result:           target,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return errors.Wrapf(err, "can not parse request body")
	}
	err = decoder.Decode(request)
	return errors.Wrapf(err, "can not parse request body")
}

func (eventDetails *EventDetails) EncodeResponse() ([]byte, error) {
	// TODO map eventDetails root struct to the eventDetailsAPI
	// Is []byte correct return type?
	return nil, nil
}

type TagKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
