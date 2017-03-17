package interfaces

type SDK struct {
	Name     string   `json:"name" node:"name"`
	Version  string   `json:"version" node:"version"`
	ClientIP string   `json:"clientIP" node:"client_ip"`
	Upstream Upstream `json:"upstream"`
}

type Upstream struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IsNewer bool   `json:"isNewer"`
}

func (sdk *SDK) UnmarshalRecord(nodeBlob interface{}) error {
	err := DecodeRecord("sdk", "sentry.interfaces.Sdk", nodeBlob, sdk)
	sdk.Upstream.Name = sdk.Name                                // TODO check original code
	sdk.Upstream.URL = "https://docs.sentry.io/clients/python/" // TODO remove hardcode
	return err
}

func (sdk *SDK) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
