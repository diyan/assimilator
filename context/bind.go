package context

import (
	"errors"

	"github.com/diyan/assimilator/conf"
	"github.com/diyan/assimilator/db"
	"github.com/diyan/assimilator/db/store"
	"github.com/diyan/assimilator/models"
	"github.com/labstack/echo"
)

type Binder struct {
	config    conf.Config
	dbTxMaker db.TxMakerFunc
}

func NewBinder(config conf.Config, dbTxMaker db.TxMakerFunc) Binder {
	return Binder{
		config:    config,
		dbTxMaker: dbTxMaker,
	}
}

func (b *Binder) Base(handler func(Base) error) func(echo.Context) error {
	return func(c echo.Context) error {
		baseContext, err := b.getBaseContext(c)
		if err != nil {
			return err
		}
		return handler(*baseContext)
	}
}

func (b *Binder) Organization(handler func(Organization) error) func(echo.Context) error {
	return func(c echo.Context) error {
		orgContext, err := b.getOrganizationContext(c)
		if err != nil {
			return err
		}
		return handler(*orgContext)
	}
}

func (b *Binder) Project(handler func(Project) error) func(echo.Context) error {
	return func(c echo.Context) error {
		projectContext, err := b.getProjectContext(c)
		if err != nil {
			return err
		}
		return handler(*projectContext)
	}
}

func (b *Binder) getBaseContext(c echo.Context) (*Base, error) {
	tx, err := b.dbTxMaker()
	if err != nil {
		return nil, err
	}
	// TODO remove hardcode
	user := models.User{ID: 1, Name: "admin"}
	return &Base{Context: c, Tx: tx, User: user}, nil
}

func (b *Binder) getOrganizationContext(c echo.Context) (*Organization, error) {
	orgSlug := c.Param("organization_slug")
	// TODO check with regex pattern, validate if that orgSlug exists in db
	if orgSlug == "" {
		return nil, errors.New("'organization_slug' was not provided")
	}
	baseContext, err := b.getBaseContext(c)
	if err != nil {
		return nil, err
	}
	orgStore := store.NewOrganizationStore()
	/* TODO check Sentry source. We may need to check permissions
	   err := tx.SelectBySql(`
	   		select
	   			o.*
	   		from sentry_organization o
	   			join sentry_organizationmember om on o.id = om.organization_id
	   		where o.slug = ? and om.user_id = ? and o.status = ?`,
	   		orgSlug, userID, models.OrganizationStatusVisible).
	   		LoadStruct(&org)
	   	if err != nil {
	   		return errors.Wrap(err, "can not read organization")
	   	}
	*/
	org, err := orgStore.GetOrganization(baseContext.Tx, orgSlug)
	if err != nil {
		return nil, err
	}
	return &Organization{Base: *baseContext, Organization: *org}, nil
}

func (b *Binder) getProjectContext(c echo.Context) (*Project, error) {
	projectSlug := c.Param("project_slug")
	if projectSlug == "" {
		return nil, errors.New("'project_slug' was not provided")
	}
	orgContext, err := b.getOrganizationContext(c)
	if err != nil {
		return nil, err
	}
	// TODO return ResourceDoesNotExist if record was not found
	// TODO check project permissions -> self.check_object_permissions(request, project)
	projectStore := store.NewProjectStore()
	project, err := projectStore.GetProject(
		orgContext.Tx, orgContext.Organization.Slug, projectSlug)
	if err != nil {
		return nil, err
	}
	return &Project{
		Base:         orgContext.Base,
		Organization: orgContext.Organization,
		Project:      project,
	}, nil
}
