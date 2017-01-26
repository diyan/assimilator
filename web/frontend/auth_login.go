package frontend

import (
	"github.com/labstack/echo"
)

func authRegisterGetView(c echo.Context) error {
	return authLoginGetView(c)
}

func authRegisterPostView(c echo.Context) error {
	return authLoginPostView(c)
}

func authLoginGetView(c echo.Context) error {
	return renderNotImplemented(c)
}

func authLoginPostView(c echo.Context) error {
	return renderNotImplemented(c)
}
