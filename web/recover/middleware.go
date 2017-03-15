package recover

import (
	"fmt"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/conf"
	simpleflake "github.com/intelekshual/go-simpleflake"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// NewMiddleware creates middleware that recover from panics
func NewMiddleware(config conf.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					req := c.Request()
					logEvent := logrus.Fields{
						"event": "internal_error", // TODO consider internalError naming
						// TODO set request_id and other request-specific into in the separate middleware
						"host":   req.Host,
						"uri":    req.URL.String(),
						"method": req.Method,
						//"path":   path,
						// TODO check X-Real-IP and X-Forwarded-For headers
						"remote_ip": req.RemoteAddr,
					}
					eventID, err := simpleflake.New()
					// TODO generate request_id in separate logging middleware
					logEvent["request_id"] = eventID
					logEvent["event_id"] = eventID
					if err != nil {
						logrus.Error(errors.Wrap(err, "failed to generate simpleflake ID"))
					}
					if _, ok := r.(stackTracer); ok {
						// stackTracer interface means that error was handled gracefully
						// TODO add more domain-specific error interface
						logEvent["event"] = "app_error"
						logEvent["error_msg"] = fmt.Sprintf("%v", r)
						err = fmt.Errorf("%+v\n", r)
					} else if e, ok := r.(error); ok {
						logEvent["event"] = "internal_error"
						stack := make([]byte, 1<<20)
						stackLen := runtime.Stack(stack, false)
						err = fmt.Errorf("%+v\n%s", e, stack[:stackLen])
					} else {
						logEvent["event"] = "internal_error"
						err = fmt.Errorf("%v", r)
					}

					//c.Error(err)
					logrus.WithFields(logEvent).Error(err)
				}
			}()
			return next(c)
		}
	}
}
