package interfaces

type SDK struct {
	Name     string   `node:"name"      json:"name"`
	Version  string   `node:"version"   json:"version"`
	ClientIP string   `node:"client_ip" json:"clientIP"`
	Upstream Upstream `json:"upstream"`
}

type Upstream struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IsNewer bool   `json:"isNewer"`
}

func (sdk *SDK) DecodeRecord(record interface{}) error {
	err := DecodeRecord("sdk", "sentry.interfaces.Sdk", record, sdk)
	sdk.Upstream.Name = sdk.Name                                // TODO check original code
	sdk.Upstream.URL = "https://docs.sentry.io/clients/python/" // TODO remove hardcode
	return err
}

func (sdk *SDK) DecodeRequest(request map[string]interface{}) error {
	return nil
}
