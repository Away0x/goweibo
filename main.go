package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/lexkong/log"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/helpers"
	"gin_weibo/app/models"
	"gin_weibo/config"
	"gin_weibo/database"
	"gin_weibo/database/factory"
	"gin_weibo/routes"
	"gin_weibo/routes/named"

	"github.com/spf13/pflag"
)

var (
	// 需要 mock data，注意该操作会覆盖数据库；只在非 release 时生效
	needMock = pflag.BoolP("mock", "m", false, "need mock data")
)

func main() {
	// 解析命令行参数
	pflag.Parse()

	// 初始化配置
	config.InitConfig()

	// gin config
	g := gin.New()
	setupGin(g)

	// db config
	db := database.InitDB()
	// mock data
	if do := factoryMake(); do {
		return
	}
	// db migrate
	db.AutoMigrate(
		&models.User{},
		&models.PasswordReset{},
	)
	defer db.Close()

	// router config
	routes.Register(g)

	// 启动
	fmt.Printf("\n\n-------------------------------------------------- Start to listening the incoming requests on http address: %s --------------------------------------------------\n\n", config.AppConfig.Addr)
	if err := http.ListenAndServe(config.AppConfig.Addr, g); err != nil {
		log.Fatal("http server 启动失败", err)
	}
}

// 配置 gin
func setupGin(g *gin.Engine) {
	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.AppConfig.PublicPath, config.AppConfig.PublicPath)
	g.StaticFile("/favicon.ico", config.AppConfig.PublicPath+"/favicon.ico")

	// 模板配置
	// 注册模板函数
	g.SetFuncMap(template.FuncMap{
		"Mix":    helpers.Mix,
		"Static": helpers.Static,
		"Route":  named.G,
	})
	g.LoadHTMLGlob(config.AppConfig.ViewsPath + "/**/*")
}

// 数据 mock
func factoryMake() (do bool) {
	// 只有非 release 时才可用该函数
	if config.AppConfig.RunMode == config.RunmodeRelease {
		return false
	}
	status := *needMock

	if !status {
		return false
	}

	fmt.Print("\n\n-------------------------------------------------- MOCK --------------------------------------------------\n\n")
	factory.UsersTableSeeder(true)
	return true
}
