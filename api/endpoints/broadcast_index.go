package api

import "github.com/labstack/echo"

func BroadcastIndexGetEndpoint(c echo.Context) error {
	c.Logger().Debug("not yet implemented")
	// TODO Workaround to fix 'this.state.broadcasts.filter is not a function' error
	return c.JSON(200, []int{})
}

func BroadcastIndexPutEndpoint(c echo.Context) error {
	c.Logger().Debug("not yet implemented")
	return c.NoContent(200)
}
