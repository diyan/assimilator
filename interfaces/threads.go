package interfaces

type Threads struct {
}

func init() {
	Register(&Threads{})
}

func (*Threads) KeyAlias() string {
	return "thread"
}

func (*Threads) KeyCanonical() string {
	return "sentry.interfaces.Threads"
}
