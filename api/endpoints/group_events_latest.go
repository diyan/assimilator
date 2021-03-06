package api

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/interfaces"
	"github.com/diyan/assimilator/models"
)

type Event struct {
	models.Event
	models.EventDetails        `kv:",squash"`
	interfaces.EventInterfaces `kv:",squash"`
	PreviousEventID            *string `json:"previousEventID"`
	NextEventID                *string `json:"nextEventID"`
}

func GroupEventsLatestGetEndpoint(c context.Base) error {
	// TODO
	// 1. ? get default project to filter out issues by issue_id
	// 2. get latest event_id for issue_id that was provided in url segment
	// 3. call ProjectEventDetailsGetEndpoint and provide event_id
	projectStore := store.NewProjectStore()
	project, err := projectStore.GetProject(c.Tx, "acme-team", "acme")
	if err != nil {
		return err
	}
	eventID := 1
	// TODO move all code below to the ProjectEventDetailsGetEndpoint handler
	eventStore := store.NewEventStore()
	event, err := eventStore.GetEvent(c.Tx, project.ID, eventID)
	if err != nil {
		return err
	}
	apiEvent := Event{Event: *event}
	if event.DetailsRef != nil {
		detailsMap, err := eventStore.GetEventDetailsMap(c.Tx, *event.DetailsRef)
		if err != nil {
			return err
		}
		if err := models.DecodeRecord(detailsMap, &apiEvent); err != nil {
			return err
		}
		if err := apiEvent.EventDetails.DecodeRecord(detailsMap); err != nil {
			return err
		}
		if err := apiEvent.EventInterfaces.DecodeRecord(detailsMap); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, apiEvent)
	}
	return c.NoContent(http.StatusOK)
}
