package interfaces

type Query struct {
}

func init() {
	Register(&Query{})
}

func (*Query) KeyAlias() string {
	return "query"
}

func (*Query) KeyCanonical() string {
	return "sentry.interfaces.Query"
}
