package api

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"regexp"

	"github.com/diyan/assimilator/interfaces"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
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
	ProjectID                  int
	EventID                    string `in:"event_id"`
	models.EventDetails        `in:",squash"`
	interfaces.EventInterfaces `in:",squash"`
}

func bindRequest(project models.Project, requestBody io.ReadCloser, event *EventDetails) error {
	event.ProjectID = project.ID

	eventMap := map[string]interface{}{}
	if err := json.NewDecoder(requestBody).Decode(&eventMap); err != nil {
		return err
	}
	// TODO Ensure all keys are expected
	eventMap = interfaces.ToAliasKeys(eventMap)
	if err := models.DecodeRequest(eventMap, &event); err != nil {
		return err
	}
	if !validPlatforms[event.Platform] {
		event.Platform = "other"
	}
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
	//pp.Print(event)
	return c.JSON(200, map[string]string{"id": event.EventID})
}

func newUUIDHexString() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
