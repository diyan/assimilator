package models

import (
	"regexp"
	"strings"
	"time"
)

var shortIDRe *regexp.Regexp

const (
	GroupStatusUnresolved         = 0
	GroupStatusResolved           = 1
	GroupStatusIgnored            = 2
	GroupStatusPendingDeletion    = 3
	GroupStatusDeletionInProgress = 4
	GroupStatusPendingMerge       = 5
	// GroupStatusMuted status will be removed in Sentry 9.0
	GroupStatusMuted = GroupStatusIgnored
)

// Group is an aggregated message which summarizes a set of Events.
type Group struct {
	ID        int       `db:"id" json:"id,string"`
	Logger    string    `db:"logger" json:"-"`
	Level     int       `db:"level" json:"-"`
	Message   string    `db:"message" json:"title"`
	Culprit   *string   `db:"view" json:"culprit"`
	Status    int       `db:"status" json:"-"`
	TimesSeen int       `db:"times_seen" json:"-"`
	LastSeen  time.Time `db:"last_seen" json:"lastSeen"`
	FirstSeen time.Time `db:"first_seen" json:"firstSeen"`
	// Data is a GzippedDictField
	Data           *string    `db:"data" json:"-"`
	Score          int        `db:"score" json:"-"`
	ProjectID      *int       `db:"project_id" json:"-"`
	TimeSpentTotal int        `db:"time_spent_total" json:"-"`
	TimeSpentCount int        `db:"time_spent_count" json:"-"`
	ResolvedAt     *time.Time `db:"resolved_at" json:"-"`
	ActiveAt       *time.Time `db:"active_at" json:"-"`
	IsPublic       *bool      `db:"is_public" json:"isPublic"`
	Platform       *string    `db:"platform" json:"-"`
	NumComments    *int       `db:"num_comments" json:"numComments"`
	FirstReleaseID *int       `db:"first_release_id" json:"-"`
	ShortID        *int       `db:"short_id" json:"-"`
}

func init() {
	shortIDRe = regexp.MustCompile(`^(.*?)(?:[\s_-])([A-Za-z0-9-._]+)$`)
}

func LooksLikeShortID(value string) bool {
	return shortIDRe.MatchString(strings.TrimSpace(value))
}
