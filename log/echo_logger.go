package log

import (
	"github.com/diyan/assimilator/conf"
	"github.com/diyan/echox/log"
	"github.com/labstack/echo"
)

// NewEchoLogger returns logger for the Echo web framework
func NewEchoLogger(config conf.Config) echo.Logger {
	return log.Logrus()
}
