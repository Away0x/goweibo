package routes

import (
  "goweibo/app/controllers/api"
  "goweibo/app/services"
  "goweibo/core"
  "goweibo/core/context"
  "goweibo/core/pkg/captcha"
  _ "goweibo/docs"
  "goweibo/routes/wrapper"

  "github.com/labstack/echo/v4/middleware"
  echoSwagger "github.com/swaggo/echo-swagger"
  appMiddlewares "goweibo/routes/middlewares"
)

const (
  // APIPrefix api prefix
  APIPrefix = "/api"
  CaptchaURLPrefix = "/captcha"
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
    router.GET("/api-doc/*", echoSwagger.WrapHandler).Name = "api-doc"
  }

  // 验证码
  router.RegisterHandler(router.GET, CaptchaURLPrefix + "/:id", func(c *context.AppContext) error {
    return captcha.Handler(c)
  }).Name = "captcha"

  e := router.Group(APIPrefix, middleware.CORS())

  // 验证码
  captchaCon := e.Group(CaptchaURLPrefix)
  {
    cc := api.NewCaptchaController(CaptchaURLPrefix)
    router.RegisterHandler(captchaCon.GET, "", cc.New).Name = "captcha.new"
    router.RegisterHandler(captchaCon.POST, "", cc.Verify).Name = "captcha.verify"
  }

  auth := e.Group("/token")
  {
    tc := api.NewTokenController()
    router.RegisterHandler(auth.POST, "/store", tc.Store).Name = "token.store"
    router.RegisterHandler(auth.PUT, "/refresh", wrapper.TokenAuth(tc.Refresh)).Name = "token.refresh"
  }

  user := e.Group("/user", appMiddlewares.AuthToken)
  {
    uc := api.NewUserController(services.NewUserServices())
    router.RegisterHandler(user.GET, "", wrapper.TokenAuth(uc.Index)).Name = "user.index"
    router.RegisterHandler(user.GET, ":id", wrapper.User(uc.Show)).Name = "user.show"
    router.RegisterHandler(user.POST, "", uc.Create).Name = "user.create"
  }
}
