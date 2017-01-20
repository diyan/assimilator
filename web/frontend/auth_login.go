package frontend

import (
	"github.com/labstack/echo"
)

func AuthRegisterGetView(c echo.Context) error {
	return AuthLoginGetView(c)
}

func AuthRegisterPostView(c echo.Context) error {
	return AuthLoginPostView(c)
}

func AuthLoginGetView(c echo.Context) error {
	return RenderNotImplemented(c)
}

func AuthLoginPostView(c echo.Context) error {
	return RenderNotImplemented(c)
}
