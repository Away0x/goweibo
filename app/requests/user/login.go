package user

import (
	userModel "gin_weibo/app/models/user"
	"gin_weibo/app/requests"

	"gin_weibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

type UserLoginForm struct {
	Email    string
	Password string
}

// Validate : 验证函数
func (u *UserLoginForm) Validate() (errors []string) {
	errors = requests.RunValidators(
		requests.ValidatorMap{
			"email": {
				requests.RequiredValidator(u.Email),
				requests.MaxLengthValidator(u.Email, 255),
				requests.EmailValidator(u.Email),
			},
			"password": {
				requests.RequiredValidator(u.Password),
			},
		},
		requests.ValidatorMsgArr{
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱格式错误",
			},
			"password": {
				"密码不能为空",
			},
		},
	)

	return errors
}

// ValidateAndLogin 验证参数并且获取用户
func (u *UserLoginForm) ValidateAndGetUser(c *gin.Context) (user *userModel.User, errors []string) {
	errors = u.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	// 通过邮箱获取用户，并且判断密码是否正确
	user, err := userModel.GetByEmail(u.Email)
	if err != nil {
		errors = append(errors, "该邮箱没有注册过用户: "+err.Error())
		return nil, errors
	}

	if err := user.Compare(u.Password); err != nil {
		flash.NewDangerFlash(c, "很抱歉，您的邮箱和密码不匹配")
		return nil, errors
	}

	return user, []string{}
}
