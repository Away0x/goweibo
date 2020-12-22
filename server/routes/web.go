package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func registerWeb(e *echo.Echo) {
	ee := e.Group("/")

	ee.GET("welcome", func(c echo.Context) error {
		now := time.Now()

		return c.Render(http.StatusOK, "welcome.tpl", map[string]interface{}{
			"time": now.Format("2006-01-02"),
		})
	}).Name = "welcome"
}
