package interfaces

type CSP struct {
}

func (csp *CSP) DecodeRecord(record interface{}) error {
	return nil
}

func (csp *CSP) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("csp", "sentry.interfaces.Csp", request, csp)
	return err
}
