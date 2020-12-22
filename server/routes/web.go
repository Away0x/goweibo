package routes

import (
	"goweibo/core"
	"goweibo/core/context"
	"time"

	"goweibo/core/pkg/session"
)

func registerWeb(router *core.Application) {
	e := router.Group("", session.NewMiddleware(session.MiddlewareOptions{
		HttpOnly:    true,
		Path:        "/",
		MaxAge:      86400 * 30,
		SessionName: core.GetConfig().String("APP.NAME"),
		SessionKey:  core.GetConfig().String("APP.KEY"),
	}))

	router.RegisterHandler(e.GET, "welcome", func(c *context.AppContext) error {
		now := time.Now()

		return c.AWHtml("welcome", context.TplData{
			"time": now.Format("2006-01-02"),
		})
	})
}
