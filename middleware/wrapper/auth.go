// 只有登录用户才可访问
package wrapper

import (
	"gin_weibo/app/auth"
	"gin_weibo/app/controllers"
	userModel "gin_weibo/app/models/user"

	"github.com/gin-gonic/gin"
)

type (
	AuthHandlerFunc = func(*gin.Context, *userModel.User)
)

// Auth : 登录用户才可访问
func Auth(handler AuthHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户未登录则跳转到登录页
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser == nil || err != nil {
			controllers.RedirectToLoginPage(c)
			return
		}

		handler(c, currentUser)
	}
}
