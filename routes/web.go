package routes

import (
	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers/followers"
	"gin_weibo/app/controllers/password"
	"gin_weibo/app/controllers/sessions"
	staticpage "gin_weibo/app/controllers/static_page"
	"gin_weibo/app/controllers/status"
	"gin_weibo/app/controllers/user"
	"gin_weibo/middleware/wrapper"
	"gin_weibo/routes/named"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{
		g.GET("/", staticpage.Home)
		// 绑定路由 path 和 路由 name，之后可通过 named.G("root") 或 named.GR("root") 获取到路由 path
		// 模板文件中可通过 {{ Route "root" }} 或 {{ RelativeRoute "root" }} 获取 path
		named.Name(g, "root", "GET", "/")

		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "GET", "/help")

		g.GET("/about", staticpage.About)
		named.Name(g, "about", "GET", "/about")
	}

	// ------------------------------ user ------------------------------
	{
		g.GET("/signup", wrapper.Guest(user.Create))
		named.Name(g, "signup", "GET", "/signup")

		g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
		// 带参路由绑定，可通过 named.G("signup.confirm", token) 或 named.GR("signup.confirm", token) 获取 path
		// 模板文件中可通过 {{ Route "signup.confirm" .token }} 或 {{ RelativeRoute "signup.confirm" .token }} 获取 path
		named.Name(g, "signup.confirm", "GET", "/signup/confirm/:token") //

		userRouter := g.Group("/users")
		{
			// 创建用户页面
			userRouter.GET("/create", wrapper.Guest(user.Create))
			named.Name(userRouter, "users.create", "GET", "/create")
			// 保存新用户
			userRouter.POST("", wrapper.Guest(user.Store))
			named.Name(userRouter, "users.store", "POST", "")

			// 用户列表页面
			userRouter.GET("", wrapper.Auth(user.Index))
			named.Name(userRouter, "users.index", "GET", "")
			// 展示具体用户页面
			userRouter.GET("/show/:id", wrapper.Auth(user.Show))
			named.Name(userRouter, "users.show", "GET", "/show/:id")

			// 编辑用户页面
			userRouter.GET("/edit/:id", wrapper.Auth(user.Edit))
			named.Name(userRouter, "users.edit", "GET", "/edit/:id")
			// 修改用户
			userRouter.POST("/update/:id", wrapper.Auth(user.Update))
			named.Name(userRouter, "users.update", "POST", "/update/:id")

			// 删除用户
			userRouter.POST("/destroy/:id", wrapper.Auth(user.Destroy))
			named.Name(userRouter, "users.destroy", "POST", "/destroy/:id")

			// 用户关注者列表
			userRouter.GET("/followings/:id", wrapper.Auth(user.Followings))
			named.Name(userRouter, "users.followings", "GET", "/followings/:id")
			// 用户粉丝列表
			userRouter.GET("/followers/:id", wrapper.Auth(user.Followers))
			named.Name(userRouter, "users.followers", "GET", "/followers/:id")
			// 关注用户
			userRouter.POST("/followers/store/:id", wrapper.Auth(followers.Store))
			named.Name(userRouter, "followers.store", "POST", "/followers/store/:id")
			// 取消关注用户
			userRouter.POST("/followers/destroy/:id", wrapper.Auth(followers.Destroy))
			named.Name(userRouter, "followers.destroy", "POST", "/followers/destroy/:id")
		}
	}

	// ------------------------------ sessions ------------------------------
	{
		// 登录页面
		g.GET("/login", wrapper.Guest(sessions.Create))
		named.Name(g, "login.create", "GET", "/login")
		// 登录
		g.POST("/login", wrapper.Guest(sessions.Store))
		named.Name(g, "login.store", "POST", "/login")
		// 登出
		g.POST("/logout", sessions.Destroy)
		named.Name(g, "login.destroy", "POST", "/logout")
		named.Name(g, "logout", "POST", "/logout")
	}

	// ------------------------------ password ------------------------------
	passwordRouter := g.Group("/password")
	{
		// 显示重置密码的邮箱发送页面
		passwordRouter.GET("/reset", wrapper.Guest(password.ShowLinkRequestsForm))
		named.Name(passwordRouter, "password.request", "GET", "/reset")
		// 邮箱发送重设链接
		passwordRouter.POST("/email", wrapper.Guest(password.SendResetLinkEmail))
		named.Name(passwordRouter, "password.email", "POST", "/email")
		// 密码更新页面
		passwordRouter.GET("/reset/:token", wrapper.Guest(password.ShowResetForm))
		named.Name(passwordRouter, "password.reset", "GET", "/reset/:token")
		// 执行密码更新操作
		passwordRouter.POST("/reset", wrapper.Guest(password.Reset))
		named.Name(passwordRouter, "password.update", "POST", "/reset")
	}

	// ------------------------------ statuses ------------------------------
	statusRouter := g.Group("/statuses")
	{
		// 处理创建微博的请求
		statusRouter.POST("", wrapper.Auth(status.Store))
		named.Name(statusRouter, "statuses.store", "POST", "")
		// 处理删除微博的请求
		statusRouter.POST("/destroy/:id", wrapper.Auth(status.Destroy))
		named.Name(statusRouter, "statuses.destroy", "POST", "/destroy/:id")
	}

}
