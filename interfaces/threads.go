package interfaces

type Threads struct {
}

func (threads *Threads) DecodeRecord(record interface{}) error {
	return nil
}

func (threads *Threads) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("threads", "sentry.interfaces.Threads", request, threads)
	return err
}
