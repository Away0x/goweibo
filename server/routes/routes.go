package routes

import (
  "goweibo/core"
  "goweibo/core/pkg/session"
  "goweibo/core/sdhandler"
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

  // 可以指定请求从 POST 重写为其他 (DELETE、PUT、PATCH ...)
  // form value 中需要携带 _method 参数
	router.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

  // 去除 url 尾部 /
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

  // 服务器健康自检
  sdhandler.RegisterSDHandlers(router.Echo, "/sd")

	// 注册路由
	registerError(router)
	registerWeb(router)
	registerAPI(router)
}
