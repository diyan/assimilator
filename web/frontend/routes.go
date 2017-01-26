package frontend

import "github.com/labstack/echo"

// RegisterFrontendRoutes add routes with frontend views to the Echo's root router
func RegisterFrontendRoutes(e *echo.Echo) {
	e.POST("/api/hooks/mailgun/inbound/", mailgunInboundWebhookPostView)
	e.POST("/api/hooks/release/:plugin_id/:project_id/:signature/", releaseWebhookPostView)
	g := e.Group("/api/embed/error-page/")
	g.GET("", errorPageEmbedGetView)
	g.POST("", errorPageEmbedPostView)

	// Auth
	g = e.Group("/auth/link/:organization_slug/")
	g.GET("", authOrganizationLoginGetView)
	g.POST("", authOrganizationLoginPostView)
	g = e.Group("/auth/login/")
	g.GET("", authLoginGetView)
	g.POST("", authLoginPostView)
	g = e.Group("/auth/login/:organization_slug)/")
	g.GET("", authOrganizationLoginGetView)
	g.POST("", authOrganizationLoginPostView)
	g = e.Group("/auth/2fa/")
	g.GET("", twoFactorAuthGetView)
	g.POST("", twoFactorAuthPostView)
	e.GET("/auth/2fa/u2fappid.json", u2FAppID) // see sentry.web.frontend.twofactor.u2f_appid
	g = e.Group("/auth/sso/")
	g.GET("", authProviderLoginGetView)
	g.POST("", authProviderLoginPostView)
	e.GET("/auth/logout/", authLogoutGetView)
	g = e.Group("/auth/reactivate/")
	g.GET("", reactivateAccountGetView)
	g.POST("", reactivateAccountPostView)
	g = e.Group("/auth/register/")
	g.GET("", authRegisterGetView) // same as getAuthLoginView
	g.POST("", authRegisterPostView)

	// Account
	e.GET("/login-redirect/", loginRedirect)
	g = e.Group("/account/sudo/")
	g.GET("", sudoGetView)
	g.POST("", sudoPostView)
	e.GET("/account/confirm-email/", accountStartConfirmEmail)
	// :user_id [\d]+
	// :hash [0-9a-zA-Z]+
	e.GET("/account/confirm-email/:user_id/:hash/", accountConfirmEmail)
	g = e.Group("/account/recover/")
	g.GET("", accountRecoverGetView)
	g.POST("", accountRecoverPostView)
	g = e.Group("/account/recover/confirm/:user_id/:hash/")
	g.GET("", accountRecoverConfirmGetView)
	g.POST("", accountRecoverConfirmPostView)
	g = e.Group("/account/settings/")
	g.GET("", accountSettingsGetView)
	g.POST("", accountSettingsPostView)
	g = e.Group("/account/settings/2fa/")
	g.GET("", twoFactorSettingsGetView)
	g.POST("", twoFactorSettingsPostView)
	g = e.Group("/account/settings/2fa/recovery/")
	g.GET("", recoveryCodeSettingsGetView)
	g.POST("", recoveryCodeSettingsPostView)
	g = e.Group("/account/settings/2fa/totp/")
	g.GET("", tOTPSettingsGetView)
	g.POST("", tOTPSettingsPostView)
	g = e.Group("/account/settings/2fa/sms/")
	g.GET("", smsSettingsGetView)
	g.POST("", smsSettingsPostView)
	g = e.Group("/account/settings/2fa/u2f/")
	g.GET("", u2FSettingsGetView)
	g.POST("", u2FSettingsPostView)
	e.GET("/account/settings/avatar/", avatarSettings)
	g = e.Group("/account/settings/appearance/")
	g.GET("", appearanceSettingsGetView)
	g.POST("", appearanceSettingsPostView)
	e.GET("/account/settings/identities/", listIdentities)
	// :identity_id [^\/]+
	g = e.Group("/account/settings/identities/:identity_id/disconnect/")
	g.GET("", disconnectIdentityGetView)
	g.POST("", disconnectIdentityPostView)
	g = e.Group("/account/settings/notifications/")
	g.GET("", accountNotificationGetView)
	g.POST("", accountNotificationPostView)
	g = e.Group("/account/settings/security/")
	g.GET("", accountSecurityGetView)
	g.POST("", accountSecurityPostView)
	g = e.Group("/account/settings/emails/")
	g.GET("", showEmailsGetView)
	g.POST("", showEmailsPostView)

	// compatibility
	g = e.Group("/account/settings/notifications/unsubscribe/:project_id/")
	g.GET("", emailUnsubscribeProjectGetView)
	g.POST("", emailUnsubscribeProjectPostView)

	g = e.Group("/account/notifications/unsubscribe/:project_id/")
	g.GET("", emailUnsubscribeProjectGetView)
	g.POST("", emailUnsubscribeProjectPostView)
	// :issue_id \d+
	g = e.Group("/account/notifications/unsubscribe/issue/:issue_id/")
	g.GET("", unsubscribeIssueNotificationsGetView)
	g.GET("", unsubscribeIssueNotificationsPostView)

	g = e.Group("/account/remove/")
	g.GET("", removeAccountGetView)
	g.POST("", removeAccountPostView)
	// TODO mount social auth handlers under /account/settings/social/* prefix
	//e.GET("/account/settings/social/*', include('social_auth.urls'))

	// Admin
	// TODO HIGH. Review the rest of Sentry handlers an update routes with correct stubs
	e.GET("/manage/queue/", getAdminQueueView)
	e.GET("/manage/status/environment/", getStatusEnv)
	e.GET("/manage/status/packages/", getStatusPackages)
	e.GET("/manage/status/mail/", getStatusMail)
	e.GET("/manage/status/warnings/", getStatusWarnings)

	// Admin - Users
	e.GET("/manage/users/new/", getCreateNewUser)
	e.GET("/manage/users/:user_id/", getEditUser)
	e.GET("/manage/users/:user_id/remove/", getRemoveUser)

	// Admin - Plugins
	// :slug [\w_-]+
	e.GET("/manage/plugins/:slug/", getConfigurePlugin)

	// e.GET("/", getReactPageView)
	e.GET("/manage/*", getSentryAdminOverview)

	e.GET("/api/", getSentryAPI)
	e.GET("/api/new-token/", getReactPageView)

	e.GET("/out/", getOutView)

	// Organizations
	// :organization_slug [\w_-]+
	e.GET("/:organization_slug/", getOrganizationHomeView)

	e.GET("/organizations/new/", getCreateOrganizationView)
	e.GET("/organizations/:organization_slug/api-keys/", getOrganizationApiKeysView)
	// :key_id [\w_-]+
	e.GET("/organizations/:organization_slug/api-keys/:key_id/", getOrganizationApiKeySettingsView)
	e.GET("/organizations/:organization_slug/auth/", getOrganizationAuthSettingsView)
	e.GET("/organizations/:organization_slug/members/", getOrganizationMembersView)
	e.GET("/organizations/:organization_slug/members/new/", getCreateOrganizationMemberView)
	e.GET("/organizations/:organization_slug/members/:member_id/", getOrganizationMemberSettingsView)
	e.GET("/organizations/:organization_slug/stats/", getSentryOrganizationStatsView)
	e.GET("/organizations/:organization_slug/settings/", getOrganizationSettingsView)
	// :team_slug [\w_-]+
	e.GET("/organizations/:organization_slug/teams/:team_slug/remove/", getRemoveTeamView)
	e.GET("/organizations/:organization_slug/teams/new/", getCreateTeamView)
	e.GET("/organizations/:organization_slug/projects/new/", getCreateProjectView)
	e.GET("/organizations/:organization_slug/remove/", getRemoveOrganizationView)
	// TODO CONTINUE AT THIS PONT. Develop not-implemented handlers in correct golang files
	e.GET("/organizations/:organization_slug/restore/", getRestoreOrganizationView)
	// :token \w+
	e.GET("/accept/:member_id/:token/", getAcceptOrganizationInviteView)

	// Settings - Projects
	// :project_slug [\w_-]+
	e.GET("/:organization_slug/:project_slug/settings/", getProjectSettingsView)
	e.GET("/:organization_slug/:project_slug/settings/issue-tracking/", getProjectIssueTrackingView)
	e.GET("/:organization_slug/:project_slug/settings/release-tracking/", getProjectReleaseTrackingView)
	e.GET("/:organization_slug/:project_slug/settings/keys/", getProjectKeysView)
	e.GET("/:organization_slug/:project_slug/settings/keys/new/", getCreateProjectKeyView)
	// :key_id \d+
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/edit/", getEditProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/remove/", getRemoveProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/disable/", getDisableProjectKeyView)
	e.GET("/:organization_slug/:project_slug/settings/keys/:key_id/enable/", getEnableProjectKeyView)

	e.GET("/:organization_slug/:project_slug/settings/plugins/", getProjectPluginsView)
	// :slug [\w_-]+
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/", getProjectPluginConfigureView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/reset/", getProjectPluginResetView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/disable/", getProjectPluginDisableView)
	e.GET("/:organization_slug/:project_slug/settings/plugins/:slug/enable/", getProjectPluginEnableView)

	e.GET("/:organization_slug/:project_slug/settings/remove/", getRemoveProjectView)

	e.GET("/:organization_slug/:project_slug/settings/tags/", getProjectTagsView)

	e.GET("/:organization_slug/:project_slug/settings/quotas/", getProjectQuotasView)

	e.GET("/:organization_slug/:project_slug/settings/alerts/rules/new/", getProjectRuleEditView)
	// :rule_id \d+
	e.GET("/:organization_slug/:project_slug/settings/alerts/rules/:rule_id/", getProjectRuleEditView)

	// :avatar_id [^\/]+
	e.GET("/avatar/:avatar_id/", getUserAvatarPhotoView)

	// Generic
	e.GET("/", getHomeView)

	// Generic API
	// TODO disable auth on ReactPageView for 2 routes below
	// :share_id [\w_-]+
	e.GET("/share/group/share_id/", getGenericReactPageView)
	e.GET("/share/issue/share_id/", getGenericReactPageView)

	// Keep named URL for for things using reverse
	// :short_id [\w_-]+
	e.GET("/:organization_slug/issues/:short_id/", getSentryShortIDView)
	// :group_id \d+
	e.GET("/:organization_slug/:project_id/issues/:group_id/", getSentryGroupView)
	e.GET("/:organization_slug/:project_id/", getSentryStream)

	// :event_id_or_latest (\d+|latest)
	e.GET("/:organization_slug/:project_slug/group/:group_id/events/:event_id_or_latest/json/", getGroupEventJsonView)
	e.GET("/:organization_slug/:project_slug/issues/:group_id/events/:event_id_or_latest/json/", getGroupEventJsonView)
	// :key [^\/]+
	e.GET("/:organization_slug/:project_slug/issues/:group_id/tags/:key/export/", getGroupTagExportView)
	e.GET("/:organization_slug/:project_slug/issues/:group_id/actions/:slug/", getGroupPluginActionView)
}
