package interfaces

type DebugMeta struct {
}

func (*DebugMeta) KeyAlias() string {
	return "debug_meta"
}

func (*DebugMeta) KeyCanonical() string {
	return "sentry.interfaces.DebugMeta"
}

func (debug *DebugMeta) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, debug)
}

func (debug *DebugMeta) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, debug)
}
