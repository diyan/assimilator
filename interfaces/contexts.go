package interfaces

type Contexts struct{}

// alias name is `contexts`
func (contexts *Contexts) DecodeRecord(record interface{}) error {
	return nil
}

func (contexts *Contexts) DecodeRequest(request map[string]interface{}) error {
	return nil
}
