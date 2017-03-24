package interfaces

// LogEntry is an interface consisted of a ``message`` arg, an an optional
// ``params`` arg for formatting, and an optional ``formatted`` message which
// is the result of ``message`` combined with ``params``.
//
// If your message cannot be parameterized, then the message interface
// will serve no benefit.
//
// - ``message`` must be no more than 1000 characters in length.
//
// {
//     "message": "My raw message with interpreted strings like %s",
//     "formatted": "My raw message with interpreted strings like this",
//     "params": ["this"]
// }
type LogEntry struct {
	Message   string        `kv:"message"   in:"message"   json:"-"`
	Formatted string        `kv:"formatted" in:"formatted" json:"-"`
	Params    []interface{} `kv:"params"    in:"params"    json:"-"`
}

func (entry *LogEntry) DecodeRecord(record interface{}) error {
	return nil
}

func (entry *LogEntry) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("logentry", "sentry.interfaces.Message", request, entry)
	return err
}
