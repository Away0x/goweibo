package routes

import (
	"github.com/labstack/echo/v4"
)

func registerWeb(e *echo.Echo) {
	ee := e.Group("")

	ee.GET("", func(c echo.Context) error {
		return c.JSON(200, 123)
	})
}
