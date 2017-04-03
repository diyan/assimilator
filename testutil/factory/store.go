package factory

import (
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"
)

func (tf TestFactory) SaveOrganization(org models.Organization) {
	store := store.NewOrganizationStore()
	tf.noError(store.SaveOrganization(tf.tx, org))
}

func (tf TestFactory) SaveOrganizationMember(orgMember models.OrganizationMember) {
	store := store.NewOrganizationStore()
	tf.noError(store.SaveOrganizationMember(tf.tx, orgMember))
}

func (tf TestFactory) SaveTeam(team models.Team) {
	store := store.NewTeamStore()
	tf.noError(store.SaveTeam(tf.tx, team))
}

func (tf TestFactory) SaveTeamMember(teamMember models.OrganizationMemberTeam) {
	store := store.NewTeamStore()
	tf.noError(store.SaveMember(tf.tx, teamMember))
}

func (tf TestFactory) SaveProject(project models.Project) {
	store := store.NewProjectStore()
	tf.noError(store.SaveProject(tf.tx, project))
}

func (tf TestFactory) SaveEnvironment(environment models.Environment) {
	store := store.NewProjectStore()
	tf.noError(store.SaveEnvironment(tf.tx, environment))
}

func (tf TestFactory) SaveTags(tags ...*models.TagKey) {
	store := store.NewProjectStore()
	tf.noError(store.SaveTags(tf.tx, tags...))
}

func (tf TestFactory) SaveProjectSearches(searches ...models.SavedSearch) {
	store := store.NewProjectStore()
	tf.noError(store.SaveSearches(tf.tx, searches...))
}

func (tf TestFactory) SaveEventGroup(group models.Group) {
	store := store.NewProjectStore()
	tf.noError(store.SaveEventGroup(tf.tx, group))
}

func (tf TestFactory) SaveEvent(event models.Event) {
	store := store.NewEventStore()
	tf.noError(store.SaveEvent(tf.tx, event))
}

func (tf TestFactory) SaveEventNodeBlob(nodeBlob models.NodeBlob) {
	store := store.NewEventStore()
	tf.noError(store.SaveNodeBlob(tf.tx, nodeBlob))
}

func (tf TestFactory) SaveUser(user models.User) {
	store := store.NewUserStore()
	tf.noError(store.SaveUser(tf.tx, user))
}
