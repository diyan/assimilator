package interfaces

type CSP struct {
}

func (*CSP) KeyAlias() string {
	return "csp"
}

func (*CSP) KeyCanonical() string {
	return "sentry.interfaces.Csp"
}

func (csp *CSP) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, csp)
}

func (csp *CSP) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, csp)
}
