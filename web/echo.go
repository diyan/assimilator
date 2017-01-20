package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/k0kubun/pp"
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
		if c.Request().Method == http.MethodHead { // Issue #608
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
					var err error
					switch r := r.(type) {
					case error:
						err = r
					default:
						err = fmt.Errorf("%v", r)
					}
					pp.Print(err)
					//stack := make([]byte, config.StackSize)
					//length := runtime.Stack(stack, !config.DisableStackAll)
					//if !config.DisablePrintStack {
					//	c.Logger().Printf("[%s] %s %s\n", color.Red("PANIC RECOVER"), err, stack[:length])
					//}
					//c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
