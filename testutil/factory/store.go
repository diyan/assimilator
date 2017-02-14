package factory

import (
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"
)

func (tf TestFactory) SaveOrganization(org models.Organization) {
	orgStore := store.NewOrganizationStore(tf.ctx)
	tf.noError(orgStore.SaveOrganization(org))
}

func (tf TestFactory) SaveOrganizationMember(orgMember models.OrganizationMember) {
	orgStore := store.NewOrganizationStore(tf.ctx)
	tf.noError(orgStore.SaveOrganizationMember(orgMember))
}

func (tf TestFactory) SaveTeam(team models.Team) {
	teamStore := store.NewTeamStore(tf.ctx)
	tf.noError(teamStore.SaveTeam(team))
}

func (tf TestFactory) SaveTeamMember(teamMember models.OrganizationMemberTeam) {
	teamStore := store.NewTeamStore(tf.ctx)
	tf.noError(teamStore.SaveMember(teamMember))
}

func (tf TestFactory) SaveProject(project models.Project) {
	projectStore := store.NewProjectStore(tf.ctx)
	tf.noError(projectStore.SaveProject(project))
}

func (tf TestFactory) SaveEnvironment(environment models.Environment) {
	projectStore := store.NewProjectStore(tf.ctx)
	tf.noError(projectStore.SaveEnvironment(environment))
}

func (tf TestFactory) SaveTags(tags ...*models.TagKey) {
	projectStore := store.NewProjectStore(tf.ctx)
	tf.noError(projectStore.SaveTags(tags...))
}

func (tf TestFactory) SaveProjectSearches(searches ...models.SavedSearch) {
	projectStore := store.NewProjectStore(tf.ctx)
	tf.noError(projectStore.SaveSearches(searches...))
}

func (tf TestFactory) SaveUser(user models.User) {
	store := store.NewUserStore(tf.ctx)
	tf.noError(store.SaveUser(user))
}
