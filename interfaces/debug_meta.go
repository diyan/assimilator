package interfaces

type DebugMeta struct {
}

func init() {
	Register(&DebugMeta{})
}

func (*DebugMeta) KeyAlias() string {
	return "debug_meta"
}

func (*DebugMeta) KeyCanonical() string {
	return "sentry.interfaces.DebugMeta"
}
