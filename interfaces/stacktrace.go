package interfaces

import (
	"encoding/json"

	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/pkg/errors"
)

type Stacktrace struct {
	HasSystemFrames bool    `kv:"has_system_frames" in:"-"      json:"hasSystemFrames"`
	FramesOmitted   *bool   `kv:"frames_omitted"    in:"-"      json:"framesOmitted"` // TODO type?
	Frames          []Frame `kv:"frames"            in:"frames" json:"frames"`
}

// TODO https://docs.sentry.io/clientdev/interfaces/stacktrace/ image_addr symbol?
type Frame struct {
	ColumnNumber       *int    `kv:"colno"              in:"colno"              json:"colNo"`
	LineNumber         int     `kv:"lineno"             in:"lineno"             json:"lineNo"`
	InstructionOffset  *int    `kv:"instruction_offset" in:"instruction_offset" json:"instructionOffset"` // TODO type?
	InstructionAddress *string `kv:"instruction_addr"   in:"instruction_addr"   json:"instructionAddr"`   // TODO type?
	Symbol             *string `kv:"symbol"             in:"-"                  json:"symbol"`            // TODO type?
	SymbolAddress      *string `kv:"symbol_addr"        in:"symbol_addr"        json:"symbolAddr"`        // TODO type?
	AbsolutePath       string  `kv:"abs_path"           in:"abs_path"           json:"absPath"`
	Module             string  `kv:"module"             in:"module"             json:"module"`
	Package            *string `kv:"package"            in:"package"            json:"package"`
	Platform           *string `kv:"platform"           in:"platform"           json:"platform"` // TODO type?
	Errors             *string `kv:"errors"             in:"-"                  json:"errors"`   // TODO type?
	InApp              bool    `kv:"in_app"             in:"in_app"             json:"inApp"`
	Filename           string  `kv:"filename"           in:"filename"           json:"filename"`
	Function           string  `kv:"function"           in:"function"           json:"function"`

	Context         FrameContext                `kv:"-"            in:"-"            json:"context"`
	ContextLineNode string                      `kv:"context_line" in:"context_line" json:"-"`
	PreContextNode  []string                    `kv:"pre_context"  in:"pre_context"  json:"-"`
	PostContextNode []string                    `kv:"post_context" in:"post_context" json:"-"`
	Variables       map[string]interface{}      `kv:"-"            in:"vars"         json:"vars"`
	VariablesNode   map[interface{}]interface{} `kv:"vars"         in:"-"            json:"-"`
}

type FrameContext []FrameContextLine

type FrameContextLine struct {
	LineNumber int
	Line       string
}

func (*Stacktrace) KeyAlias() string {
	return "stacktrace"
}

func (*Stacktrace) KeyCanonical() string {
	return "sentry.interfaces.Stacktrace"
}

func (contextLine FrameContextLine) MarshalJSON() ([]byte, error) {
	return json.Marshal([]interface{}{contextLine.LineNumber, contextLine.Line})
}

func (stacktrace *Stacktrace) DecodeRecord(record interface{}) error {
	err := DecodeRecord(record, stacktrace)
	// TODO remove hardcoded value
	if stacktrace.FramesOmitted != nil && !*stacktrace.FramesOmitted {
		stacktrace.FramesOmitted = nil
	}
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

func (stacktrace *Stacktrace) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, stacktrace)
}
