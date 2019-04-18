package routes

import (
	"github.com/gin-gonic/gin"

	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/user"
)

func registerWeb(g *gin.Engine) {
	// static page
	g.GET("/", staticpage.Home)
	g.GET("/help", staticpage.Help)
	g.GET("/about", staticpage.About)

	// user
	g.GET("/signup", user.Create)
	userRouter := g.Group("/users")
	{
		// 用户列表
		userRouter.GET("", user.Index)
		// 创建用户页面
		userRouter.GET("/create", user.Create)
		// 展示具体用户页面
		userRouter.GET("/show/:id", user.Show)
		// 编辑用户页面
		userRouter.GET("/edit/:id", user.Edit)
		// 保存新用户
		userRouter.POST("", user.Store)
		// 修改用户
		userRouter.POST("/update/:id", user.Update)
		// 删除用户
		userRouter.POST("/destory/:id", user.Destroy)
	}
}
