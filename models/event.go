package models

import "time"

// Event represents an individual event.
type Event struct {
	ID        int     `db:"id" json:"id,string"`
	GroupID   *int    `db:"group_id" json:"groupID,string"`
	EventID   *string `db:"message_id" json:"eventID"`
	ProjectID *int    `db:"project_id" json:"-"`
	Message   string  `db:"message" json:"-"`
	Platform  *string `db:"platform" json:"platform"`
	TimeSpent *int    `db:"time_spent" json:"-"`
	// NOTE data has NodeField type
	Data *string `db:"data" json:"-"`
	// TODO check that `dateCreated` is the name for JSON
	DateCreated time.Time `db:"datetime" json:"dateCreated"`
}
