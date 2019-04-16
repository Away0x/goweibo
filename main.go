package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_weibo/config"
	"gin_weibo/routes"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// gin config
	gin.SetMode(config.AppConfig.RunMode)
	g := gin.New()

	// router config
	routes.Register(g)

	// 启动
	log.Printf("Start to listening the incoming requests on http address: %s", config.AppConfig.Addr)
	log.Fatal(http.ListenAndServe(config.AppConfig.Addr, g).Error())
}
