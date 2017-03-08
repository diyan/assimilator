package interfaces

import (
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/pkg/errors"
)

type SDK struct {
	Name     string   `json:"name" pickle:"name"`
	Version  string   `json:"version" pickle:"version"`
	ClientIP string   `json:"clientIP" pickle:"client_ip"`
	Upstream Upstream `json:"upstream" pickle:"-"`
}

type Upstream struct {
	Name    string `json:"name" pickle:"-"`
	URL     string `json:"url" pickle:"-"`
	IsNewer bool   `json:"isNewer"`
}

func (sdk *SDK) UnmarshalRecord(nodeBlob interface{}) error {
	// TODO safe cast to map[interface{}]interface{}
	// TODO safe get from map using `sdk` alias key
	// TODO safe get from map using `sentry.interfaces.Sdk` canonical key
	if err := pickle.UnpackInto(&sdk).From(nodeBlob.(map[interface{}]interface{})["sdk"], nil); err != nil {
		return errors.Wrapf(err, "can not convert node blob to sentry.interfaces.Sdk")
	}
	sdk.Upstream.Name = sdk.Name                                // TODO check original code
	sdk.Upstream.URL = "https://docs.sentry.io/clients/python/" // TODO remove hardcode
	return nil
}

func (sdk *SDK) UnmarshalAPI(rawEvent map[string]interface{}) error {
	return nil
}
