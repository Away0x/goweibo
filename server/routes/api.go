package routes

import (
	"goweibo/app/controllers/api"
	"goweibo/core"
	_ "goweibo/docs"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	// APIPrefix api prefix
	APIPrefix = "/api"
)

// @title goweibo Api
// @version 1.0
// @description goweibo api document

// @contact.name Away0x
// @contact.url https://github.com/Away0x
// @contact.email wutong0910@foxmail.com

// @host localhost:9999
// @BasePath /api

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func registerAPI(router *core.Application) {
	if core.GetConfig().IsDev() {
		router.GET("/apidoc/*", echoSwagger.WrapHandler).Name = "apidoc"
	}

	e := router.Group(APIPrefix, middleware.CORS())

	router.RegisterHandler(e.GET, "test", api.Test).Name = "api-test"
}
