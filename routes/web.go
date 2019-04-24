package routes

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers/password"
	"gin_weibo/app/controllers/sessions"
	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/user"
	"gin_weibo/middleware/wrapper"
	"gin_weibo/routes/named"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{
		g.GET("/", staticpage.Home)
		named.Name(g, "root", "/")
		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "/help")
		g.GET("/about", staticpage.About)
		named.Name(g, "about", "/about")
	}

	// ------------------------------ user ------------------------------
	{
		g.GET("/signup", wrapper.Guest(user.Create))
		named.Name(g, "signup", "/signup")
		g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
		named.Name(g, "signup.confirm", "/signup/confirm/:token")

		userRouter := g.Group("/users")
		{
			// 创建用户页面
			userRouter.GET("/create", wrapper.Guest(user.Create))
			named.Name(userRouter, "users.create", "/create")
			// 保存新用户
			userRouter.POST("", wrapper.Guest(user.Store))
			named.Name(userRouter, "users.store", "")

			// 用户列表页面
			userRouter.GET("", wrapper.Auth(user.Index))
			named.Name(userRouter, "users.index", "")
			// 展示具体用户页面
			userRouter.GET("/show/:id", wrapper.Auth(user.Show))
			named.Name(userRouter, "users.show", "/show/:id")

			// 编辑用户页面
			userRouter.GET("/edit/:id", wrapper.Auth(user.Edit))
			named.Name(userRouter, "users.edit", "/edit/:id")
			// 修改用户
			userRouter.POST("/update/:id", wrapper.Auth(user.Update))
			named.Name(userRouter, "users.update", "/update/:id")

			// 删除用户
			userRouter.POST("/destory/:id", wrapper.Auth(user.Destory))
			named.Name(userRouter, "users.destory", "/destory/:id")
		}
	}

	// ------------------------------ sessions ------------------------------
	{
		// 登录页面
		g.GET("/login", wrapper.Guest(sessions.Create))
		named.Name(g, "login.create", "/login")
		// 登录
		g.POST("/login", wrapper.Guest(sessions.Store))
		named.Name(g, "login.store", "/login")
		// 登出
		g.POST("/logout", sessions.Destroy)
		named.Name(g, "login.destory", "/logout")
		named.Name(g, "logout", "/logout")
	}

	// ------------------------------ password ------------------------------
	passwordRouter := g.Group("/password")
	{
		// 显示重置密码的邮箱发送页面
		passwordRouter.GET("/reset", wrapper.Guest(password.ShowLinkRequestsForm))
		named.Name(passwordRouter, "password.request", "/reset")
		// 邮箱发送重设链接
		passwordRouter.POST("/email", wrapper.Guest(password.SendResetLinkEmail))
		named.Name(passwordRouter, "password.email", "/email")
		// 密码更新页面
		passwordRouter.GET("/reset/:token", wrapper.Guest(password.ShowResetForm))
		named.Name(passwordRouter, "password.reset", "/reset/:token")
		// 执行密码更新操作
		passwordRouter.POST("/reset", wrapper.Guest(password.Reset))
		named.Name(passwordRouter, "password.update", "/reset")
	}
}
