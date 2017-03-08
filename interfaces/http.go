package interfaces

type HTTP struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}

func (request *HTTP) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (request *HTTP) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
