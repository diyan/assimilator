package interfaces

type HTTP struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}

func (request *HTTP) DecodeRecord(record interface{}) error {
	return nil
}

func (request *HTTP) DecodeRequest(rawRequest map[string]interface{}) error {
	return nil
}
