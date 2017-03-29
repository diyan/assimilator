package interfaces

// The SDK used to transmit this event.
//
// {
//     "name": "sentry-unity",
//     "version": "1.0"
// }
type SDK struct {
	Name     string   `kv:"name"      in:"name"    json:"name"`
	Version  string   `kv:"version"   in:"version" json:"version"`
	ClientIP string   `kv:"client_ip" in:"-"       json:"clientIP"`
	Upstream Upstream `json:"upstream"  in:"-"`
}

type Upstream struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IsNewer bool   `json:"isNewer"`
}

func init() {
	Register(&SDK{})
}

func (*SDK) KeyAlias() string {
	return "sdk"
}

func (*SDK) KeyCanonical() string {
	return "sentry.interfaces.Sdk"
}

func (sdk *SDK) DecodeRecord(record map[string]interface{}) error {
	sdk.Upstream.Name = sdk.Name                                // TODO check original code
	sdk.Upstream.URL = "https://docs.sentry.io/clients/python/" // TODO remove hardcode
	return nil
}
