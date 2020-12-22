package routes

import (
	"goweibo/core"
	"goweibo/core/context"
	"time"
)

func registerWeb(router *core.Application) {
	router.RegisterHandler(router.GET, "welcome", func(c *context.AppContext) error {
		now := time.Now()

		return c.AWHtml("welcome", context.TplData{
			"time": now.Format("2006-01-02"),
		})
	}).Name = "welcome"
}
