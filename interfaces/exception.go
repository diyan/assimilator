package interfaces

type Exception struct {
}

func (exception *Exception) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (exception *Exception) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
