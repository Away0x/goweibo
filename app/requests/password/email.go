package password

import (
	passwordResetModel "gin_weibo/app/models/password_reset"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/app/requests"
)

// PasswordEmailForm -
type PasswordEmailForm struct {
	Email string
}

func (p *PasswordEmailForm) emailExistValidator() requests.ValidatorFunc {
	return func() (msg string) {
		if _, err := userModel.GetByEmail(p.Email); err == nil {
			return ""
		}
		return "该邮箱不存在"
	}
}

// Validate : 验证函数
func (u *PasswordEmailForm) Validate() (errors []string) {
	errors = requests.RunValidators(
		requests.ValidatorMap{
			"email": {
				requests.RequiredValidator(u.Email),
				requests.MaxLengthValidator(u.Email, 255),
				requests.EmailValidator(u.Email),
				u.emailExistValidator(),
			},
		},
		requests.ValidatorMsgArr{
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱格式错误",
				"该邮箱不存在",
			},
		},
	)

	return errors
}

// ValidateAndGetToken 验证参数并且创建验证 pwd 的 token
func (p *PasswordEmailForm) ValidateAndGetToken() (pwd *passwordResetModel.PasswordReset, errors []string) {
	errors = p.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	pwd = &passwordResetModel.PasswordReset{
		Email: p.Email,
	}

	if err := pwd.Create(); err != nil {
		errors = append(errors, "失败: "+err.Error())
		return nil, errors
	}

	return pwd, []string{}
}
