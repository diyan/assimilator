package interfaces

type HTTP struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}

func (*HTTP) KeyAlias() string {
	return "request"
}

func (*HTTP) KeyCanonical() string {
	return "sentry.interfaces.Http"
}

func (request *HTTP) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, request)
}

func (request *HTTP) DecodeRequest(rawRequest map[string]interface{}) error {
	return DecodeRequest(rawRequest, request)
}
