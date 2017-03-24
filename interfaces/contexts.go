package interfaces

type Contexts struct{}

func (contexts *Contexts) DecodeRecord(record interface{}) error {
	return nil
}

func (contexts *Contexts) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("contexts", "sentry.interfaces.Contexts", request, contexts)
	return err
}
