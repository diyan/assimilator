package interfaces

type Message struct {
}

type messageRecord struct {
	Message string `pickle:"message"`
}

func (message *Message) UnmarshalRecord(nodeBlob interface{}) error {
	return nil
}

func (message *Message) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
