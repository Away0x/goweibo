package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/helpers"
	"gin_weibo/app/models"
	"gin_weibo/config"
	"gin_weibo/database"
	"gin_weibo/routes"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// gin config
	g := gin.New()
	setupGin(g)

	// db config
	db := database.InitDB()
	db.AutoMigrate(
		&models.User{},
	)
	defer db.Close()

	// router config
	routes.Register(g)

	// 启动
	log.Printf("Start to listening the incoming requests on http address: %s", config.AppConfig.Addr)
	log.Fatal(http.ListenAndServe(config.AppConfig.Addr, g).Error())
}

// 配置 gin
func setupGin(g *gin.Engine) {
	// 启动模式配置
	gin.SetMode(config.AppConfig.RunMode)

	// 项目静态文件配置
	g.Static("/"+config.ProjectConfig.PublicPath, config.ProjectConfig.PublicPath)
	g.StaticFile("/favicon.ico", config.ProjectConfig.PublicPath+"/favicon.ico")

	// 模板配置
	// 注册模板函数
	g.SetFuncMap(template.FuncMap{
		"Mix":    helpers.Mix,
		"Static": helpers.Static,
	})
	g.LoadHTMLGlob("resources/views/**/*")
}
