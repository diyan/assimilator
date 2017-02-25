package store

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type EventStore struct {
	c echo.Context
}

func NewEventStore(c echo.Context) EventStore {
	return EventStore{c: c}
}

func (s EventStore) GetEvent(projectID, eventID int) (models.Event, error) {
	db, err := db.FromE(s.c)
	event := models.Event{}
	if err != nil {
		return event, errors.Wrap(err, "can not get issue event")
	}
	_, err = db.SelectBySql(`
            select m.*
                from sentry_message m
            where m.project_id  = ? and m.id = ?`,
		projectID, eventID).
		LoadStructs(&event)
	if err != nil {
		return event, errors.Wrap(err, "can not get issue event")
	}
	return event, nil
}

// TODO move to the nodeStore
func (s EventStore) GetNodeBlob(nodeID string) (models.NodeBlob, error) {
	db, err := db.FromE(s.c)
	nodeBlob := models.NodeBlob{}
	if err != nil {
		return nodeBlob, errors.Wrap(err, "can not get node blob")
	}
	_, err = db.SelectBySql(`
            select n.*
                from nodestore_node n
            where n.id = ?`,
		nodeID).
		LoadStructs(&nodeBlob)
	if err != nil {
		return nodeBlob, errors.Wrap(err, "can not get node blob")
	}
	return nodeBlob, nil
}

func (s EventStore) SaveEvent(event models.Event) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save issue event")
	}
	_, err = db.InsertInto("sentry_message").
		Columns("id", "group_id", "message_id", "project_id", "message",
			"platform", "time_spent", "data", "datetime").
		Record(event).
		Exec()
	return errors.Wrap(err, "failed to save issue event")
}

// TODO move to the nodeStore
func (s EventStore) SaveNodeBlob(nodeBlob models.NodeBlob) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save node blob")
	}
	_, err = db.InsertInto("nodestore_node").
		Columns("id", "data", "timestamp").
		Record(nodeBlob).
		Exec()
	return errors.Wrap(err, "failed to save node blob")
}
