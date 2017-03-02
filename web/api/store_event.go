package api

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/models"
	"github.com/k0kubun/pp"
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
	EventID     string
	ProjectID   int
	Logger      string
	Platform    string
	Culprit     string
	Release     string
	Timestamp   time.Time
	Fingerprint []string
	Modules     map[string]interface{} // TODO ensure type is not map[string]string
	Tags        []TagKeyValue
	Extra       map[string]interface{} // TODO ensure type is not map[string]string
	Errors      []models.EventError

	// Built-in interfaces, sorted by name
	AppleCrashReport interface{}
	Breadcrumbs      interface{}
	Contexts         interface{}
	CSP              interface{}
	DebugMeta        interface{}
	Device           interface{}
	Exception        interface{}
	LogEntry         interface{}
	Query            interface{}
	Repos            interface{}
	Request          interface{}
	SDK              interface{}
	Stacktrace       interface{}
	Template         interface{}
	Threads          interface{}
	User             interface{}
}

// TODO consider move to the interfaces package
type TagKeyValue struct {
	Key, Value string
}

// TODO move to the interfaces package
type RequestInfo struct {
	Headers map[string]string `json:"headers"`
	URL     string            `json:"url"`
}

type ExceptionInfo struct {
}

type BreadcrumbInfo struct {
}

type UserInfo struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	IPAddress string `json:"ip_address"`
}

func anyTypeToString(v interface{}) string {
	if v != nil {
		return fmt.Sprint(v)
	}
	return ""
}

func simpleTypeToString(v interface{}) (rv string, ok bool) {
	ok = true
	switch v.(type) {
	case string, bool, float64: // JSON basic types
		rv = fmt.Sprint(v)
	case nil:
		rv = ""
	default:
		ok = false
	}
	return
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

	if rawEventID, ok := rawEvent["event_id"]; ok {
		if eventID, ok := rawEventID.(string); ok {
			if eventIDRegex.MatchString(eventID) {
				event.EventID = eventID
			} else {
				log.WithFields(log.Fields{
					"event_id":     event.EventID,
					"event_id_len": len(event.EventID),
				}).Debug("Discarded invalid value for event_id")
				event.Errors = append(event.Errors, models.EventError{
					Type:  models.EventErrorInvalidData,
					Name:  "event_id",
					Value: event.EventID,
				})
				event.EventID = newUUIDHexString()
			}
		} else {
			return errors.New("Invalid value for event_id")
		}
	} else {
		event.EventID = newUUIDHexString()
	}

	if rawCulprit, ok := rawEvent["culprit"]; ok {
		if culprit, ok := rawCulprit.(string); ok {
			event.Culprit = culprit
		} else {
			return errors.New("Invalid value for culprit")
		}
	}

	event.Logger = anyTypeToString(rawEvent["logger"]) // TODO check code
	event.Release = anyTypeToString(rawEvent["release"])

	if !validPlatforms[anyTypeToString(rawEvent["platform"])] {
		event.Platform = "other"
	}

	if timestamp, ok := rawEvent["timestamp"]; ok {
		if err := bindTimestamp(timestamp, event); err != nil {
			log.WithFields(log.Fields{
				"timestamp": timestamp,
				"err":       err,
			}).Debug("Discarded invalid value for timestamp")
			event.Errors = append(event.Errors, models.EventError{
				Type:  models.EventErrorInvalidData,
				Name:  "timestamp",
				Value: timestamp,
			})
		}
	}

	if rawFingerprint, ok := rawEvent["fingerprint"]; ok {
		invalidFingerprint := false
		if fingerprintSlice, ok := rawFingerprint.([]interface{}); ok {
			for _, fingerprintLine := range fingerprintSlice {
				if value, ok := simpleTypeToString(fingerprintLine); ok {
					event.Fingerprint = append(event.Fingerprint, anyTypeToString(value))
				} else {
					invalidFingerprint = true
				}
				switch value := fingerprintLine.(type) {

				case string, bool, float64:
					event.Fingerprint = append(event.Fingerprint, anyTypeToString(value))
				default:
					invalidFingerprint = true
				}
			}
		} else {
			invalidFingerprint = true
		}
		if invalidFingerprint {
			log.WithFields(log.Fields{
				"fingerprint": rawFingerprint,
				"err":         "array of booleans, numbers, strings is expected",
			}).Debug("Discarded invalid value for fingerprint")
			event.Errors = append(event.Errors, models.EventError{
				Type:  models.EventErrorInvalidData,
				Name:  "fingerprint",
				Value: rawFingerprint,
			})
			event.Fingerprint = nil // all or nothing
		}
	}

	if rawModules, ok := rawEvent["modules"]; ok {
		if modules, ok := rawModules.(map[string]interface{}); ok {
			event.Modules = modules
		} else {
			log.WithFields(log.Fields{
				"modules": modules,
				"err":     "type is not map[string]interface{}",
			}).Debug("Discarded invalid value for modules")
			event.Errors = append(event.Errors, models.EventError{
				Type:  models.EventErrorInvalidData,
				Name:  "modules",
				Value: rawModules,
			})
		}
	}

	if rawExtra, ok := rawEvent["extra"]; ok {
		if extra, ok := rawExtra.(map[string]interface{}); ok {
			event.Extra = extra
		} else {
			log.WithFields(log.Fields{
				"extra": extra,
				"err":   "type is not map[string]interface{}",
			}).Debug("Discarded invalid value for extra")
			event.Errors = append(event.Errors, models.EventError{
				Type:  models.EventErrorInvalidData,
				Name:  "extra",
				Value: rawExtra,
			})
		}
	}

	// Valid tags are both {"tagKey": "tagValue"} and [["tagKey", "tagValue"]]
	if rawTags, ok := rawEvent["tags"]; ok {
		if tagsMap, ok := rawTags.(map[string]interface{}); ok {
			// TODO process tags as map
			pp.Print(tagsMap)
		} else if tagsSlice, ok := rawTags.([]interface{}); ok {
			// TODO process tags as slice
			pp.Print(tagsSlice)
		} else {
			log.WithFields(log.Fields{
				"tags": rawTags,
				"err":  "type is neither map[string]interface{} nor []interface{}",
			}).Debug("Discarded invalid value for tags")
			event.Errors = append(event.Errors, models.EventError{
				Type:  models.EventErrorInvalidData,
				Name:  "tags",
				Value: rawTags,
			})
		}
	}

	pp.Print(event)
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

// TODO timestamp could be either float or string
func bindTimestamp(timestamp interface{}, event *EventDetails) error {
	return nil
}

func newUUIDHexString() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
