package interfaces

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/pkg/errors"
)

type Stacktrace struct {
	HasSystemFrames bool    `json:"hasSystemFrames"`
	FramesOmitted   *bool   `json:"framesOmitted"` // TODO type?
	Frames          []Frame `json:"frames"`
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
	Variables          map[string]interface{} `json:"vars"`
}

type FrameContext []FrameContextLine

type FrameContextLine struct {
	LineNumber int
	Line       string
}

type stacktraceRecord struct {
	HasSystemFrames bool          `pickle:"has_system_frames"`
	FramesOmitted   *bool         `pickle:"frames_omitted"` // TODO type?
	Frames          []frameRecord `pickle:"frames"`
}

type frameRecord struct {
	ColumnNumber       *int                        `pickle:"colno"`
	LineNumber         int                         `pickle:"lineno"`
	InstructionAddress *string                     `pickle:"instruction_addr"` // TODO type?
	Symbol             *string                     `pickle:"symbol"`           // TODO type?
	SymbolAddress      *string                     `pickle:"symbol_addr"`      // TODO type?
	AbsolutePath       string                      `pickle:"abs_path"`
	Module             string                      `pickle:"module"`
	Package            *string                     `pickle:"package"`
	Platform           *string                     `pickle:"platform"` // TODO type?
	Errors             *string                     `pickle:"errors"`   // TODO type?
	InApp              bool                        `pickle:"in_app"`
	Filename           string                      `pickle:"filename"`
	Function           string                      `pickle:"function"`
	ContextLine        string                      `pickle:"context_line"`
	PreContext         []string                    `pickle:"pre_context"`
	PostContext        []string                    `pickle:"post_context"`
	Variables          map[interface{}]interface{} `pickle:"vars"`
}

func (contextLine FrameContextLine) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{contextLine.LineNumber, contextLine.Line})
}

func (stacktrace *Stacktrace) UnmarshalRecord(nodeBlob interface{}) error {
	// TODO safe cast to map[interface{}]interface{}
	// TODO safe get from map using `stacktrace` alias key
	// TODO safe get from map using `sentry.interfaces.Stacktrace` canonical key
	record := stacktraceRecord{}
	if err := pickle.UnpackInto(&record).From(nodeBlob.(map[interface{}]interface{})["sentry.interfaces.Stacktrace"], nil); err != nil {
		return errors.Wrapf(err, "can not convert node blob to sentry.interfaces.Stacktrace")
	}
	for _, frameRecord := range record.Frames {
		frame := Frame{
			ColumnNumber:       frameRecord.ColumnNumber,
			LineNumber:         frameRecord.LineNumber,
			InstructionAddress: frameRecord.InstructionAddress,
			Symbol:             frameRecord.Symbol,
			SymbolAddress:      frameRecord.SymbolAddress,
			AbsolutePath:       frameRecord.AbsolutePath,
			Module:             frameRecord.Module,
			Package:            frameRecord.Package,
			Platform:           frameRecord.Platform,
			Errors:             frameRecord.Errors,
			InApp:              frameRecord.InApp,
			Filename:           frameRecord.Filename,
			Function:           frameRecord.Function,
		}
		//frame.InstructionAddress = padHexAddr(frameRecord.InstructionAddress, padAddr)
		//frame.SymbolAddress = padHexAddr(frameRecord.SymbolAddressRaw, padAddr)
		frame.Context = getFrameContext(
			frameRecord.LineNumber,
			frameRecord.ContextLine,
			frameRecord.PreContext,
			frameRecord.PostContext,
			frameRecord.Filename,
			frameRecord.Module,
		)
		frame.Variables = map[string]interface{}{}
		err := fillTypedVars(frameRecord.Variables, frame.Variables)
		if err != nil {
			return errors.Wrap(err, "failed to decode frame variables")
		}
		stacktrace.Frames = append(stacktrace.Frames, frame)
	}
	return nil
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

func fillTypedVars(sourceMap map[interface{}]interface{}, destMap map[string]interface{}) error {
	for nameBlob, valueBlob := range sourceMap {
		name := nameBlob.(string)
		switch value := valueBlob.(type) {
		case map[interface{}]interface{}:
			nestedMap := map[string]interface{}{}
			destMap[name] = nestedMap
			if err := fillTypedVars(value, nestedMap); err != nil {
				return err
			}
		case pickle.PickleNone:
			destMap[name] = nil
		case int64:
			destMap[name] = int(value)
		case []interface{}, string, bool:
			destMap[name] = value
		default:
			return errors.Errorf("unexpected type %T", value)
		}
	}
	return nil
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

// NOTE We are expecting here rawStacktrace instead of rawEvent
func (stacktrace *Stacktrace) UnmarshalAPI(rawStacktrace map[string]interface{}) error {
	rawFrames, ok := rawStacktrace["frames"].([]interface{})
	if !ok {
		return nil
	}
	for _, rawFrame := range rawFrames {
		frameMap := rawFrame.(map[string]interface{})
		frame := Frame{
			Filename:     frameMap["filename"].(string),
			Function:     frameMap["function"].(string),
			LineNumber:   int(frameMap["lineno"].(float64)),
			ColumnNumber: pointer.ToInt(int(frameMap["colno"].(float64))),
			InApp:        frameMap["in_app"].(bool),
		}
		stacktrace.Frames = append(stacktrace.Frames, frame)
	}

	return nil
}
