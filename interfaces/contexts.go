package interfaces

type Contexts struct{}

func init() {
	Register(&Contexts{})
}

func (*Contexts) KeyAlias() string {
	return "contexts"
}

func (*Contexts) KeyCanonical() string {
	return "sentry.interfaces.Contexts"
}
