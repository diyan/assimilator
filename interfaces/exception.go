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
	Values          []ExceptionValue `in:"values" json:"values"`
	HasSystemFrames bool             `in:"-"      json:"hasSystemFrames"`
	ExcOmitted      bool             `in:"-"      json:"excOmitted"`
}

type ExceptionValue struct {
	Type            string      `in:"type"       json:"type"`
	Value           string      `in:"value"      json:"value"`
	Module          string      `in:"module"     json:"module"`
	Mechanism       interface{} `in:"mechanism"  json:"mechanism"`
	Stacktrace      Stacktrace  `in:"stacktrace" json:"stacktrace"`
	HasSystemFrames bool        `in:"-"          json:"hasSystemFrames"` // TODO ?
	SlimFrames      bool        `in:"-"          json:"slimFrames"`      // TODO ?
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
