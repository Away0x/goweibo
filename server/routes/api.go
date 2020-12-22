package routes

import (
	"goweibo/core"
	"goweibo/core/context"

	"github.com/labstack/echo/v4/middleware"
)

const (
	// APIPrefix api prefix
	APIPrefix = "/api"
)

func registerAPI(router *core.Application) {
	e := router.Group(APIPrefix, middleware.CORS())

	router.RegisterHandler(e.GET, "test", func(c *context.AppContext) error {
		return c.AWSuccessJSON(context.RespData{
			"hello": "world",
		})
	})
}
