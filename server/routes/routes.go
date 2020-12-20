package routes

import (
	"goweibo/core"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	staticPath = "/public"
)

// Register 注册路由
func Register(e *echo.Echo) {
	e.Use(middleware.Recover())

	if core.GetConfig().IsDev() {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `${status}   ${method}   ${latency_human}               ${uri}`,
		}))
	}

	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Pre(middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		RedirectCode: http.StatusMovedPermanently,
	}))

	if core.GetConfig().Bool("APP.GZIP") {
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Skipper: func(c echo.Context) bool {
				// 只 gzip 静态文件
				return !strings.HasPrefix(c.Request().URL.Path, staticPath)
			},
		}))
	}

	// 静态文件路由
	e.Static(staticPath, core.GetConfig().String("APP.PUBLIC_DIR"))
	e.File("/favicon.ico", core.GetConfig().String("APP.PUBLIC_DIR")+"/favicon.ico")

	// 注册路由
	registerError(e)
	registerWeb(e)
	registerAPI(e)
}
