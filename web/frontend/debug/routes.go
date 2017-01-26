package debug

import (
	"github.com/labstack/echo"
)

// RegisterDebugViews adds routes with debug views to the Echo's route group
func RegisterDebugViews(g *echo.Group) {
	// TODO stubs were created only for GET verbs
	g.GET("/mail/alert/", debugMailAlertGetView)
	g.GET("/mail/note/", debugNoteEmailGetView)
	g.GET("/mail/new-release/", debugNewReleaseEmailGetView)
	g.GET("/mail/assigned/", debugAssignedEmailGetView)
	g.GET("/mail/assigned/self/", debugSelfAssignedEmailGetView)
	g.GET("/mail/digest/", debugMailDigestGetView)
	g.GET("/mail/report/", debugMailReportGetView)
	g.GET("/mail/regression/", debugRegressionEmailGetView)
	g.GET("/mail/regression/release/", debugRegressionReleaseEmailGetView)
	g.GET("/mail/resolved/", debugResolvedEmailGetView)
	g.GET("/mail/resolved-in-release/", debugResolvedInReleaseEmailGetView)
	g.GET("/mail/resolved-in-release/upcoming/", debugResolvedInReleaseUpcomingEmailGetView)
	g.GET("/mail/request-access/", debugMailRequestAccess)
	g.GET("/mail/access-approved/", debugMailAccessApproved)
	g.GET("/mail/invitation/", debugMailAccessInvitation)
	g.GET("/mail/confirm-email/", debugMailConfirmEmail)
	g.GET("/mail/recover-account/", debugMailRecoverAccount)
	g.GET("/mail/unassigned/", debugUnassignedEmailGetView)
	g.GET("/mail/org-delete-confirm/", debugMailOrgDeleteConfirmGetView)
	g.GET("/embed/error-page/", debugErrorPageEmbedGetView)
	g.GET("/trigger-error/", debugTriggerErrorGetView)
	g.GET("/auth-confirm-identity/", debugAuthConfirmIdentity)
	g.GET("/auth-confirm-link/", debugAuthConfirmLink)
	g.GET("/icons/", iconsGetView)
}
