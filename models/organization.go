package models

import "time"

// Organization represents a group of individuals which maintain ownership of projects.
type Organization struct {
	ID          int       `db:"id" json:"id,string"`
	Name        string    `db:"name" json:"name"`
	Slug        string    `db:"slug" json:"slug"`
	DateCreated time.Time `db:"date_added" json:"dateCreated"`
}

const (
	OrganizationStatusVisible           = 0
	OrganizationStatusPendingDeletion   = 1
	OrganiztionStatusDeletionInProgress = 2
)

/*CREATE TABLE sentry_organization (
    id integer NOT NULL,
    name character varying(64) NOT NULL,
    status integer NOT NULL,
    date_added timestamp with time zone NOT NULL,
    slug character varying(50) NOT NULL,
    flags bigint NOT NULL,
    default_role character varying(32) NOT NULL,
    CONSTRAINT sentry_organization_status_check CHECK ((status >= 0))
);
*/
