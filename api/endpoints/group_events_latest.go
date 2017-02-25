package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"net/http"

	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type nodeInfo struct {
	NodeID string // base64 encoded UUID
}

type Event struct {
	models.Event
	EventDetails
	// TOOD add Message string `json:"message"` that will return MessageInfo.Message
	PreviousEventID *string `json:"previousEventID"`
	NextEventID     *string `json:"nextEventID"`
}

// TODO add fields release, message (getter of MessageInfo.Message),
//  size, dateReceived, entries (type message and type stacktrace),
//  context, contexts

type EventDetails struct {
	Ref        int         `json:"-"`
	RefVersion int         `json:"-"`
	Version    string      `json:"-"`
	Type       string      `json:"type"`
	Message    MessageInfo `json:"-"`
	// TODO check type info
	Errors       []interface{}          `json:"errors"`
	Tags         []TagInfo              `json:"tags"`
	SDK          SDKInfo                `json:"sdk"`
	ReceivedTime int                    `json:"received"`
	Packages     map[string]string      `json:"packages"`
	Extra        map[string]interface{} `json:"extra"`
	// TODO double type info
	Fingerprint []string `json:"fingerprint"`
	// TODO check type info
	Metadata map[string]string `json:"metadata"`
	//Stacktrace Stacktrace        `json:"sentry.interfaces.Stacktrace"`
	Entries []interface{} `json:"entries"`

	// TODO check type info for user, userReport
	User       *string `json:"user"`
	UserReport *string `json:"userReport"`
}

type MessageInfo struct {
	Message string `json:"message"`
}

type TagInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SDKInfo struct {
	Name     string       `json:"name"`
	Version  string       `json:"version"`
	ClientIP string       `json:"clientIP"`
	Upstream UpstreamInfo `json:"upstream,omitempty"`
}

type UpstreamInfo struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IsNewer bool   `json:"isNewer"`
}

type Stacktrace struct {
	HasSystemFrames bool `json:"hasSystemFrames"`
	// TODO check type
	FramesOmitted *bool   `json:"framesOmitted"`
	Frames        []Frame `json:"frames"`
}

type Frame struct {
	LineNumber   int                    `json:"lineNo"`
	AbsolutePath string                 `json:"absPath"`
	Module       string                 `json:"module"`
	InApp        bool                   `json:"inApp"`
	Filename     string                 `json:"filename"`
	Function     string                 `json:"function"`
	ContextLine  string                 `json:"context_line"`
	PreContext   []string               `json:"pre_context"`
	PostContext  []string               `json:"post_context"`
	Variables    map[string]interface{} `json:"vars"`
}

func toEventDetails(nodeBlob interface{}) (*EventDetails, error) {
	// TODO consider do all operation in unsafe manner + use recover inside this function
	rv := EventDetails{}
	nodeMap := nodeBlob.(map[interface{}]interface{})
	rv.Ref = int(nodeMap["_ref"].(int64))
	rv.RefVersion = int(nodeMap["_ref_version"].(int64))
	rv.Version = nodeMap["version"].(string)
	rv.Type = nodeMap["type"].(string)
	rv.Message = MessageInfo{
		Message: nodeMap["sentry.interfaces.Message"].(map[interface{}]interface{})["message"].(string),
	}
	rv.Entries = append(rv.Entries, map[string]interface{}{
		"type": "message",
		"data": rv.Message,
	})
	rv.Errors = nodeMap["errors"].([]interface{})
	for _, tagInterface := range nodeMap["tags"].([]interface{}) {
		tagSlice := tagInterface.([]interface{})
		rv.Tags = append(rv.Tags, TagInfo{
			Key:   tagSlice[0].(string),
			Value: tagSlice[1].(string),
		})
	}
	sdkMap := nodeMap["sdk"].(map[interface{}]interface{})
	rv.SDK = SDKInfo{
		Name:     sdkMap["name"].(string),
		Version:  sdkMap["version"].(string),
		ClientIP: sdkMap["client_ip"].(string),
	}
	// TODO check `received` type, why it's float?
	rv.ReceivedTime = int(nodeMap["received"].(float64))
	rv.Packages = map[string]string{}
	for name, version := range nodeMap["modules"].(map[interface{}]interface{}) {
		rv.Packages[name.(string)] = version.(string)
	}
	rv.Metadata = map[string]string{}
	for key, value := range nodeMap["metadata"].(map[interface{}]interface{}) {
		rv.Metadata[key.(string)] = value.(string)
	}
	rv.Extra = map[string]interface{}{}
	for key, value := range nodeMap["extra"].(map[interface{}]interface{}) {
		rv.Extra[key.(string)] = value
	}
	rv.Fingerprint = []string{}
	for _, item := range nodeMap["fingerprint"].([]interface{}) {
		rv.Fingerprint = append(rv.Fingerprint, item.(string))
	}

	stacktrace := Stacktrace{}
	stacktraceMap := nodeMap["sentry.interfaces.Stacktrace"].(map[interface{}]interface{})
	stacktrace.HasSystemFrames = stacktraceMap["has_system_frames"].(bool)
	stacktrace.FramesOmitted = nil
	for _, frameInterface := range stacktraceMap["frames"].([]interface{}) {
		frameMap := frameInterface.(map[interface{}]interface{})
		frame := Frame{
			LineNumber:   int(frameMap["lineno"].(int64)),
			AbsolutePath: frameMap["abs_path"].(string),
			Module:       frameMap["module"].(string),
			InApp:        frameMap["in_app"].(bool),
			Filename:     frameMap["filename"].(string),
			Function:     frameMap["function"].(string),
			ContextLine:  frameMap["context_line"].(string),
			Variables:    map[string]interface{}{},
		}
		if preContext, ok := frameMap["pre_context"].([]interface{}); ok {
			for _, preContextLine := range preContext {
				frame.PreContext = append(frame.PreContext, preContextLine.(string))
			}
		}
		if postContext, ok := frameMap["post_context"].([]interface{}); ok {
			for _, postContextLine := range postContext {
				frame.PostContext = append(frame.PostContext, postContextLine.(string))
			}
		}
		err := fillTypedVars(frameMap["vars"].(map[interface{}]interface{}), frame.Variables)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode frame variables")
		}
		stacktrace.Frames = append(stacktrace.Frames, frame)
	}
	rv.Entries = append(rv.Entries, map[string]interface{}{
		"type": "stacktrace",
		"data": stacktrace,
	})
	return &rv, nil
}

