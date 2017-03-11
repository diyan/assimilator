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
	Values          []ExceptionValue `json:"values"`
	HasSystemFrames bool             `json:"hasSystemFrames"`
	ExcOmitted      bool             `json:"excOmitted"`
}

type ExceptionValue struct {
	Type            string      `json:"type"`
	Value           string      `json:"value"`
	Module          string      `json:"module"`
	Mechanism       interface{} `json:"mechanism"`
	Stacktrace      Stacktrace  `json:"stacktrace"`
	HasSystemFrames bool        `json:"hasSystemFrames"` // TODO ?
	SlimFrames      bool        `json:"slimFrames"`      // TODO ?
}

func (exception *Exception) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (exception *Exception) UnmarshalAPI(rawEvent map[string]interface{}) error {
	rawException, ok := rawEvent["exception"].(map[string]interface{})
	if !ok {
		rawException, ok = rawEvent["sentry.interfaces.Exception"].(map[string]interface{})
		if !ok {
			return nil
		}
	}
	// TODO validate that `values` is present and type is []interface{}
	hasSystemFrames := hasSystemFrames(rawException)
	for _, rawErrValue := range rawException["values"].([]interface{}) {
		errMap := rawErrValue.(map[string]interface{})
		errValue := ExceptionValue{
			HasSystemFrames: hasSystemFrames,
			SlimFrames:      false,
		}
		// TODO handle error
		errValue.Stacktrace.UnmarshalAPI(errMap["stacktrace"].(map[string]interface{}))
		exception.Values = append(exception.Values, errValue)
	}
	// TODO process `exc_omitted` if provided
	// TODO call `func slimExceptionData(exception *Exception)`
	return nil
}

//  TODO implement
func hasSystemFrames(v interface{}) bool {
	return false
}
