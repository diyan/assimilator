package store

import (
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type TeamStore struct {
	c echo.Context
}

func NewTeamStore(c echo.Context) TeamStore {
	return TeamStore{c: c}
}

func (s TeamStore) SaveTeam(team models.Team) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save team")
	}
	_, err = db.InsertInto("sentry_team").
		Columns("id", "slug", "name", "date_added", "status", "organization_id").
		Record(team).
		Exec()
	return errors.Wrap(err, "failed to save team")
}

func (s TeamStore) SaveMember(teamMember models.OrganizationMemberTeam) error {
	db, err := db.FromE(s.c)
	if err != nil {
		return errors.Wrap(err, "failed to save organization team member")
	}
	_, err = db.InsertInto("sentry_organizationmember_teams").
		Columns("id", "organizationmember_id", "team_id", "is_active").
		Record(teamMember).
		Exec()
	return errors.Wrap(err, "failed to save organization team member")
}
