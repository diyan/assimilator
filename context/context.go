package context

import (
	"github.com/diyan/assimilator/models"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
)

// TODO consider add conf.Config
type Base struct {
	echo.Context
	Tx   *dbr.Tx
	User models.User
}

type Organization struct {
	Base
	Organization models.Organization
}

// TODO How to embed context.Organization, so it won't shadow models.Organization
type Project struct {
	Base
	Organization models.Organization
	Project      models.Project
}
