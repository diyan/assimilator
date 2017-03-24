package interfaces

type Template struct {
}

func (template *Template) DecodeRecord(record interface{}) error {
	return nil
}

func (template *Template) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("template", "sentry.interfaces.Template", request, template)
	return err
}
