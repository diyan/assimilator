package conf

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type contextKey string

func (c contextKey) String() string {
	return "config context key " + string(c)
}

var (
	contextKeyConfig = contextKey("config")
)

// Config struct holds all application settings
type Config struct {
	Port            int    `mapstructure:"port"`
	DatabaseURL     string `mapstructure:"db_url"`
	InitialTeam     string `mapstructure:"initial_team"`
	InitialProject  string `mapstructure:"initial_project"`
	InitialKey      string `mapstructure:"initial_key"`
	InitialPlatform string `mapstructure:"initial_platform"`
}

// FromEC returns app config associated with echo's Context
func FromEC(c echo.Context) Config {
	if conf, ok := c.Request().Context().Value(contextKeyConfig).(Config); ok {
		return conf
	}
	panic(errors.New("failed to get app config, is conf.NewMiddleware(config) call missing?"))
}

// NewMiddleware creates middleware that sets app config for each HTTP request
func NewMiddleware(config Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := context.WithValue(req.Context(), contextKeyConfig, config)
			next.ServeHTTP(rw, req.WithContext(ctx))
		})
	}
}
