package interfaces

type Template struct {
}

func (*Template) KeyAlias() string {
	return "template"
}

func (*Template) KeyCanonical() string {
	return "sentry.interfaces.Template"
}

func (template *Template) DecodeRecord(record interface{}) error {
	return DecodeRecord(record, template)

}

func (template *Template) DecodeRequest(request map[string]interface{}) error {
	return DecodeRequest(request, template)
}
