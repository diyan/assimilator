package interfaces

type Contexts struct{}

// alias name is `contexts`
func (contexts *Contexts) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (contexts *Contexts) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
