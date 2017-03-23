package interfaces

// Exception consists of a list of values. In most cases, this list
// contains a single exception, with an optional stacktrace interface.
//
// Each exception has a mandatory `value` argument and optional `type` and
// `module` arguments describing the exception class type and module
// namespace.
//
// You can also optionally bind a stacktrace interface to an exception. The
// spec is identical to `sentry.interfaces.Stacktrace`.
//
// {
//     "values": [{
//         "type": "ValueError",
//         "value": "My exception value",
//         "module": "__builtins__",
//         "mechanism": {},
//         "stacktrace": {
//             // see sentry.interfaces.Stacktrace
//         }
//     }]
// }
//
//    Values should be sent oldest to newest, this includes both the stacktrace
//    and the exception itself.
type Exception struct {
	Values          []ExceptionValue `input:"values" json:"values"`
	HasSystemFrames bool             `input:"-"      json:"hasSystemFrames"`
	ExcOmitted      bool             `input:"-"      json:"excOmitted"`
}

type ExceptionValue struct {
	Type            string      `input:"type"       json:"type"`
	Value           string      `input:"value"      json:"value"`
	Module          string      `input:"module"     json:"module"`
	Mechanism       interface{} `input:"mechanism"  json:"mechanism"`
	Stacktrace      Stacktrace  `input:"stacktrace" json:"stacktrace"`
	HasSystemFrames bool        `input:"-"          json:"hasSystemFrames"` // TODO ?
	SlimFrames      bool        `input:"-"          json:"slimFrames"`      // TODO ?
}

func (exception *Exception) DecodeRecord(record interface{}) error {
	return nil
}

func (exception *Exception) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("exception", "sentry.interfaces.Exception", request, exception)
	// hasSystemFrames := hasSystemFrames(rawException)
	// TODO iterate values, set HasSystemFrames
	// TODO process `exc_omitted` if provided
	// TODO call `func slimExceptionData(exception *Exception)`
	return err
}

//  TODO implement
func hasSystemFrames(v interface{}) bool {
	return false
}
