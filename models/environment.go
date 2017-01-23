package models

import "time"

// Environment ...
type Environment struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   int       `db:"project_id" json:"-"`
	Name        string    `db:"name" json:"name"`
	DateCreated time.Time `db:"date_added" json:"-"`
}

/*
describe sentry_environment
+------------+--------------------------+------------------------------------------------------------------+
| Column     | Type                     | Modifiers                                                        |
|------------+--------------------------+------------------------------------------------------------------|
| id         | integer                  |  not null default nextval('sentry_environment_id_seq'::regclass) |
| project_id | integer                  |  not null                                                        |
| name       | character varying(64)    |  not null                                                        |
| date_added | timestamp with time zone |  not null                                                        |
+------------+--------------------------+------------------------------------------------------------------+
*/
