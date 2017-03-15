package log

import (
	"github.com/diyan/assimilator/conf"

	mwx "github.com/diyan/echox/middleware"
	"github.com/labstack/echo"
)

// NewAccessLogMiddleware creates middleware that logs HTTP requests
func NewAccessLogMiddleware(config conf.Config) echo.MiddlewareFunc {
	// TODO add configuration flag to enable/disable access logs
	// TODO if disabled return either nil-func or noop middleware
	return mwx.LogrusLogger(nil)
}
