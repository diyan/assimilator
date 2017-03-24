package interfaces

type Repos struct {
}

func (repos *Repos) DecodeRecord(record interface{}) error {
	return nil
}

func (repos *Repos) DecodeRequest(request map[string]interface{}) error {
	err := DecodeRequest("repos", "sentry.interfaces.Repos", request, repos)
	return err
}
