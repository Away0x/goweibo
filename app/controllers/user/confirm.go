package user

import (
	"gin_weibo/app/auth"
	"gin_weibo/app/controllers"
	"gin_weibo/app/models"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/pkg/flash"
	"time"

	"github.com/gin-gonic/gin"
)

// ConfirmEmail : 邮箱验证、用户激活
func ConfirmEmail(c *gin.Context) {
	token := c.Param("token")

	user, err := userModel.GetByActivationToken(token)
	if user == nil || err != nil {
		controllers.Render404(c)
		return
	}

	// 更新用户
	user.Activated = models.TrueTinyint
	user.ActivationToken = ""
	user.EmailVerifiedAt = time.Now()
	if err = user.Update(false); err != nil {
		flash.NewSuccessFlash(c, "用户激活失败: "+err.Error())
		controllers.RedirectRouter(c, "root")
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "恭喜你，激活成功！")
	controllers.RedirectRouter(c, "users.show", user.ID)
}
