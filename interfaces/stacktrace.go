package interfaces

import (
	"encoding/json"

	"github.com/AlekSi/pointer"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/pkg/errors"
)

type Stacktrace struct {
	HasSystemFrames bool    `node:"has_system_frames" json:"hasSystemFrames"`
	FramesOmitted   *bool   `node:"frames_omitted" json:"framesOmitted"` // TODO type?
	Frames          []Frame `node:"frames" json:"frames"`
}

type Frame struct {
	ColumnNumber       *int                        `node:"colno" json:"colNo"`
	LineNumber         int                         `node:"lineno" json:"lineNo"`
	InstructionOffset  *int                        `node:"-" json:"instructionOffset"`              // TODO type?
	InstructionAddress *string                     `node:"instruction_addr" json:"instructionAddr"` // TODO type?
	Symbol             *string                     `node:"symbol" json:"symbol"`                    // TODO type?
	SymbolAddress      *string                     `node:"symbol_addr" json:"symbolAddr"`           // TODO type?
	AbsolutePath       string                      `node:"abs_path" json:"absPath"`
	Module             string                      `node:"module" json:"module"`
	Package            *string                     `node:"package" json:"package"`
	Platform           *string                     `node:"platform" json:"platform"` // TODO type?
	Errors             *string                     `node:"errors" json:"errors"`     // TODO type?
	InApp              bool                        `node:"in_app" json:"inApp"`
	Filename           string                      `node:"filename" json:"filename"`
	Function           string                      `node:"function" json:"function"`
	Context            FrameContext                `node:"-" json:"context"`
	ContextLineNode    string                      `node:"context_line" json:"-"`
	PreContextNode     []string                    `node:"pre_context" json:"-"`
	PostContextNode    []string                    `node:"post_context" json:"-"`
	Variables          map[string]interface{}      `node:"-" json:"vars"`
	VariablesNode      map[interface{}]interface{} `node:"vars" json:"-"`
}

type FrameContext []FrameContextLine

type FrameContextLine struct {
	LineNumber int
	Line       string
}

func (contextLine FrameContextLine) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{contextLine.LineNumber, contextLine.Line})
}

func (stacktrace *Stacktrace) UnmarshalRecord(nodeBlob interface{}) error {
	err := DecodeRecord("stacktrace", "sentry.interfaces.Stacktrace", nodeBlob, stacktrace)
	for i := 0; i < len(stacktrace.Frames); i++ {
		frame := &stacktrace.Frames[i]
		//frame.InstructionAddress = padHexAddr(frame.InstructionAddress, padAddr)
		//frame.SymbolAddress = padHexAddr(frame.SymbolAddressRaw, padAddr)
		// TODO refactor getFrameContext into `decodeFrameContext(frame Frame)`
		frame.Context = getFrameContext(
			frame.LineNumber,
			frame.ContextLineNode,
			frame.PreContextNode,
			frame.PostContextNode,
			frame.Filename,
			frame.Module,
		)
		frame.ContextLineNode = ""
		frame.PreContextNode = nil
		frame.PostContextNode = nil
		frame.Variables = map[string]interface{}{}

		err := fillTypedVars(frame.VariablesNode, frame.Variables)
		if err != nil {
			return errors.Wrap(err, "failed to decode frame variables")
		}
		frame.VariablesNode = nil
	}
	return err
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
		case []interface{}, string, bool, nil:
			destMap[name] = value
		default:
			return errors.Errorf("unexpected type %T", value)
		}
	}
	return nil
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
