package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// SystemHealth ...
type SystemHealth struct {
	Healthy  SystemHealthStatus `json:"healthy"`
	Problems []string           `json:"problems"` // TODO []SystemHealthProblem
}

// SystemHealthStatus ..
type SystemHealthStatus struct {
	CeleryAppVersionCheck bool `json:"CeleryAppVersionCheck"`
	CeleryAliveCheck      bool `json:"CeleryAliveCheck"`
}

// TODO replace stub with real implementation
func SystemHealthGetEndpoint(c echo.Context) error {
	health := SystemHealth{
		Healthy: SystemHealthStatus{
			CeleryAppVersionCheck: true, CeleryAliveCheck: true},
		Problems: []string{},
	}
	return c.JSON(http.StatusOK, health)

}

