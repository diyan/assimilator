package interfaces

type CSP struct {
}

func init() {
	Register(&CSP{})
}

func (*CSP) KeyAlias() string {
	return "csp"
}

func (*CSP) KeyCanonical() string {
	return "sentry.interfaces.Csp"
}
