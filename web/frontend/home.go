package frontend

import (
	"net/http"

	"github.com/diyan/assimilator/context"
	"github.com/diyan/assimilator/models"
	"github.com/pkg/errors"
)

// Organization represents a group of individuals which maintain ownership of projects.
type Organization struct {
	Slug string
}

func getHomeView(c context.Base) error {
	return redirectToOrg(c)
}

func redirectToOrg(c context.Base) error {
	org, err := getActiveOrganization(c)
	if err != nil {
		return err
	}
	if err == nil {
		orgURI := c.Echo().URI(getOrganizationHomeView, org.Slug)
		return c.Redirect(http.StatusFound, orgURI)
	}
	//} else if !features.Has("organizations:create") {
	//    return c.Render(http.StatusForbidden "sentry/no-organization-access", nil)
	//}
	return c.HTML(http.StatusNotImplemented, "sentry-create-organization page is not implemented")
}

func getActiveOrganization(c context.Base) (*models.Organization, error) {
	// TODO get active organization for current user
	// TODO this method should take an optional organizationSlug argument
	orgSlug := c.Param("organization_slug")
	userID := 1 // TODO get ID from context.request.user.id
	onlyVisible := true
	org := models.Organization{}
	query := c.Tx.SelectBySql(`
		select o.*
		from sentry_organization o
			join sentry_organizationmember om on o.id = om.organization_id
		where om.user_id = ?`,
		userID)
	if onlyVisible {
		query = query.Where("o.status = ?", models.OrganizationStatusVisible)
	}
	// TODO if scope then filter out results by it
	if orgSlug != "" {
		query = query.Where("o.slug = ?", orgSlug)
	}
	/*
		V1
		if err := query.Limit(1).LoadStruct(&org); err != nil {
			return nil, errors.Wrap(err, "can not get active organization")
		}
		return &org, nil

		V2
		err = query.Limit(1).LoadStruct(&org)
		err = errors.Wrap(err, "can not get active organization")
		return &org, err
	*/
	err := query.Limit(1).LoadStruct(&org)
	err = errors.Wrap(err, "can not get active organization")
	return &org, err
}
