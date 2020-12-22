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

func registerAPI(app *core.Application) {
	e := app.Engine.Group(APIPrefix, middleware.CORS())

	app.RegisterHandler(e.GET, "test", func(c *context.AppContext) error {
		return c.SuccessResp(context.RespData{
			"hello": "world",
		})
	})
}
