package interfaces

type HTTP struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}

func init() {
	Register(&HTTP{})
}

func (*HTTP) KeyAlias() string {
	return "request"
}

func (*HTTP) KeyCanonical() string {
	return "sentry.interfaces.Http"
}
