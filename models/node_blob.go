package models

import "time"

type NodeBlob struct {
	ID          string    `db:"id"`
	Data        string    `db:"data"`
	DateCreated time.Time `db:"timestamp"`
}
