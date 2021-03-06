package bootstrap

import (
  "fmt"
  "goweibo/core"
  "goweibo/core/pkg/strutils"
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
  core.GetApplication().RegisterRoutes(routes.Register)
  // 输出路由配置
  core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")
  // init render
  SetupServerRender()

  fmt.Printf("\napp runmode is %s, %s\n\n", core.GetConfig().AppRunMode(), core.GetConfig().String("APP.URL"))
}

// RunServer 启动 server
func RunServer() {
  core.GetApplication().Echo.Logger.Fatal(core.GetApplication().Start(core.GetConfig().String("APP.ADDR")))
}

// SetupServerRender 初始化 echo 渲染器
func SetupServerRender() {
  render := tpl.NewRenderer()
  tpl.SetupTpl(&tpl.Config{
    GetRoutePath: core.GetApplication().RoutePath,
    GeneratePublicPath: func(path string) string {
      staticURL := core.GetConfig().String("APP.STATIC_URL")
      if core.GetConfig().IsDev() {
        return fmt.Sprintf("%s%s?v=%s", staticURL, path, strutils.RandomCreateBytes(6))
      }
      return fmt.Sprintf("%s%s", staticURL, path)
    },
  })

  // template dir
  render.AddDirectory(core.GetConfig().String("APP.TEMPLATE_DIR"))

  // template global var
  globalVar := pongo2.Context{
    "APP_NAME":    core.GetConfig().String("APP.NAME"),
    "APP_RUNMODE": string(core.GetConfig().AppRunMode()),
    "APP_URL":     core.GetConfig().String("APP.URL"),
  }

  render.UseContextProcessor(func(echoCtx echo.Context, pongoCtx pongo2.Context) {
    pongoCtx.Update(globalVar)

    tpldata := pongo2.Context{}

    pongoCtx.Update(tpldata)
  })

  core.GetApplication().Renderer = render

  // tags
  pongo2.RegisterTag("route", tpl.RouteTag)
  pongo2.RegisterTag("static", tpl.StaticTag)
}
