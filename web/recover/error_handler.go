package recover

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/diyan/assimilator/conf"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// NewEchoErrorHandler creates HTTP error handler for the Echo web framework
func NewEchoErrorHandler(config conf.Config) func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
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
		logrus.Error(err)
	}
}
