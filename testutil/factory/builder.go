package factory

import (
	"time"

	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
)

var time_of_2999_01_01__00_00_00 = time.Date(2999, time.January, 1, 0, 0, 0, 0, time.UTC)

// TODO Define naming for fixture-per-test and fixture-per-suite
func MakeAppConfig() conf.Config {
	return conf.Config{
		Port:            3000,
		DatabaseURL:     "postgres://postgres@localhost/sentry_ci?sslmode=disable",
		InitialTeam:     "ACME-Team",
		InitialProject:  "ACME",
		InitialKey:      "763a78a695424ed687cf8b7dc26d3161:763a78a695424ed687cf8b7dc26d3161",
		InitialPlatform: "python",
	}
}

func (tf TestFactory) MakeTags() []*models.TagKey {
	tag1 := models.TagKey{
		ID:        1,
		ProjectID: 1,
		Key:       "server_name",
	}
	tag2 := tag1
	tag2.ID = 2
	tag2.Key = "level"
	return []*models.TagKey{&tag1, &tag2}
}

func (tf TestFactory) MakeOrganization() models.Organization {
	return models.Organization{
		ID:          1,
		Name:        "ACME-Team",
		Slug:        "acme-team",
		Status:      models.OrganizationStatusVisible,
		Flags:       1, // TODO Introduce constants
		DefaultRole: "member",
		DateCreated: time_of_2999_01_01__00_00_00,
	}
}

func (tf TestFactory) MakeOrganizationMember() models.OrganizationMember {
	return models.OrganizationMember{
		ID:              1,
		OrganizationID:  1,
		UserID:          1,
		Type:            50, // TODO introduce constant
		DateCreated:     time_of_2999_01_01__00_00_00,
		Email:           dbr.NullString{},
		HasGlobalAccess: true,
		Flags:           0,
		Role:            "owner",
		Token:           dbr.NullString{},
	}
}

func (tf TestFactory) MakeTeam() models.Team {
	return models.Team{
		ID:             1,
		Slug:           "acme-team",
		Name:           "ACME-Team",
		OrganizationID: 1,
		Status:         models.TeamStatusVisible,
		DateCreated:    time_of_2999_01_01__00_00_00,
	}
}

func (tf TestFactory) MakeTeamMember() models.OrganizationMemberTeam {
	return models.OrganizationMemberTeam{
		ID:                   1,
		OrganizationMemberID: 1,
		TeamID:               1,
		IsActive:             true,
	}
}

func (tf TestFactory) MakeProject() models.Project {
	return models.Project{
		ID:             1,
		TeamID:         1,
		OrganizationID: 1,
		Name:           "ACME",
		Slug:           "acme",
		Public:         false,
		Status:         models.ProjectStatusVisible,
		FirstEvent:     time_of_2999_01_01__00_00_00,
		DateCreated:    time_of_2999_01_01__00_00_00,
	}
}
