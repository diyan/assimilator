package models

import "time"

// SavedSearch is a model for saved search query.
type SavedSearch struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   int       `db:"project_id" json:"-"`
	Name        string    `db:"name" json:"name"`
	Query       string    `db:"query" json:"query"`
	DateCreated time.Time `db:"date_added" json:"dateCreated"`
	IsDefault   bool      `db:"is_default" json:"isDefault"`
	// TODO JSON payload contains isUserDefault property
}
