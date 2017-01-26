package debug

import (
	"net/http"

	"github.com/labstack/echo"
)

func renderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"We are sorry, this page was not yet ported from Sentry")
}

func debugMailAlertGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugNoteEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugNewReleaseEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugAssignedEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugSelfAssignedEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailDigestGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailReportGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugRegressionEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugRegressionReleaseEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugResolvedEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugResolvedInReleaseEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugResolvedInReleaseUpcomingEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailRequestAccess(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailAccessApproved(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailAccessInvitation(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailConfirmEmail(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailRecoverAccount(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugUnassignedEmailGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugMailOrgDeleteConfirmGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugErrorPageEmbedGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugTriggerErrorGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugAuthConfirmIdentity(c echo.Context) error {
	return renderNotImplemented(c)
}

func debugAuthConfirmLink(c echo.Context) error {
	return renderNotImplemented(c)
}

func iconsGetView(c echo.Context) error {
	return renderNotImplemented(c)
}
