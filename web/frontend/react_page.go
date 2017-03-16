package frontend

import (
	"net/http"

	"github.com/labstack/echo"
)

func getOrganizationHomeView(c echo.Context) error {
	return getReactPageView(c)
}

func getSentryAdminOverview(c echo.Context) error {
	return getReactPageView(c)
}

func getSentryAPI(c echo.Context) error {
	return getReactPageView(c)
}

func getSentryOrganizationStatsView(c echo.Context) error {
	return getReactPageView(c)
}

func getSentryShortIDView(c echo.Context) error {
	return getReactPageView(c)
}

func GetSentryGroupView(c echo.Context) error {
	return getReactPageView(c)
}

func getSentryStream(c echo.Context) error {
	return getReactPageView(c)
}

func getGenericReactPageView(c echo.Context) error {
	return getReactPageView(c)
}

func getReactPageView(c echo.Context) error {
	// restrict orgSlug to the regex [\w_-]+
	//orgSlug := c.Param("organization-slug")
	//log.Print(orgSlug, " is an organization-slug")
	// TODO implement generic ReactPageView
	return c.Render(http.StatusOK, "sentry/bases/react.html", map[string]interface{}{
		"request":        struct{ LANGUAGE_CODE string }{LANGUAGE_CODE: "en"},
		"sentry_version": struct{ build string }{build: "8.12.0"},
		// TODO this does not work because foo is not exported
		//"sentry_version":   template.SentryVersion{foo: "8.12.0"},
		"CSRF_COOKIE_NAME": "sc", // TODO Move into settings
		"ALLOWED_HOSTS":    "",
	})
}
