package models

import "time"

// SavedSearch is a model for saved search query.
type SavedSearch struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   string    `db:"project_id" json:"-"`
	Name        string    `db:"name" json:"name"`
	Query       string    `db:"query" json:"query"`
	DateCreated time.Time `db:"date_added" json:"dateCreated"`
	IsDefault   bool      `db:"is_default" json:"isDefault"`
	// TODO JSON payload contains isUserDefault property
}

/*
describe sentry_savedsearch
+------------+--------------------------+------------------------------------------------------------------+
| Column     | Type                     | Modifiers                                                        |
|------------+--------------------------+------------------------------------------------------------------|
| id         | integer                  |  not null default nextval('sentry_savedsearch_id_seq'::regclass) |
| project_id | integer                  |  not null                                                        |
| name       | character varying(128)   |  not null                                                        |
| query      | text                     |  not null                                                        |
| date_added | timestamp with time zone |  not null                                                        |
| is_default | boolean                  |  not null                                                        |
+------------+--------------------------+------------------------------------------------------------------+
*/
