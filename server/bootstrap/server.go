package bootstrap

import (
	"fmt"
	"goweibo/core"
	"goweibo/core/pkg/tpl"
	"goweibo/routes"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo/v4"
)

// SetupServer 初始化 server
func SetupServer() {
	// 初始化日志
	core.SetupLog()

	e := echo.New()
	e.Debug = core.GetConfig().IsDev()
	e.HideBanner = true

	core.NewApplication(e)
	// 注册路由
	routes.Register(e)
	// 输出路由配置
	core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")
	// init render
	SetupServerRender(e)

	fmt.Printf("\n\napp runmode is %s\n\n", core.GetConfig().AppRunMode())
	// 启动 server
	e.Logger.Fatal(core.GetApplicationEngine().Start(core.GetConfig().String("APP.PORT")))
}

// SetupServerRender 初始化 echo 渲染器
func SetupServerRender(e *echo.Echo) {
	render := tpl.NewRenderer()
	tpl.SetupTpl(&tpl.Config{
		GetRoutePath: core.GetApplication().RoutePath,
	})

	// template dir
	render.AddDirectory(core.GetConfig().String("APP.TEMPLATE_DIR"))

	// template global var
	globalVar := pongo2.Context{
		"APP_NAME":    core.GetConfig().String("APP.NAME"),
		"APP_RUNMODE": string(core.GetConfig().AppRunMode()),
		"APP_URL":     core.GetConfig().String("APP.URL"),
		"route":       core.GetApplication().RoutePath,
	}

	render.UseContextProcessor(func(echoCtx echo.Context, pongoCtx pongo2.Context) {
		pongoCtx.Update(globalVar)

		tpldata := pongo2.Context{}

		pongoCtx.Update(tpldata)
	})
	e.Renderer = render

	// tags
	pongo2.RegisterTag("route", tpl.RouteTag)
}
