package interfaces

type Query struct {
}

func (*Query) KeyAlias() string {
	return "query"
}

func (*Query) KeyCanonical() string {
	return "sentry.interfaces.Query"
}

func (query *Query) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, query)
}

func (query *Query) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, query)
}
