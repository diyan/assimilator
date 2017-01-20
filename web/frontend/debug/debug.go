package debug

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderNotImplemented(c echo.Context) error {
	return c.HTML(
		http.StatusNotImplemented,
		"We are sorry, this page was not yet ported from Sentry")
}

func DebugMailAlertGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugNoteEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugNewReleaseEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugAssignedEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugSelfAssignedEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailDigestGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailReportGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugRegressionEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugRegressionReleaseEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugResolvedEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugResolvedInReleaseEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugResolvedInReleaseUpcomingEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailRequestAccess(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailAccessApproved(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailAccessInvitation(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailConfirmEmail(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailRecoverAccount(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugUnassignedEmailGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugMailOrgDeleteConfirmGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugErrorPageEmbedGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugTriggerErrorGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugAuthConfirmIdentity(c echo.Context) error {
	return RenderNotImplemented(c)
}

func DebugAuthConfirmLink(c echo.Context) error {
	return RenderNotImplemented(c)
}

func IconsGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}
