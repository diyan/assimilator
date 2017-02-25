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

func (tf TestFactory) MakeEvent() models.Event {
	return models.Event{
		ID:          1,
		GroupID:     pointer.ToInt(1),
		ProjectID:   pointer.ToInt(1),
		EventID:     pointer.ToString("dcf8c1d1cd284d3fbfeffb43ddb7c0f7"),
		Message:     "This is a test message generated using ``raven test`` __main__ in <module>",
		DateCreated: time_of_2999_01_01__00_00_00,
		TimeSpent:   nil,
		Platform:    pointer.ToString("python"),
		Data:        pointer.ToString("eJzTSCkw5ApWz8tPSY3PTFHnKjAC8vwMIkLyI3Iygsoqnf3TIjMMi4MKohxtbYHSxlzFegCg1g+U"),
	}
}

func (tf TestFactory) MakeEventNodeBlob() models.NodeBlob {
	return models.NodeBlob{
		ID:          "N0XToXlhRvyCOfYh1sRpZA==",
		Data:        `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`,
		DateCreated: time_of_2999_01_01__00_00_00,
	}
}

func (tf TestFactory) MakeEventNodeBlobV2() models.NodeBlob {
	return models.NodeBlob{
		ID:          "N0XToXlhRvyCOfYh1sRpZA==",
		Data:        `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`,
		DateCreated: time_of_2999_01_01__00_00_00,
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
