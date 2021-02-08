package routes

import (
	"goweibo/core"
	"goweibo/core/context"
  "goweibo/core/pkg/timeutils"
  "time"
)

func registerWeb(router *core.Application) {
	router.RegisterHandler(router.GET, "welcome", func(c *context.AppContext) error {
		return c.AWHtml("welcome", context.TplData{
			"time": timeutils.FormatTime(time.Now()),
		})
	}).Name = "welcome"
}
