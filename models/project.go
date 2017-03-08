package models

import "time"

// Project is a permission based namespace which generally
// is the top level entry point for all data.
type Project struct {
	ID             int       `db:"id" json:"id,string"`
	TeamID         int       `db:"team_id" json:"-"`
	OrganizationID int       `db:"organization_id" json:"-"`
	Name           string    `db:"name" json:"name"`
	Slug           string    `db:"slug" json:"slug"`
	Public         bool      `db:"public" json:"isPublic"`
	Status         int       `db:"status" json:"-"`
	FirstEvent     time.Time `db:"first_event" json:"firstEvent"`
	DateCreated    time.Time `db:"date_added" json:"dateCreated"`
}

const (
	ProjectStatusVisible            = 0
	ProjectStatusHidden             = 1
	ProjectStatusPendingDeletion    = 2
	ProjectStatusDeletionInProgress = 3
)
