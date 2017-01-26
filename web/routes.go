package web

import (
	"net/http"

	apiV0 "github.com/diyan/assimilator/api"
	"github.com/diyan/assimilator/web/api"
	"github.com/diyan/assimilator/web/frontend"
	"github.com/diyan/assimilator/web/frontend/debug"

	"github.com/labstack/echo"
)

/*
APIView base class has HTTP OPTIONS handler that respond with actual list of supported methods


*/

/*
e.Any("/api/store/", api.StoreView)

func StoreView(c echo.Context) {
	if c.Request().Method() == "POST" {
		return postStore(c)
	} else if c.Request().Method() == "GET" {
		return getStore(c)
	} else {
		call next middleware
	}
}

store := e.Group("/api/store/")
store.Get("", api.GetStoreView)
store.Post("", api.PostStoreView)

store.Get("", api.StoreGetView)
store.Post("", api.StorePostView)


store.Any("", api.AnyStoreView)
*/

// RegisterRoutes ..
func RegisterRoutes(e *echo.Echo) {
	// TODO call registerDebugViews only if getattr(settings, 'DEBUG_VIEWS', settings.DEBUG)
	g := e.Group("/debug")
	debug.RegisterDebugViews(g)
	// The static version is either a 10 digit timestamp, a sha1, or md5 hash
	// :version \d{10}|[a-f0-9]{32,40}
	// TODO Use general-purpose static middleware or custom implementation?
	//e.GET("/_static/:module/*", frontend.GetStaticMedia)
	//e.GET("/_static/:version/:module/*", frontend.GetStaticMedia)

	// API
	g = e.Group("/api")
	api.RegisterAPIRoutes(g)
	// API version 0
	g = e.Group("/api/0")
	apiV0.RegisterAPIRoutes(g)
	e.POST("/api/hooks/mailgun/inbound/", frontend.MailgunInboundWebhookPostView)
	e.POST("/api/hooks/release/:plugin_id/:project_id/:signature/", frontend.ReleaseWebhookPostView)
	g = e.Group("/api/embed/error-page/")
	g.GET("", frontend.ErrorPageEmbedGetView)
	g.POST("", frontend.ErrorPageEmbedPostView)

	// Auth
	g = e.Group("/auth/link/:organization_slug/")
	g.GET("", frontend.AuthOrganizationLoginGetView)
	g.POST("", frontend.AuthOrganizationLoginPostView)
	g = e.Group("/auth/login/")
	g.GET("", frontend.AuthLoginGetView)
	g.POST("", frontend.AuthLoginPostView)
	g = e.Group("/auth/login/:organization_slug)/")
	g.GET("", frontend.AuthOrganizationLoginGetView)
	g.POST("", frontend.AuthOrganizationLoginPostView)
	g = e.Group("/auth/2fa/")
	g.GET("", frontend.TwoFactorAuthGetView)
	g.POST("", frontend.TwoFactorAuthPostView)
	e.GET("/auth/2fa/u2fappid.json", frontend.U2FAppID) // see sentry.web.frontend.twofactor.u2f_appid
	g = e.Group("/auth/sso/")
	g.GET("", frontend.AuthProviderLoginGetView)
	g.POST("", frontend.AuthProviderLoginPostView)
	e.GET("/auth/logout/", frontend.AuthLogoutGetView)
	g = e.Group("/auth/reactivate/")
	g.GET("", frontend.ReactivateAccountGetView)
	g.POST("", frontend.ReactivateAccountPostView)
	g = e.Group("/auth/register/")
	g.GET("", frontend.AuthRegisterGetView) // same as getAuthLoginView
	g.POST("", frontend.AuthRegisterPostView)

	// Account
	e.GET("/login-redirect/", frontend.LoginRedirect)
	g = e.Group("/account/sudo/")
	g.GET("", frontend.SudoGetView)
	g.POST("", frontend.SudoPostView)
	e.GET("/account/confirm-email/", frontend.AccountStartConfirmEmail)
	// :user_id [\d]+
	// :hash [0-9a-zA-Z]+
	e.GET("/account/confirm-email/:user_id/:hash/", frontend.AccountConfirmEmail)
	g = e.Group("/account/recover/")
	g.GET("", frontend.AccountRecoverGetView)
	g.POST("", frontend.AccountRecoverPostView)
	g = e.Group("/account/recover/confirm/:user_id/:hash/")
	g.GET("", frontend.AccountRecoverConfirmGetView)
	g.POST("", frontend.AccountRecoverConfirmPostView)
	g = e.Group("/account/settings/")
	g.GET("", frontend.AccountSettingsGetView)
	g.POST("", frontend.AccountSettingsPostView)
	g = e.Group("/account/settings/2fa/")
	g.GET("", frontend.TwoFactorSettingsGetView)
	g.POST("", frontend.TwoFactorSettingsPostView)
	g = e.Group("/account/settings/2fa/recovery/")
	g.GET("", frontend.RecoveryCodeSettingsGetView)
	g.POST("", frontend.RecoveryCodeSettingsPostView)
	g = e.Group("/account/settings/2fa/totp/")
	g.GET("", frontend.TOTPSettingsGetView)
	g.POST("", frontend.TOTPSettingsPostView)
	g = e.Group("/account/settings/2fa/sms/")
	g.GET("", frontend.SmsSettingsGetView)
	g.POST("", frontend.SmsSettingsPostView)
	g = e.Group("/account/settings/2fa/u2f/")
	g.GET("", frontend.U2FSettingsGetView)
	g.POST("", frontend.U2FSettingsPostView)
	e.GET("/account/settings/avatar/", frontend.AvatarSettings)
	g = e.Group("/account/settings/appearance/")
	g.GET("", frontend.AppearanceSettingsGetView)
	g.POST("", frontend.AppearanceSettingsPostView)
	e.GET("/account/settings/identities/", frontend.ListIdentities)
	// :identity_id [^\/]+
	g = e.Group("/account/settings/identities/:identity_id/disconnect/")
	g.GET("", frontend.DisconnectIdentityGetView)
	g.POST("", frontend.DisconnectIdentityPostView)
	g = e.Group("/account/settings/notifications/")
	g.GET("", frontend.AccountNotificationGetView)
	g.POST("", frontend.AccountNotificationPostView)
	g = e.Group("/account/settings/security/")
	g.GET("", frontend.AccountSecurityGetView)
	g.POST("", frontend.AccountSecurityPostView)
	g = e.Group("/account/settings/emails/")
	g.GET("", frontend.ShowEmailsGetView)
	g.POST("", frontend.ShowEmailsPostView)

	// compatibility
	g = e.Group("/account/settings/notifications/unsubscribe/:project_id/")
	g.GET("", frontend.EmailUnsubscribeProjectGetView)
	g.POST("", frontend.EmailUnsubscribeProjectPostView)

	g = e.Group("/account/notifications/unsubscribe/:project_id/")
	g.GET("", frontend.EmailUnsubscribeProjectGetView)
	g.POST("", frontend.EmailUnsubscribeProjectPostView)
	// :issue_id \d+
	g = e.Group("/account/notifications/unsubscribe/issue/:issue_id/")
	g.GET("", frontend.UnsubscribeIssueNotificationsGetView)
	g.GET("", frontend.UnsubscribeIssueNotificationsPostView)

	g = e.Group("/account/remove/")
	g.GET("", frontend.RemoveAccountGetView)
	g.POST("", frontend.RemoveAccountPostView)
	// TODO mount social auth handlers under /account/settings/social/* prefix
	//e.GET("/account/settings/social/*', include('social_auth.urls'))

	// Admin
	// TODO HIGH. Review the rest of Sentry handlers an update routes with correct stubs
	e.GET("/manage/queue/", frontend.GetAdminQueueView)
	e.GET("/manage/status/environment/", frontend.GetStatusEnv)
	e.GET("/manage/status/packages/", frontend.GetStatusPackages)
	e.GET("/manage/status/mail/", frontend.GetStatusMail)
	e.GET("/manage/status/warnings/", frontend.GetStatusWarnings)

	// Admin - Users
	e.GET("/manage/users/new/", frontend.GetCreateNewUser)
	e.GET("/manage/users/:user_id/", frontend.GetEditUser)
	e.GET("/manage/users/:user_id/remove/", frontend.GetRemoveUser)

	// Admin - Plugins
	// :slug [\w_-]+
	e.GET("/manage/plugins/:slug/", frontend.GetConfigurePlugin)

	// e.GET("/", frontend.GetReactPageView)
	e.GET("/manage/*", frontend.GetSentryAdminOverview)

	// Legacy Redirects
	e.GET("/docs/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://docs.sentry.io/hosted/")
	})
	e.GET("/docs/api/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, "https://docs.sentry.io/hosted/api/")
	})

	e.GET("/api/", frontend.GetSentryAPI)
	e.GET("/api/new-token/", frontend.GetReactPageView)

	e.GET("/out/", frontend.GetOutView)

	// Organizations
	// :organization_slug [\w_-]+
	e.GET("/:organization_slug/", frontend.GetOrganizationHomeView)

	e.GET("/organizations/new/", frontend.GetCreateOrganizationView)
	e.GET("/organizations/:organization_slug/api-keys/", frontend.GetOrganizationApiKeysView)
	// :key_id [\w_-]+
	e.GET("/organizations/:organization_slug/api-keys/:key_id/", frontend.GetOrganizationApiKeySettingsView)
	e.GET("/organizations/:organization_slug/auth/", frontend.GetOrganizationAuthSettingsView)
	e.GET("/organizations/:organization_slug/members/", frontend.GetOrganizationMembersView)
	e.GET("/organizations/:organization_slug/members/new/", frontend.GetCreateOrganizationMemberView)
	e.GET("/organizations/:organization_slug/members/:member_id/", frontend.GetOrganizationMemberSettingsView)
	e.GET("/organizations/:organization_slug/stats/", frontend.GetSentryOrganizationStatsView)
	e.GET("/organizations/:organization_slug/settings/", frontend.GetOrganizationSettingsView)
	// :team_slug [\w_-]+
	e.GET("/organizations/:organization_slug/teams/:team_slug/remove/", frontend.GetRemoveTeamView)
	e.GET("/organizations/:organization_slug/teams/new/", frontend.GetCreateTeamView)
	e.GET("/organizations/:organization_slug/projects/new/", frontend.GetCreateProjectView)
	e.GET("/organizations/:organization_slug/remove/", frontend.GetRemoveOrganizationView)
	// TODO CONTINUE AT THIS PONT. Develop not-implemented handlers in correct golang files
	e.GET("/organizations/:organization_slug/restore/", frontend.GetRestoreOrganizationView)
	// :token \w+
	e.GET("/accept/:member_id/:token/", frontend.GetAcceptOrganizationInviteView)

	// Settings - Projects
	// :project_slug [\w_-]+
	e.GET("/:organization_slug/:project_slug/settings/", frontend.GetProjectSettingsView)
	e.GET("/:organization_slug/:project_slug/settings/issue-tracking/", frontend.GetProjectIssueTrackingView)
	e.GET("/:organization_slug/:project_slug/settings/release-tracking/", frontend.GetProjectReleaseTrackingView)
	e.GET("/:organization_slug/:project_slug/settings/keys/", frontend.GetProjectKeysView)
	e.GET("/:organization_slug/:project_slug/settings/keys/new/", frontend.GetCreateProjectKeyView)
	// :key_id \d+
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/edit/", frontend.GetEditProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/remove/", frontend.GetRemoveProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/disable/", frontend.GetDisableProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/enable/", frontend.GetEnableProjectKeyView)

	e.GET("/:organization_slug/:project_slug/settings/plugins/", frontend.GetProjectPluginsView)
	// :slug [\w_-]+
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/", frontend.GetProjectPluginConfigureView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/reset/", frontend.GetProjectPluginResetView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/disable/", frontend.GetProjectPluginDisableView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/enable/", frontend.GetProjectPluginEnableView)

	e.GET("/:organization_slug/:project_slug/settings/remove/", frontend.GetRemoveProjectView)

	e.GET("/:organization_slug/:project_slug/settings/tags/", frontend.GetProjectTagsView)

	e.GET("/:organization_slug/:project_slug/settings/quotas/", frontend.GetProjectQuotasView)

	e.GET("/:organization_slug/:project_slug/settings/alerts/rules/new/", frontend.GetProjectRuleEditView)
	// :rule_id \d+
	e.GET("/:organization_slug/:project_slug/settings/alerts/rules/:rule_id/", frontend.GetProjectRuleEditView)

	// :avatar_id [^\/]+
	e.GET("/avatar/:avatar_id/", frontend.GetUserAvatarPhotoView)

	// Generic
	e.GET("/", frontend.GetHomeView)

	e.GET("/robots.txt", api.GetRobotsTxt)

	// Force a 404 of favicon.ico.
	// This url is commonly requested by browsers, and without
	// blocking this, it was treated as a 200 OK for a react page view.
	// A side effect of this is it may cause a bad redirect when logging in
	// since this gets stored in session as the last viewed page.
	// See: https://github.com/getsentry/sentry/issues/2195
	e.GET("/favicon.ico", func(c echo.Context) error {
		return c.NoContent(404)
	})
	// crossdomain.xml
	e.GET("/crossdomain.xml", api.GetCrossdomainXMLIndex)
	e.GET("/api/:project_id/crossdomain.xml", api.GetCrossdomainXML)

	// plugins
	// TODO mount plugin handlers under /plugins/* prefix
	// e.GET("/plugins/*", include('sentry.plugins.base.urls'))

	// Generic API
	// TODO disable auth on ReactPageView for 2 routes below
	// :share_id [\w_-]+
	e.GET("/share/group/share_id/", frontend.GetGenericReactPageView)
	e.GET("/share/issue/share_id/", frontend.GetGenericReactPageView)

	// Keep named URL for for things using reverse
	// :short_id [\w_-]+
	e.GET("/:organization_slug/issues/:short_id/", frontend.GetSentryShortIDView)
	// :group_id \d+
	e.GET("/:organization_slug/:project_id/issues/:group_id/", frontend.GetSentryGroupView)
	e.GET("/:organization_slug/:project_id/", frontend.GetSentryStream)

	// :event_id_or_latest (\d+|latest)
	e.GET("/:organization_slug/:project_slug/group/:group_id/events/:event_id_or_latest/json/", frontend.GetGroupEventJsonView)
	e.GET("/:organization_slug/:project_slug/issues/:group_id/events/:event_id_or_latest/json/", frontend.GetGroupEventJsonView)
	// :key [^\/]+
	e.GET("/:organization_slug/:project_slug/issues/:group_id/tags/:key/export/", frontend.GetGroupTagExportView)
	e.GET("/:organization_slug/:project_slug/issues/:group_id/actions/:slug/", frontend.GetGroupPluginActionView)

	// Legacy
	// TODO check original implementation
	// url(r'/", react_page_view),
}
