package requests

import (
	"gin_weibo/app/models"
)

// 以后可以改为 tag 来调用验证器函数
type UserForm struct {
	Name                 string
	Email                string
	Password             string
	PasswordConfirmation string
}

func (u *UserForm) emailUniqueValidator() validatorFunc {
	return func() (msg string) {
		m := &models.User{}
		if _, err := m.GetByEmail(u.Email); err != nil {
			return ""
		}
		return "邮箱已经被注册过了"
	}
}

// Validate : 验证函数
func (u *UserForm) Validate() (errors []string) {
	errors = RunValidators(
		validatorMap{
			"name": {
				RequiredValidator(u.Name),
				MaxLengthValidator(u.Name, 50),
			},
			"email": {
				RequiredValidator(u.Email),
				MaxLengthValidator(u.Email, 255),
				u.emailUniqueValidator(),
			},
			"password": {
				RequiredValidator(u.Password),
				MixLengthValidator(u.Password, 6),
				EqualValidator(u.Password, u.PasswordConfirmation),
			},
		},
		validatorMsgArr{
			"name": {
				"名称不能为空",
				"名称长度不能大于 50 个字符",
			},
			"email": {
				"邮箱不能为空",
				"邮箱长度不能大于 255 个字符",
				"邮箱已经被注册过了",
			},
			"password": {
				"密码不能为空",
				"密码长度不能小于 6 个字符",
				"两次输入的密码不一致",
			},
		},
	)

	return errors
}

func (u *UserForm) ValidateAndSave() (user *models.User, errors []string) {
	errors = u.Validate()

	if len(errors) != 0 {
		return nil, errors
	}

	// 创建用户
	user = &models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}

	if err := user.Encrypt(); err != nil {
		errors = append(errors, "用户创建失败: "+err.Error())
		return nil, errors
	}

	if err := user.Create(); err != nil {
		errors = append(errors, "用户创建失败: "+err.Error())
		return nil, errors
	}

	return user, []string{}
}
