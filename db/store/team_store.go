package store

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/pkg/errors"
)

type TeamStore struct {
}

func NewTeamStore() TeamStore {
	return TeamStore{}
}

func (s TeamStore) SaveTeam(tx *dbr.Tx, team models.Team) error {
	_, err := tx.InsertInto("sentry_team").
		Columns("id", "slug", "name", "date_added", "status", "organization_id").
		Record(team).
		Exec()
	return errors.Wrap(err, "failed to save team")
}

func (s TeamStore) SaveMember(tx *dbr.Tx, teamMember models.OrganizationMemberTeam) error {
	_, err := tx.InsertInto("sentry_organizationmember_teams").
		Columns("id", "organizationmember_id", "team_id", "is_active").
		Record(teamMember).
		Exec()
	return errors.Wrap(err, "failed to save organization team member")
}
