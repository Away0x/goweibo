package routes

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers/sessions"
	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/user"
	"gin_weibo/middleware/wrapper"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{
		g.GET("/", staticpage.Home)
		g.GET("/help", staticpage.Help)
		g.GET("/about", staticpage.About)
	}

	// ------------------------------ user ------------------------------
	{
		g.GET("/signup", wrapper.Guest(user.Create))
		g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
		userRouter := g.Group("/users")
		{
			// 创建用户页面
			userRouter.GET("/create", wrapper.Guest(user.Create))
			// 保存新用户
			userRouter.POST("", wrapper.Guest(user.Store))

			// 用户列表页面
			userRouter.GET("", wrapper.Auth(user.Index))
			// 展示具体用户页面
			userRouter.GET("/show/:id", wrapper.Auth(user.Show))

			// 编辑用户页面
			userRouter.GET("/edit/:id", wrapper.Auth(user.Edit))
			// 修改用户
			userRouter.POST("/update/:id", wrapper.Auth(user.Update))

			// 删除用户
			userRouter.POST("/destory/:id", wrapper.Auth(user.Destory))
		}
	}

	// ------------------------------ sessions ------------------------------
	{
		// 登录页面
		g.GET("/login", wrapper.Guest(sessions.Create))
		// 登录
		g.POST("/login", wrapper.Guest(sessions.Store))
		// 登出
		g.POST("/logout", sessions.Destroy)
	}
}
