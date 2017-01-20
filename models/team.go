package models

import "time"

// Team represents a group of individuals which maintain ownership of projects.
type Team struct {
	ID             int       `db:"id" json:"id,string"`
	Slug           string    `db:"slug" json:"slug"`
	Name           string    `db:"name" json:"name"`
	DateCreated    time.Time `db:"date_added" json:"dateCreated"`
	Status         int       `db:"status" json:"status"`
	OrganizationID int       `db:"organization_id" json:"-"`
}

const (
	TeamStatusVisible            = 0
	TeamStatusPendingDeletion    = 1
	TeamStatusDeletionInProgress = 2
)

/*
class Team(Model):
    """
    A team represents a group of individuals which maintain ownership of projects.
    """
    __core__ = True

    organization = FlexibleForeignKey('sentry.Organization')
    slug = models.SlugField()
    name = models.CharField(max_length=64)
    status = BoundedPositiveIntegerField(choices=(
        (TeamStatus.VISIBLE, _('Active')),
        (TeamStatus.PENDING_DELETION, _('Pending Deletion')),
        (TeamStatus.DELETION_IN_PROGRESS, _('Deletion in Progress')),
    ), default=TeamStatus.VISIBLE)
    date_added = models.DateTimeField(default=timezone.now, null=True)

    objects = TeamManager(cache_fields=(
        'pk',
        'slug',
    ))

    class Meta:
        app_label = 'sentry'
        db_table = 'sentry_team'
        unique_together = (('organization', 'slug'),)

    __repr__ = sane_repr('name', 'slug')

    def __unicode__(self):
        return u'%s (%s)' % (self.name, self.slug)

    def save(self, *args, **kwargs):
        if not self.slug:
            lock = locks.get('slug:team', duration=5)
            with TimedRetryPolicy(10)(lock.acquire):
                slugify_instance(self, self.name, organization=self.organization)
            super(Team, self).save(*args, **kwargs)
        else:
            super(Team, self).save(*args, **kwargs)

    @property
    def member_set(self):
        return self.organization.member_set.filter(
            organizationmemberteam__team=self,
            organizationmemberteam__is_active=True,
            user__is_active=True,
        ).distinct()

    def has_access(self, user, access=None):
        from sentry.models import AuthIdentity, OrganizationMember

        warnings.warn('Team.has_access is deprecated.', DeprecationWarning)

        queryset = self.member_set.filter(
            user=user,
        )
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

    def transfer_to(self, organization):
        """
        Transfers a team and all projects under it to the given organization.
        """
        from sentry.models import (
            OrganizationAccessRequest, OrganizationMember,
            OrganizationMemberTeam, Project
        )

        try:
            with transaction.atomic():
                self.update(organization=organization)
        except IntegrityError:
            # likely this means a team already exists, let's try to coerce to
            # it instead of a blind transfer
            new_team = Team.objects.get(
                organization=organization,
                slug=self.slug,
            )
        else:
            new_team = self

        Project.objects.filter(
            team=self,
        ).exclude(
            organization=organization,
        ).update(
            team=new_team,
            organization=organization,
        )

        # remove any pending access requests from the old organization
        if self != new_team:
            OrganizationAccessRequest.objects.filter(
                team=self,
            ).delete()

        # identify shared members and ensure they retain team access
        # under the new organization
        old_memberships = OrganizationMember.objects.filter(
            teams=self,
        ).exclude(
            organization=organization,
        )
        for member in old_memberships:
            try:
                new_member = OrganizationMember.objects.get(
                    user=member.user,
                    organization=organization,
                )
            except OrganizationMember.DoesNotExist:
                continue

            try:
                with transaction.atomic():
                    OrganizationMemberTeam.objects.create(
                        team=new_team,
                        organizationmember=new_member,
                    )
            except IntegrityError:
                pass

        OrganizationMemberTeam.objects.filter(
            team=self,
        ).exclude(
            organizationmember__organization=organization,
        ).delete()

        if new_team != self:
            cursor = connections[router.db_for_write(Team)].cursor()
            # we use a cursor here to avoid automatic cascading of relations
            # in Django
            try:
                cursor.execute('DELETE FROM sentry_team WHERE id = %s', [self.id])
            finally:
                cursor.close()

    def get_audit_log_data(self):
        return {
            'id': self.id,
            'slug': self.slug,
            'name': self.name,
            'status': self.status,
        }
*/
