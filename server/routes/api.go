package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// APIPrefix api prefix
	APIPrefix = "/api"
)

func registerAPI(e *echo.Echo) {
	e.Group(APIPrefix, middleware.CORS())
}
