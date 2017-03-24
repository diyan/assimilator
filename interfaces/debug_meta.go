package interfaces

type DebugMeta struct {
}

func (debug *DebugMeta) DecodeRecord(record interface{}) error {
	return nil
}

func (debug *DebugMeta) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("debug_meta", "sentry.interfaces.DebugMeta", request, debug)
	return err
}
