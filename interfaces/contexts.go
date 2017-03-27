package interfaces

type Contexts struct{}

func (*Contexts) KeyAlias() string {
	return "contexts"
}

func (*Contexts) KeyCanonical() string {
	return "sentry.interfaces.Contexts"
}

func (contexts *Contexts) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, contexts)
}

func (contexts *Contexts) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, contexts)
}
