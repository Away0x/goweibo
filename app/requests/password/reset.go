package password

import (
	passwordResetModel "gin_weibo/app/models/password_reset"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/app/requests"
)

// PassWordResetForm -
type PassWordResetForm struct {
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}

func (p *PassWordResetForm) tokenExistValidator() requests.ValidatorFunc {
	return func() (msg string) {
		if m, err := passwordResetModel.GetByToken(p.Token); err == nil {
			p.Email = m.Email
			return ""
		}
		return "该 token 不存在"
	}
}

// Validate : 验证函数
func (p *PassWordResetForm) Validate() (errors []string) {
	errors = requests.RunValidators(
		requests.ValidatorMap{
			"password": {
				requests.RequiredValidator(p.Password),
				requests.MixLengthValidator(p.Password, 6),
				requests.EqualValidator(p.Password, p.PasswordConfirmation),
			},
			"token": {
				requests.RequiredValidator(p.Token),
				p.tokenExistValidator(),
			},
		},
		requests.ValidatorMsgArr{
			"password": {
				"密码不能为空",
				"密码长度不能小于 6 个字符",
				"两次输入的密码不一致",
			},
			"token": {
				"token 不能为空",
				"该 token 不存在",
			},
		},
	)

	return errors
}

// ValidateAndUpdateUser 验证参数并且创建验证 pwd 的 token
func (p *PassWordResetForm) ValidateAndUpdateUser() (user *userModel.User, errors []string) {
	errors = p.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	// 验证成功，删除 token
	if err := passwordResetModel.DeleteByToken(p.Token); err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}

	// 更新用户密码
	user, err := userModel.GetByEmail(p.Email)
	if err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}
	user.Password = p.Password
	if err = user.Update(true); err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}

	return user, []string{}
}
