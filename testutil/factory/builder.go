package factory

import (
	"time"

	"github.com/AlekSi/pointer"
	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/models"
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

func (tf TestFactory) MakeProjectSearches() []models.SavedSearch {
	return []models.SavedSearch{
		models.SavedSearch{
			ID:          1,
			ProjectID:   1,
			Name:        "Unresolved Issues",
			Query:       "is:unresolved",
			DateCreated: time_of_2999_01_01__00_00_00,
			IsDefault:   true,
		},
		models.SavedSearch{
			ID:          2,
			ProjectID:   1,
			Name:        "Needs Triage",
			Query:       "is:unresolved is:unassigned",
			DateCreated: time_of_2999_01_01__00_00_00,
			IsDefault:   false,
		},
		models.SavedSearch{
			ID:          3,
			ProjectID:   1,
			Name:        "Assigned To Me",
			Query:       "is:unresolved assigned:me",
			DateCreated: time_of_2999_01_01__00_00_00,
			IsDefault:   false,
		},
		models.SavedSearch{
			ID:          4,
			ProjectID:   1,
			Name:        "My Bookmarks",
			Query:       "is:unresolved bookmarks:me",
			DateCreated: time_of_2999_01_01__00_00_00,
			IsDefault:   false,
		},
		models.SavedSearch{
			ID:          5,
			ProjectID:   1,
			Name:        "New Today",
			Query:       "is:unresolved age:-24h",
			DateCreated: time_of_2999_01_01__00_00_00,
			IsDefault:   false,
		},
	}
}

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
		Email:           nil,
		HasGlobalAccess: true,
		Flags:           0,
		Counter:         0,
		Role:            "owner",
		Token:           nil,
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

func (tf TestFactory) MakeEnvironment() models.Environment {
	return models.Environment{
		ID:          1,
		ProjectID:   1,
		Name:        "",
		DateCreated: time_of_2999_01_01__00_00_00,
	}
}

func (tf TestFactory) MakeEventGroup() models.Group {
	return models.Group{
		ID:        1,
		ProjectID: pointer.ToInt(1),
		Logger:    "",
		Level:     20, // TODO Add enums
		Message:   "This is a test message generated using ``raven test`` __main__ in <module>",
		Culprit:   pointer.ToString("__main__ in <module>"),
		Status:    0, // TODO Add enums
		TimesSeen: 1,
		LastSeen:  time_of_2999_01_01__00_00_00,
		FirstSeen: time_of_2999_01_01__00_00_00,
		// TODO Data is most likely base64 -> gzip -> dict
		Data:           pointer.ToString("eJwdykEOgjAUBND9P0V3sDKpQO0JvADErf2xY20CpOF/SLy9xWQ2M/PaWCyNzcyizw0v5AOxoXKlu+390PXeOUsyNvotqHtXbcSb91lr689ngXJk5doHamNxlWjW+eQ3ekyfLKaGjULULBDhBJOwYmNFNLvkNZkQNj6w/lEIVDyJXH6p7jGr"),
		Score:          1485348661, // TODO what does this mean?
		TimeSpentTotal: 0,
		TimeSpentCount: 0,
		ResolvedAt:     nil,
		ActiveAt:       pointer.ToTime(time_of_2999_01_01__00_00_00),
		IsPublic:       pointer.ToBool(false),
		Platform:       pointer.ToString("python"),
		NumComments:    pointer.ToInt(0),
		FirstReleaseID: nil,
		ShortID:        pointer.ToInt(1),
	}
}

func (tf TestFactory) MakeUser() models.User {
	return models.User{
		ID:                1,
		Username:          "admin",
		Name:              "",
		Email:             "admin@example.com",
		IsStaff:           true,
		IsActive:          true,
		IsSuperuser:       true,
		IsManaged:         false,
		IsPasswordExpired: false,
		// TODO explain what is a plain-text equivalent
		Password:           "pbkdf2_sha256$12000$GrqCKrh4gpuI$PLLnjVsHTgSDCcAv6ql0rJ3Z/5RE9oNoaHHc8D/WTtE=",
		DateCreated:        time_of_2999_01_01__00_00_00,
		LastLogin:          time_of_2999_01_01__00_00_00,
		LastPasswordChange: time_of_2999_01_01__00_00_00,
	}
}
