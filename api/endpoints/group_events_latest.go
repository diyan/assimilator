package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"net/http"

	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/interfaces"
	"github.com/diyan/assimilator/lib/conv"
	"github.com/diyan/assimilator/models"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Event struct {
	models.Event
	models.EventDetails        `kv:",squash"`
	interfaces.EventInterfaces `kv:",squash"`
	PreviousEventID            *string `json:"previousEventID"`
	NextEventID                *string `json:"nextEventID"`
}

type NodeRef struct {
	NodeID string `kv:"node_id"`
}

func unpickleZippedBase64String(blob string) (map[string]interface{}, error) {
	zippedBytes, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, errors.Wrap(err, "base64 decode failed")
	}
	zlibReader, err := zlib.NewReader(bytes.NewReader(zippedBytes))
	if err != nil {
		return nil, errors.Wrap(err, "unzip stream failed")
	}
	defer zlibReader.Close()
	unpickledBlob, err := pickle.Unpickle(zlibReader)
	if err != nil {
		return nil, errors.Wrap(err, "unpickle failed")
	}
	unpickledMap := conv.StringMap(unpickledBlob)
	return unpickledMap, nil
}

func GroupEventsLatestGetEndpoint(c echo.Context) error {
	// TODO
	// 1. ? get default project to filter out issues by issue_id
	// 2. get latest event_id for issue_id that was provided in url segment
	// 3. call ProjectEventDetailsGetEndpoint and provide event_id
	projectStore := store.NewProjectStore(c)
	project, err := projectStore.GetProject("acme-team", "acme")
	if err != nil {
		return err
	}
	eventID := 1
	// TODO move all code below to the ProjectEventDetailsGetEndpoint handler
	eventStore := store.NewEventStore(c)
	event, err := eventStore.GetEvent(project.ID, eventID)
	if err != nil {
		return err
	}
	if event.Data != nil {
		nodeRefMap, err := unpickleZippedBase64String(*event.Data)
		if err != nil {
			return errors.Wrap(err, "failed to decode event's node reference")
		}
		nodeRef := NodeRef{}
		if err := models.DecodeRecord(nodeRefMap, &nodeRef); err != nil {
			return errors.Wrap(err, "failed to decode event's node reference")
		}
		nodeBlobRow, err := eventStore.GetNodeBlob(nodeRef.NodeID)
		if err != nil {
			return errors.Wrap(err, "failed to load event's blob from node store")
		}
		eventMap, err := unpickleZippedBase64String(nodeBlobRow.Data)
		if err != nil {
			return errors.Wrap(err, "failed to decode event's blob")
		}
		apiEvent := Event{Event: event}
		// TODO we can hide DecodeRecord method inside eventStore
		//   but we need this convention for interfaces
		if err := models.DecodeRecord(interfaces.ToAliasKeys(eventMap), &apiEvent); err != nil {
			return err
		}
		if err := apiEvent.EventDetails.DecodeRecord(eventMap); err != nil {
			return err
		}
		if err := apiEvent.EventInterfaces.DecodeRecord(eventMap); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, apiEvent)
	}
	return c.NoContent(http.StatusOK)
}
