package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	simpleflake "github.com/intelekshual/go-simpleflake"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := http.StatusText(code)
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if we, ok := err.(stackTracer); ok {
		fmt.Fprintf(os.Stderr, "\n%+v\n\n", we)
	}
	//if e.Debug {
	//	msg = err.Error()
	//}
	msg = err.Error()
	if !c.Response().Committed {
		// TODO Consider use `http.MethodHead` from Go 1.6+ and drop Go 1.5
		if c.Request().Method == "HEAD" { // Issue #608
			c.NoContent(code)
		} else {
			c.String(code, msg)
		}
	}
	//e.Logger.Error(err)

	//logrus.ErrorLevel()
	//debug.PrintStack()
	//pp.Print(err)
	logrus.Error(err)
}

func RecoverMiddleware() echo.MiddlewareFunc {
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
						err = e
					} else {
						err = fmt.Errorf("%v", r)
					}

					//stack := make([]byte, config.StackSize)
					//length := runtime.Stack(stack, !config.DisableStackAll)
					//if !config.DisablePrintStack {
					//	c.Logger().Printf("[%s] %s %s\n", color.Red("PANIC RECOVER"), err, stack[:length])
					//}
					//c.Error(err)
					logrus.WithFields(logEvent).Error(err)
				}
			}()
			return next(c)
		}
	}
}
