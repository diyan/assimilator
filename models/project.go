package models

import "time"

// Project is a permission based namespace which generally
// is the top level entry point for all data.
type Project struct {
	ID             int       `db:"id" json:"id,string"`
	TeamID         int       `db:"team_id" json:"-"`
	OrganizationID int       `db:"organization_id" json:"-"`
	Name           string    `db:"name" json:"name"`
	Slug           string    `db:"slug" json:"slug"`
	Public         bool      `db:"public" json:"public"`
	Status         int       `db:"status" json:"-"`
	FirstEvent     time.Time `db:"first_event" json:"firstEvent"`
	DateCreated    time.Time `db:"date_added" json:"dateCreated"`
}

const (
	ProjectStatusVisible            = 0
	ProjectStatusHidden             = 1
	ProjectStatusPendingDeletion    = 2
	ProjectStatusDeletionInProgress = 3
)

/*
class Project(Model):
    """
    Projects are permission based namespaces which generally
    are the top level entry point for all data.
    """
    __core__ = True

    slug = models.SlugField(null=True)
    name = models.CharField(max_length=200)
    forced_color = models.CharField(max_length=6, null=True, blank=True)
    organization = FlexibleForeignKey('sentry.Organization')
    team = FlexibleForeignKey('sentry.Team')
    public = models.BooleanField(default=False)
    date_added = models.DateTimeField(default=timezone.now)
    status = BoundedPositiveIntegerField(default=0, choices=(
        (ProjectStatus.VISIBLE, _('Active')),
        (ProjectStatus.PENDING_DELETION, _('Pending Deletion')),
        (ProjectStatus.DELETION_IN_PROGRESS, _('Deletion in Progress')),
    ), db_index=True)
    # projects that were created before this field was present
    # will have their first_event field set to date_added
    first_event = models.DateTimeField(null=True)

    objects = ProjectManager(cache_fields=[
        'pk',
        'slug',
    ])

    class Meta:
        app_label = 'sentry'
        db_table = 'sentry_project'
        unique_together = (('team', 'slug'), ('organization', 'slug'))

    __repr__ = sane_repr('team_id', 'name', 'slug')

    def __unicode__(self):
        return u'%s (%s)' % (self.name, self.slug)

    def next_short_id(self):
        from sentry.models import Counter
        return Counter.increment(self)

    def save(self, *args, **kwargs):
        if not self.slug:
            lock = locks.get('slug:project', duration=5)
            with TimedRetryPolicy(10)(lock.acquire):
                slugify_instance(self, self.name, organization=self.organization)
            super(Project, self).save(*args, **kwargs)
        else:
            super(Project, self).save(*args, **kwargs)

    def get_absolute_url(self):
        return absolute_uri('/{}/{}/'.format(self.organization.slug, self.slug))

    def merge_to(self, project):
        from sentry.models import (
            Group, GroupTagValue, Event, TagValue
        )

        if not isinstance(project, Project):
            project = Project.objects.get_from_cache(pk=project)

        for group in Group.objects.filter(project=self):
            try:
                other = Group.objects.get(
                    project=project,
                )
            except Group.DoesNotExist:
                group.update(project=project)
                GroupTagValue.objects.filter(
                    project=self,
                    group_id=group,
                ).update(project=project)
            else:
                Event.objects.filter(
                    group_id=group.id,
                ).update(group_id=other.id)

                for obj in GroupTagValue.objects.filter(group=group):
                    obj2, created = GroupTagValue.objects.get_or_create(
                        project=project,
                        group=group,
                        key=obj.key,
                        value=obj.value,
                        defaults={'times_seen': obj.times_seen}
                    )
                    if not created:
                        obj2.update(times_seen=F('times_seen') + obj.times_seen)

        for fv in TagValue.objects.filter(project=self):
            TagValue.objects.get_or_create(project=project, key=fv.key, value=fv.value)
            fv.delete()
        self.delete()

    def is_internal_project(self):
        for value in (settings.SENTRY_FRONTEND_PROJECT, settings.SENTRY_PROJECT):
            if six.text_type(self.id) == six.text_type(value) or six.text_type(self.slug) == six.text_type(value):
                return True
        return False

    def get_tags(self, with_internal=True):
        from sentry.models import TagKey

        if not hasattr(self, '_tag_cache'):
            tags = self.get_option('tags', None)
            if tags is None:
                tags = [
                    t for t in TagKey.objects.all_keys(self)
                    if with_internal or not t.startswith('sentry:')
                ]
            self._tag_cache = tags
        return self._tag_cache

    # TODO: Make these a mixin
    def update_option(self, *args, **kwargs):
        from sentry.models import ProjectOption

        return ProjectOption.objects.set_value(self, *args, **kwargs)

    def get_option(self, *args, **kwargs):
        from sentry.models import ProjectOption

        return ProjectOption.objects.get_value(self, *args, **kwargs)

    def delete_option(self, *args, **kwargs):
        from sentry.models import ProjectOption

        return ProjectOption.objects.unset_value(self, *args, **kwargs)

    @property
    def callsign(self):
        return self.slug.upper()

    @property
    def color(self):
        if self.forced_color is not None:
            return '#%s' % self.forced_color
        return get_hashed_color(self.callsign or self.slug)

    @property
    def member_set(self):
        from sentry.models import OrganizationMember
        return self.organization.member_set.filter(
            id__in=OrganizationMember.objects.filter(
                organizationmemberteam__is_active=True,
                organizationmemberteam__team=self.team,
            ).values('id'),
            user__is_active=True,
        ).distinct()

    def has_access(self, user, access=None):
        from sentry.models import AuthIdentity, OrganizationMember

        warnings.warn('Project.has_access is deprecated.', DeprecationWarning)

        queryset = self.member_set.filter(user=user)

        if access is not None:
            queryset = queryset.filter(type__lte=access)

        try:
            member = queryset.get()
        except OrganizationMember.DoesNotExist:
            return False

        try:
            auth_identity = AuthIdentity.objects.get(
                auth_provider__organization=self.organization_id,
                user=member.user_id,
            )
        except AuthIdentity.DoesNotExist:
            return True

        return auth_identity.is_valid(member)

    def get_audit_log_data(self):
        return {
            'id': self.id,
            'slug': self.slug,
            'name': self.name,
            'status': self.status,
            'public': self.public,
        }

    def get_full_name(self):
        if self.team.name not in self.name:
            return '%s %s' % (self.team.name, self.name)
        return self.name

    def is_user_subscribed_to_mail_alerts(self, user):
        from sentry.models import UserOption
        is_enabled = UserOption.objects.get_value(
            user, self, 'mail:alert', None)
        if is_enabled is None:
            is_enabled = UserOption.objects.get_value(
                user, None, 'subscribe_by_default', '1') == '1'
        else:
            is_enabled = bool(is_enabled)
        return is_enabled

    def is_user_subscribed_to_workflow(self, user):
        from sentry.models import UserOption, UserOptionValue

        opt_value = UserOption.objects.get_value(
            user, self, 'workflow:notifications', None)
        if opt_value is None:
            opt_value = UserOption.objects.get_value(
                user, None, 'workflow:notifications',
                UserOptionValue.all_conversations)
        return opt_value == UserOptionValue.all_conversations
*/
