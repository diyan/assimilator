package models

import "time"

// EventMapping ...
type EventMapping struct {
	ID          int       `db:"id" json:"id,string"`
	ProjectID   int       `db:"project_id" json:"-"`
	GroupID     int       `db:"group_id" json:"-"`
	EventID     string    `db:"event_id" json:"-"`
	DateCreated time.Time `db:"date_added" json:"-"`
}

// TODO Consider add lazy load of Team, Project, Group relations

/*
sentry> describe sentry_eventmapping
+------------+--------------------------+-------------------------------------------------------------------+
| Column     | Type                     | Modifiers                                                         |
|------------+--------------------------+-------------------------------------------------------------------|
| id         | integer                  |  not null default nextval('sentry_eventmapping_id_seq'::regclass) |
| project_id | integer                  |  not null                                                         |
| group_id   | integer                  |  not null                                                         |
| event_id   | character varying(32)    |  not null                                                         |
| date_added | timestamp with time zone |  not null                                                         |
+------------+--------------------------+-------------------------------------------------------------------+


  # Implement a ForeignKey-like accessor for backwards compat
  def _set_group(self, group):
      self.group_id = group.id
      self._group_cache = group

  def _get_group(self):
      from sentry.models import Group
      if not hasattr(self, '_group_cache'):
          self._group_cache = Group.objects.get(id=self.group_id)
      return self._group_cache

  group = property(_get_group, _set_group)

  # Implement a ForeignKey-like accessor for backwards compat
  def _set_project(self, project):
      self.project_id = project.id
      self._project_cache = project

  def _get_project(self):
      from sentry.models import Project
      if not hasattr(self, '_project_cache'):
          self._project_cache = Project.objects.get(id=self.project_id)
      return self._project_cache

  project = property(_get_project, _set_project)
*/
