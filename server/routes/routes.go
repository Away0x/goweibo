package routes

import (
	"goweibo/core"
	"goweibo/core/pkg/session"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	staticPath = "/public"
)

// Register 注册路由
func Register(router *core.Application) {
  if !core.GetConfig().IsDev() {
    router.Use(middleware.Recover())
  }

	if core.GetConfig().IsDev() {
		router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${status}   ${method}   ${latency_human}               ${uri}\n",
		}))
	}

	router.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	router.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	if core.GetConfig().Bool("APP.GZIP") {
		router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				// 只 gzip 静态文件
				return !strings.HasPrefix(c.Request().URL.Path, staticPath)
			},
		}))
	}

	// session
	router.Use(session.NewMiddleware(session.MiddlewareOptions{
		HttpOnly:    true,
		Path:        "/",
		MaxAge:      86400 * 30,
		SessionName: core.GetConfig().String("APP.NAME"),
		SessionKey:  core.GetConfig().String("APP.KEY"),
	}))

	// 静态文件路由
	router.Static(staticPath, core.GetConfig().String("APP.PUBLIC_DIR"))
	router.File("/favicon.ico", core.GetConfig().String("APP.PUBLIC_DIR")+"/favicon.ico")

	// 注册路由
	registerError(router)
	registerWeb(router)
	registerAPI(router)
}
