package models

import "strings"

// TagKey stores references to available filters keys.
type TagKey struct {
	ID           int     `db:"id" json:"id,string"`
	ProjectID    int     `db:"project_id" json:"-"`
	Key          string  `db:"key" json:"key"`
	UniqueValues int     `db:"values_seen" json:"uniqueValues"`
	Name         *string `db:"label" json:"name"`
	Status       int     `db:"status" json:"-"`
}

func (tag *TagKey) PostGet() {
	// TODO introduce public const
	tagLabels := map[string]string{
		"exc_type":        "Exception Type",
		"sentry:user":     "User",
		"sentry:filename": "File",
		"sentry:function": "Function",
		"sentry:release":  "Release",
		"os":              "OS",
		"url":             "URL",
		"server_name":     "Server",
	}

	if tag.Name != nil {
		return
	}
	if label, ok := tagLabels[tag.Key]; ok {
		tag.Name = &label
	} else {
		label = strings.Title(strings.Replace(tag.Key, "_", " ", -1))
		tag.Name = &label
	}
}

const (
	TagKeyStatusVisible            = 0
	TagKeyStatusPendingDeletion    = 1
	TagKeyStatusDeletionInProgress = 2
)
