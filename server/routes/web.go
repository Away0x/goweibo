package routes

import (
	"goweibo/core"
	"goweibo/core/context"
	"time"
)

func registerWeb(app *core.Application) {
	e := app.Engine

	app.RegisterHandler(e.GET, "welcome", func(c *context.AppContext) error {
		now := time.Now()

		return c.RenderHTML("welcome", context.TplData{
			"time": now.Format("2006-01-02"),
		})
	})
}
