package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/AlekSi/pointer"
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
	ColumnNumber       *int                   `json:"colNo"`
	LineNumber         int                    `json:"lineNo"`
	InstructionOffset  *int                   `json:"instructionOffset"` // TODO type?
	InstructionAddress *string                `json:"instructionAddr"`   // TODO type?
	Symbol             *string                `json:"symbol"`            // TODO type?
	SymbolAddress      *string                `json:"symbolAddr"`        // TODO type?
	AbsolutePath       string                 `json:"absPath"`
	Module             string                 `json:"module"`
	Package            *string                `json:"package"`
	Platform           *string                `json:"platform"` // TODO type?
	Errors             *string                `json:"errors"`   // TODO type?
	InApp              bool                   `json:"inApp"`
	Filename           string                 `json:"filename"`
	Function           string                 `json:"function"`
	Context            FrameContext           `json:"context"`
	Variables          map[string]interface{} `json:"-"`
}

type FrameContext []FrameContextLine

type FrameContextLine struct {
	LineNumber int
	Line       string
}

func (contextLine FrameContextLine) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{contextLine.LineNumber, contextLine.Line})
}

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

func toStringStringMap(value interface{}) (rv map[string]string) {
	if mapValue, ok := value.(map[interface{}]interface{}); ok {
		rv := map[string]string{}
		for key, value := range mapValue {
			rv[toString(key)] = toString(value)
		}
	}
	return
}

func getFrameContext(
	lineNumber int, contextLine string, preContext, postContext []string,
	filename, module string) FrameContext {
	if lineNumber == 0 {
		return nil
	}
	if contextLine == "" && !(preContext != nil || postContext != nil) {
		return nil
	}
	context := FrameContext{}
	startLineNumber := lineNumber - len(preContext)
	atLineNumber := startLineNumber
	for _, line := range preContext {
		context = append(context, FrameContextLine{LineNumber: atLineNumber, Line: line})
		atLineNumber++
	}
	if startLineNumber < 0 {
		startLineNumber = 0
	}
	context = append(context, FrameContextLine{LineNumber: atLineNumber, Line: contextLine})
	atLineNumber++
	for _, line := range postContext {
		context = append(context, FrameContextLine{LineNumber: atLineNumber, Line: line})
		atLineNumber++
	}
	return context
}

func toEventDetails(nodeBlob interface{}) (rv *EventDetails, err error) {
	defer func() {
		if r := recover(); r != nil {
			rv = nil
			err = errors.Wrapf(r.(error), "can not convert node blob to event details")
		}
	}()
	rv = &EventDetails{}
	nodeMap := nodeBlob.(map[interface{}]interface{})
	rv.Ref = toInt(nodeMap["_ref"])
	rv.RefVersion = toInt(nodeMap["_ref_version"])
	rv.Version = toString(nodeMap["version"])
	rv.Type = toString(nodeMap["type"])
	rv.Message = MessageInfo{
		Message: toString(nodeMap["sentry.interfaces.Message"].(map[interface{}]interface{})["message"]),
	}
	rv.Entries = append(rv.Entries, map[string]interface{}{
		"type": "message",
		"data": rv.Message,
	})
	rv.Errors = nodeMap["errors"].([]interface{})
	for _, tagBlob := range nodeMap["tags"].([]interface{}) {
		tagSlice := tagBlob.([]interface{})
		rv.Tags = append(rv.Tags, TagInfo{
			Key:   tagSlice[0].(string),
			Value: tagSlice[1].(string),
		})
	}
	sdkMap := nodeMap["sdk"].(map[interface{}]interface{})
	rv.SDK = SDKInfo{
		Name:     toString(sdkMap["name"]),
		Version:  toString(sdkMap["version"]),
		ClientIP: toString(sdkMap["client_ip"]),
	}
	// TODO check `received` type, why it's float?
	rv.ReceivedTime = int(nodeMap["received"].(float64))
	rv.Packages = toStringStringMap(nodeMap["modules"])
	rv.Metadata = toStringStringMap(nodeMap["metadata"])
	rv.Extra = map[string]interface{}{}
	// TODO should use fillTypedVars method for `extra`?
	for key, value := range nodeMap["extra"].(map[interface{}]interface{}) {
		rv.Extra[toString(key)] = value
	}
	rv.Fingerprint = toStringSlice(nodeMap["fingerprint"])

	stacktrace := Stacktrace{}
	stacktraceMap := nodeMap["sentry.interfaces.Stacktrace"].(map[interface{}]interface{})
	stacktrace.HasSystemFrames = toBool(stacktraceMap["has_system_frames"])
	stacktrace.FramesOmitted = nil
	for _, frameBlob := range stacktraceMap["frames"].([]interface{}) {
		frameMap := frameBlob.(map[interface{}]interface{})
		frame := Frame{
			Filename:     toString(frameMap["filename"]),
			AbsolutePath: toString(frameMap["abs_path"]),
			Module:       toString(frameMap["module"]),
			Package:      toStringPtr(frameMap["package"]),
			Platform:     toStringPtr(frameMap["platform"]),
			//InstructionAddress: padHexAddr(frameMap["instruction_addr"], padAddr)
			//SymbolAddress: padHexAddr(frameMap["symbol_addr"], padAddr)
			Function:     toString(frameMap["function"]),
			Symbol:       toStringPtr(frameMap["symbol"]),
			LineNumber:   toInt(frameMap["lineno"]),
			ColumnNumber: toIntPtr(frameMap["colno"]),
			InApp:        toBool(frameMap["in_app"]),
			Variables:    map[string]interface{}{},
		}
		frame.Context = getFrameContext(
			frame.LineNumber,
			toString(frameMap["context_line"]),
			toStringSlice(frameMap["pre_context"]),
			toStringSlice(frameMap["post_context"]),
			frame.Filename,
			frame.Module,
		)

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
	return
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
			return err
		}
		return c.JSON(http.StatusOK, Event{Event: event, EventDetails: *eventDetails})
	}
	return c.NoContent(http.StatusOK)
}
