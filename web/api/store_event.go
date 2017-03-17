package api

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"regexp"

	"github.com/diyan/assimilator/interfaces"
	"github.com/diyan/assimilator/models"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func storeGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

var eventIDRegex = regexp.MustCompile("^[a-fA-F0-9]{32}$")

var validPlatforms = map[string]bool{
	"as3":        true,
	"c":          true,
	"cfml":       true,
	"cocoa":      true,
	"csharp":     true,
	"go":         true,
	"java":       true,
	"javascript": true,
	"node":       true,
	"objc":       true,
	"other":      true,
	"perl":       true,
	"php":        true,
	"python":     true,
	"ruby":       true,
	"elixir":     true,
	"haskell":    true,
	"groovy":     true,
}

type EventDetails struct {
	models.EventDetails
	interfaces.EventInterfaces
}

// TODO implement interfaces for breadcrumbs, request, user (used by JavaScript client)
func bindRequest(project models.Project, requestBody io.ReadCloser, event *EventDetails) error {
	event.ProjectID = project.ID

	rawEvent := map[string]interface{}{}
	if err := json.NewDecoder(requestBody).Decode(&rawEvent); err != nil {
		return err
	}
	// Ensure all keys are expected
	// Bind event attributes
	// Bind event interfaces

	metadata := mapstructure.Metadata{}
	decodeHook := mapstructure.ComposeDecodeHookFunc(models.TimeDecodeHook, models.TagsDecodeHook)
	config := mapstructure.DecoderConfig{
		DecodeHook:       decodeHook,
		Metadata:         &metadata,
		WeaklyTypedInput: false,
		Result:           &event.EventDetails,
	}
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		pp.Print(err)
		return errors.Wrapf(err, "can not parse request body to event details")
	}
	if err := decoder.Decode(rawEvent); err != nil {
		pp.Print(err)
		errors.Wrapf(err, "can not parse request body to event details")
	}
	if !validPlatforms[event.Platform] {
		event.Platform = "other"
	}
	//pp.Print(metadata.Unused)
	//pp.Print(event)

	// TODO handle error
	event.EventInterfaces.UnmarshalAPI(rawEvent)
	return nil
}

func storePostView(c echo.Context) error {
	// TODO move GetProject to the package shared between api/endpoints and web/api
	//project := GetProject(c)
	project := models.Project{ID: 1}
	event := EventDetails{}
	if err := bindRequest(project, c.Request().Body, &event); err != nil {
		return err
	}
	return c.JSON(200, map[string]string{"id": event.EventID})
}

func newUUIDHexString() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
