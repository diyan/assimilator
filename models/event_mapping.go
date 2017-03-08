package models

import "time"

// EventMapping ...
type EventMapping struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   int       `db:"project_id" json:"-"`
	GroupID     int       `db:"group_id" json:"-"`
	EventID     string    `db:"event_id" json:"-"`
	DateCreated time.Time `db:"date_added" json:"-"`
}

// TODO Consider add lazy load of Team, Project, Group relations
