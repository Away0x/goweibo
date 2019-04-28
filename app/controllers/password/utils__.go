package password

import (
	"gin_weibo/routes/named"
	"github.com/gin-gonic/gin"
	"gin_weibo/app/helpers"
	passwordResetModel "gin_weibo/app/models/password_reset"
)

func sendResetEmail(pwd *passwordResetModel.PasswordReset) error {
	subject := "重置密码！请确认你的邮箱。"
	tpl := "mail/reset_password.html"
	resetPasswordURL := named.G("password.reset", "token", pwd.Token)

	return helpers.SendMail([]string{pwd.Email}, subject, tpl, gin.H{"resetPasswordURL": resetPasswordURL})
}
