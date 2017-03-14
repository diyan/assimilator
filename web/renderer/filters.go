package renderer

import (
	"encoding/json"

	"github.com/flosch/pongo2"
)

// RegisterFilters ...
func RegisterFilters() {
	pongo2.RegisterFilter("to_json", filterToJSON)
	pongo2.RegisterFilter("split", noOpFilter)
	pongo2.RegisterFilter("list_organizations", noOpFilter)
}

func filterToJSON(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	bytes, err := json.Marshal(in.String())
	if err != nil {
		// TODO consider cast err to pongo2.Error
		panic(err)
	}
	// TODO consider use pongo2 API to convert from []byte
	return pongo2.AsValue(string(bytes)), nil
}

func noOpFilter(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	return nil, nil
}
