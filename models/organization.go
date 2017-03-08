package models

import "time"

// Organization represents a group of individuals which maintain ownership of projects.
type Organization struct {
	ID          int       `db:"id" json:"id,string"`
	Name        string    `db:"name" json:"name"`
	Slug        string    `db:"slug" json:"slug"`
	Status      int       `db:"status" json:"-"`
	Flags       int       `db:"flags" json:"-"`
	DefaultRole string    `db:"default_role" json:"-"`
	DateCreated time.Time `db:"date_added" json:"dateCreated"`
}

const (
	OrganizationStatusVisible           = 0
	OrganizationStatusPendingDeletion   = 1
	OrganiztionStatusDeletionInProgress = 2
)
