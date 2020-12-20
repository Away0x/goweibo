package routes

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func registerError(e *echo.Echo) {
	echo.NotFoundHandler = notFoundHandler
	echo.MethodNotAllowedHandler = notFoundHandler

	// e.HTTPErrorHandler = func(e error, c echo.Context) {}
}

func notFoundHandler(c echo.Context) error {
	return errors.New("not found")
}
