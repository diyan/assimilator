package models

import "time"

// Team represents a group of individuals which maintain ownership of projects.
type Team struct {
	ID             int       `db:"id" json:"id,string"`
	Slug           string    `db:"slug" json:"slug"`
	Name           string    `db:"name" json:"name"`
	OrganizationID int       `db:"organization_id" json:"-"`
	Status         int       `db:"status" json:"-"`
	DateCreated    time.Time `db:"date_added" json:"dateCreated"`
}

const (
	TeamStatusVisible            = 0
	TeamStatusPendingDeletion    = 1
	TeamStatusDeletionInProgress = 2
)
