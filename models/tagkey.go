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

/*
describe sentry_filterkey
+-------------+-----------------------+----------------------------------------------------------------+
| Column      | Type                  | Modifiers                                                      |
|-------------+-----------------------+----------------------------------------------------------------|
| id          | integer               |  not null default nextval('sentry_filterkey_id_seq'::regclass) |
| project_id  | integer               |  not null                                                      |
| key         | character varying(32) |  not null                                                      |
| values_seen | integer               |  not null                                                      |
| label       | character varying(64) |                                                                |
| status      | integer               |  not null                                                      |
+-------------+-----------------------+----------------------------------------------------------------+


class TagKey(Model):
    """
    Stores references to available filters keys.
    """
    __core__ = False

    project = FlexibleForeignKey('sentry.Project')
    key = models.CharField(max_length=MAX_TAG_KEY_LENGTH)
    values_seen = BoundedPositiveIntegerField(default=0)
    label = models.CharField(max_length=64, null=True)
    status = BoundedPositiveIntegerField(choices=(
        (TagKeyStatus.VISIBLE, _('Visible')),
        (TagKeyStatus.PENDING_DELETION, _('Pending Deletion')),
        (TagKeyStatus.DELETION_IN_PROGRESS, _('Deletion in Progress')),
    ), default=TagKeyStatus.VISIBLE)

    objects = TagKeyManager()

    class Meta:
        app_label = 'sentry'
        db_table = 'sentry_filterkey'
        unique_together = (('project', 'key'),)

    __repr__ = sane_repr('project_id', 'key')

    @classmethod
    def is_valid_key(cls, key):
        return bool(TAG_KEY_RE.match(key))

    @classmethod
    def is_reserved_key(cls, key):
        return key in INTERNAL_TAG_KEYS

    @classmethod
    def get_standardized_key(cls, key):
        if key.startswith('sentry:'):
            return key.split('sentry:', 1)[-1]
        return key

    def get_label(self):
        return self.label \
            or TAG_LABELS.get(self.key) \
            or self.key.replace('_', ' ').title()

    def get_audit_log_data(self):
        return {
            'key': self.key,
        }

*/
