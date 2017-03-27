package interfaces

type Threads struct {
}

func (*Threads) KeyAlias() string {
	return "thread"
}

func (*Threads) KeyCanonical() string {
	return "sentry.interfaces.Threads"
}

func (threads *Threads) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, threads)

}

func (threads *Threads) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, threads)
}
