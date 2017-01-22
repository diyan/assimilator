package frontend

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetOrganizationHomeView(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryAdminOverview(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryAPI(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryOrganizationStatsView(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryShortIDView(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryGroupView(c echo.Context) error {
	return GetReactPageView(c)
}

func GetSentryStream(c echo.Context) error {
	return GetReactPageView(c)
}

func GetGenericReactPageView(c echo.Context) error {
	return GetReactPageView(c)
}

func GetReactPageView(c echo.Context) error {
	// restrict orgSlug to the regex [\w_-]+
	//orgSlug := c.Param("organization-slug")
	//log.Print(orgSlug, " is an organization-slug")
	// TOOD implement generic ReactPageView
	return c.Render(http.StatusOK, "react.html", map[string]interface{}{
		"request":        struct{ LANGUAGE_CODE string }{LANGUAGE_CODE: "en"},
		"sentry_version": struct{ build string }{build: "8.12.0"},
		// TODO this does not work because foo is not exported
		//"sentry_version":   template.SentryVersion{foo: "8.12.0"},
		"CSRF_COOKIE_NAME": "sc", // TODO Move into settings
		"ALLOWED_HOSTS":    "",
	})
}
