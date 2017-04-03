package store

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"

	"github.com/AlekSi/pointer"
	"github.com/diyan/assimilator/interfaces"
	"github.com/diyan/assimilator/lib/conv"
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	pickle "github.com/hydrogen18/stalecucumber"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

type EventStore struct {
}

func NewEventStore() EventStore {
	return EventStore{}
}

func (s EventStore) GetEvent(tx *dbr.Tx, projectID, eventID int) (*models.Event, error) {
	event := models.Event{}
	_, err := tx.SelectBySql(`
            select m.*
                from sentry_message m
            where m.project_id  = ? and m.id = ?`,
		projectID, eventID).
		LoadStructs(&event)
	if err != nil {
		return nil, errors.Wrap(err, "can not get issue event")
	}
	if event.DetailsRefRaw != nil {
		nodeRefMap, err := unpickleZippedBase64String(*event.DetailsRefRaw)
		if err != nil {
			return nil, errors.Wrap(err, "can not get issue event: failed to decode reference to the event details")
		}
		event.DetailsRef = &models.NodeRef{}
		if err := models.DecodeRecord(nodeRefMap, event.DetailsRef); err != nil {
			return nil, errors.Wrap(err, "can not get issue event: failed to decode reference to the event details")
		}
		event.DetailsRefRaw = nil
	}
	return &event, nil
}

func (s EventStore) GetEventDetailsMap(tx *dbr.Tx, nodeRef models.NodeRef) (map[string]interface{}, error) {
	nodeBlob := models.NodeBlob{}
	_, err := tx.SelectBySql(`
            select n.*
                from nodestore_node n
            where n.id = ?`,
		nodeRef.NodeID).
		LoadStructs(&nodeBlob)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load event details from node store")
	}
	eventMap, err := unpickleZippedBase64String(nodeBlob.Data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode event details blob")
	}
	// TODO it's a bad idea to import interfaces from db/store
	return interfaces.ToAliasKeys(eventMap), nil
}

func (s EventStore) SaveEvent(tx *dbr.Tx, event models.Event) error {
	if event.DetailsRef != nil {
		// TODO extract into method v, err := toBase64ZipPickleString(*event.DetailsRef)
		pickleBuffer := bytes.Buffer{} // io.Writer
		// TODO how to re-use `kv` tag for pickler which uses `pickle` tag name?
		//detailsMap := map[string]interface{}{
		//	"node_id": event.DetailsRef.NodeID,
		//}
		//_, err := pickle.NewPickler(&pickleBuffer).Pickle(detailsMap)
		_, err := pickle.NewPickler(&pickleBuffer).Pickle(event.DetailsRef)
		if err != nil {
			return errors.Wrap(err, "pickle failed")
		}
		zlibBuffer := bytes.Buffer{} // io.Writer
		zlibWriter := zlib.NewWriter(&zlibBuffer)
		//defer zlibWriter.Close()
		pp.Print(string(pickleBuffer.Bytes()))
		_, err = zlibWriter.Write(pickleBuffer.Bytes())
		if err != nil {
			return errors.Wrap(err, "zip stream failed")
		}
		event.DetailsRefRaw = pointer.ToString(base64.StdEncoding.EncodeToString(zlibBuffer.Bytes()))
		event.DetailsRef = nil
	}
	_, err := tx.InsertInto("sentry_message").
		Columns("id", "group_id", "message_id", "project_id", "message",
			"platform", "time_spent", "data", "datetime").
		Record(event).
		Exec()
	return errors.Wrap(err, "failed to save issue event")
}

// TODO move to the nodeStore
func (s EventStore) SaveNodeBlob(tx *dbr.Tx, nodeBlob models.NodeBlob) error {
	_, err := tx.InsertInto("nodestore_node").
		Columns("id", "data", "timestamp").
		Record(nodeBlob).
		Exec()
	return errors.Wrap(err, "failed to save node blob")
}

func unpickleZippedBase64String(blob string) (map[string]interface{}, error) {
	zippedBytes, err := base64.StdEncoding.DecodeString(blob)
	if err != nil {
		return nil, errors.Wrap(err, "base64 decode failed")
	}
	zlibReader, err := zlib.NewReader(bytes.NewReader(zippedBytes))
	defer zlibReader.Close()
	if err != nil {
		return nil, errors.Wrap(err, "unzip stream failed")
	}
	unpickledBlob, err := pickle.Unpickle(zlibReader)
	if err != nil {
		return nil, errors.Wrap(err, "unpickle failed")
	}
	unpickledMap := conv.StringMap(unpickledBlob)
	return unpickledMap, nil
}
