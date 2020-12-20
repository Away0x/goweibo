package bootstrap

import (
	"fmt"
	"goweibo/core"
	"goweibo/routes"

	"github.com/labstack/echo/v4"
)

// SetupServer 初始化 server
func SetupServer() {
	e := echo.New()
	e.Debug = core.GetConfig().IsDev()
	e.HideBanner = true

	core.NewApplication(e)
	// 注册路由
	routes.Register(e)
	// 输出路由配置
	core.GetApplication().PrintRoutes(core.GetConfig().String("APP.TEMP_DIR") + "/routes.json")

	fmt.Printf("\n\napp runmode is %s\n\n", core.GetConfig().AppRunMode())
	// 启动 server
	e.Logger.Fatal(core.GetApplicationEngine().Start(core.GetConfig().String("APP.PORT")))
}
