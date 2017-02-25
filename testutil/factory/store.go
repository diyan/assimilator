package factory

import (
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"
)

func (tf TestFactory) SaveOrganization(org models.Organization) {
	store := store.NewOrganizationStore(tf.ctx)
	tf.noError(store.SaveOrganization(org))
}

func (tf TestFactory) SaveOrganizationMember(orgMember models.OrganizationMember) {
	store := store.NewOrganizationStore(tf.ctx)
	tf.noError(store.SaveOrganizationMember(orgMember))
}

func (tf TestFactory) SaveTeam(team models.Team) {
	store := store.NewTeamStore(tf.ctx)
	tf.noError(store.SaveTeam(team))
}

func (tf TestFactory) SaveTeamMember(teamMember models.OrganizationMemberTeam) {
	store := store.NewTeamStore(tf.ctx)
	tf.noError(store.SaveMember(teamMember))
}

func (tf TestFactory) SaveProject(project models.Project) {
	store := store.NewProjectStore(tf.ctx)
	tf.noError(store.SaveProject(project))
}

func (tf TestFactory) SaveEnvironment(environment models.Environment) {
	store := store.NewProjectStore(tf.ctx)
	tf.noError(store.SaveEnvironment(environment))
}

func (tf TestFactory) SaveTags(tags ...*models.TagKey) {
	store := store.NewProjectStore(tf.ctx)
	tf.noError(store.SaveTags(tags...))
}

func (tf TestFactory) SaveProjectSearches(searches ...models.SavedSearch) {
	store := store.NewProjectStore(tf.ctx)
	tf.noError(store.SaveSearches(searches...))
}

func (tf TestFactory) SaveEventGroup(group models.Group) {
	store := store.NewProjectStore(tf.ctx)
	tf.noError(store.SaveEventGroup(group))
}

func (tf TestFactory) SaveEvent(event models.Event) {
	store := store.NewEventStore(tf.ctx)
	tf.noError(store.SaveEvent(event))
}

func (tf TestFactory) SaveEventNodeBlob(nodeBlob models.NodeBlob) {
	store := store.NewEventStore(tf.ctx)
	tf.noError(store.SaveNodeBlob(nodeBlob))
}

func (tf TestFactory) SaveUser(user models.User) {
	store := store.NewUserStore(tf.ctx)
	tf.noError(store.SaveUser(user))
}
