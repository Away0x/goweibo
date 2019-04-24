// 只有非登录用户才可访问
package wrapper

import (
	"gin_weibo/app/auth"
	"gin_weibo/app/controllers"
	"gin_weibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

// Guest : 非登录用户才可访问
func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户已经登录了则跳转到 root page
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser != nil || err == nil {
			flash.NewInfoFlash(c, "您已登录，无需再次操作。")
			controllers.RedirectRouter(c, "root")
			return
		}

		handler(c)
	}
}
