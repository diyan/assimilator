package interfaces

type Message struct {
}

func (message *Message) DecodeRecord(record interface{}) error {
	return nil
}

func (message *Message) DecodeRequest(request map[string]interface{}) error {
	return nil
}
