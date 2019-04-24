package password

import (
	"gin_weibo/app/models"
	"gin_weibo/app/requests"
)

// PasswordEmailForm -
type PasswordEmailForm struct {
	Email string
}

func (u *PasswordEmailForm) emailExistValidator() requests.ValidatorFunc {
	return func() (msg string) {
		m := &models.User{}
		if err := m.GetByEmail(u.Email); err == nil {
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
func (u *PasswordEmailForm) ValidateAndGetToken() (pwd *models.PasswordReset, errors []string) {
	errors = u.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	pwd = &models.PasswordReset{
		Email: u.Email,
	}

	if _, err := pwd.Create(); err != nil {
		errors = append(errors, "失败: "+err.Error())
		return nil, errors
	}

	return pwd, []string{}
}

// -----------------------------------------------

// PassWordResetForm -
type PassWordResetForm struct {
	Email                string
	Token                string
	Password             string
	PasswordConfirmation string
}

func (u *PassWordResetForm) tokenExistValidator() requests.ValidatorFunc {
	return func() (msg string) {
		m := &models.PasswordReset{}
		if err := m.GetByToken(u.Token); err == nil {
			u.Email = m.Email
			return ""
		}
		return "该 token 不存在"
	}
}

// Validate : 验证函数
func (u *PassWordResetForm) Validate() (errors []string) {
	errors = requests.RunValidators(
		requests.ValidatorMap{
			"password": {
				requests.RequiredValidator(u.Password),
				requests.MixLengthValidator(u.Password, 6),
				requests.EqualValidator(u.Password, u.PasswordConfirmation),
			},
			"token": {
				requests.RequiredValidator(u.Token),
				u.tokenExistValidator(),
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
func (u *PassWordResetForm) ValidateAndUpdateUser() (user *models.User, errors []string) {
	errors = u.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	// 验证成功，删除 token
	pwd := &models.PasswordReset{}
	if err := pwd.DeleteByToken(u.Token); err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}

	// 更新用户密码
	user = &models.User{}
	if err := user.GetByEmail(u.Email); err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}
	user.Password = u.Password
	if err := user.Update(true); err != nil {
		errors = append(errors, "重置密码失败: "+err.Error())
		return nil, errors
	}

	return user, []string{}
}