func fillTypedVars(sourceMap map[interface{}]interface{}, destMap map[string]interface{}) error {
	for nameInterface, value := range sourceMap {
		name := nameInterface.(string)
		switch typedValue := value.(type) {
		case map[interface{}]interface{}:
			nestedMap := map[string]interface{}{}
			destMap[name] = nestedMap
			return fillTypedVars(typedValue, nestedMap)
		case pickle.PickleNone:
			destMap[name] = nil
		case int64:
			destMap[name] = int(typedValue)
		case []interface{}, string, bool:
			destMap[name] = typedValue
		default:
			return errors.Errorf("unexpected type %T", typedValue)
		}
	}
	return nil
}

func unpickleZippedBase64String(blob string) (interface{}, error) {
	zippedBytes, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, errors.Wrap(err, "base64 decode failed")
	}
	zlibReader, err := zlib.NewReader(bytes.NewReader(zippedBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unzip stream failed")
	}
	defer zlibReader.Close()
	unpickledBlob, err := pickle.Unpickle(zlibReader)
	if err != nil {
		return nil, errors.Wrap(err, "unpickle failed")
	}
	return unpickledBlob, nil
}

func GroupEventsLatestGetEndpoint(c echo.Context) error {
	// TODO
	// 1. ? get default project to filter out issues by issue_id
	// 2. get latest event_id for issue_id that was provided in url segment
	// 3. call ProjectEventDetailsGetEndpoint and provide event_id
	projectStore := store.NewProjectStore(c)
	project, err := projectStore.GetProject("acme-team", "acme")
	if err != nil {
		return err
	}
	eventID := 1
	// TODO move all code below to the ProjectEventDetailsGetEndpoint handler
	eventStore := store.NewEventStore(c)
	event, err := eventStore.GetEvent(project.ID, eventID)
	if err != nil {
		return err
	}
	if event.Data != nil {
		nodeInfo, err := unpickleZippedBase64String(*event.Data)
		if err != nil {
			return errors.Wrap(err, "failed to decode event's node info")
		}
		// TODO do a safe type assertion for map and for map key
		// TODO do a safe get from the map
		// errors.New("failed to decode event data. 'node_id' key not found")
		// Can we just use Event.EventID field?
		nodeID := nodeInfo.(map[interface{}]interface{})["node_id"].(string)

		nodeBlobRow, err := eventStore.GetNodeBlob(nodeID)
		if err != nil {
			return errors.Wrap(err, "failed to load event's blob from node store")
		}
		nodeBlob, err := unpickleZippedBase64String(nodeBlobRow.Data)
		if err != nil {
			return errors.Wrap(err, "failed to decode event's blob")
		}
		eventDetails, err := toEventDetails(nodeBlob)
		if err != nil {
			return errors.Wrap(err, "can not convert node blob to event details")
		}
		return c.JSON(http.StatusOK, Event{Event: event, EventDetails: *eventDetails})
	}
	return c.NoContent(http.StatusOK)
}
