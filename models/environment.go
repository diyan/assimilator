package models

import "time"

// Environment ...
type Environment struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   int       `db:"project_id" json:"-"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_added" json:"-"`
}
