package conf

import (
	"github.com/labstack/echo"
	"github.com/pkg/errors"
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
	if conf, ok := c.Get("conf.Config").(Config); ok {
		return conf
	}
	panic(errors.New("failed to get app config, is conf.NewMiddleware(config) call missing?"))
}

// TODO Add func FromC(c context.Context) Config

// NewMiddleware creates middleware that sets app config for each HTTP request
func NewMiddleware(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("conf.Config", config)
			return next(c)
		}
	}
}
