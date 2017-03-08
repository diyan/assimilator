package models

import "time"

type OrganizationMemberTeam struct {
	ID                   int  `db:"id" json:"id,string"`
	OrganizationMemberID int  `db:"organizationmember_id" json:"organizationmemberId"`
	TeamID               int  `db:"team_id" json:"teamId"`
	IsActive             bool `db:"is_active" json:"isActive"`
}

// OrganizationMember identifies relationships between teams and users.
// Users listed as team members are considered to have access to all projects
// and could be thought of as team owners (though their access level may not)
// be set to ownership.
type OrganizationMember struct {
	ID              int       `db:"id" json:"id,string"`
	OrganizationID  int       `db:"organization_id" json:"organizationId"`
	UserID          int       `db:"user_id" json:"userId"`
	Type            int       `db:"type" json:"type"`
	DateCreated     time.Time `db:"date_added" json:"dateCreated"`
	Email           *string   `db:"email" json:"email"`
	HasGlobalAccess bool      `db:"has_global_access" json:"hasGlobalAccess"`
	Flags           int64     `db:"flags" json:"flags"`
	Counter         int       `db:"counter" json:"counter"`
	Role            string    `db:"role" json:"role"`
	Token           *string   `db:"token" json:"token"`
}
