package interfaces

type Template struct {
}

func init() {
	Register(&Template{})
}

func (*Template) KeyAlias() string {
	return "template"
}

func (*Template) KeyCanonical() string {
	return "sentry.interfaces.Template"
}
