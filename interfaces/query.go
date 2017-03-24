package interfaces

type Query struct {
}

func (query *Query) DecodeRecord(record interface{}) error {
	return nil
}

func (query *Query) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("query", "sentry.interfaces.Query", request, query)
	return err
}
