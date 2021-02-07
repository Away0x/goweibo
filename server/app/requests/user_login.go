package requests

import "goweibo/core/pkg/validator"

type UserLogin struct {
  Email    string `valid:"email"`
  Password string `valid:"password"`
}

func (u *UserLogin) Options() validator.Options {
  return validator.Options{
    Rules: validator.MapData{
      "email":    {"required", "max:255", "email"},
      "password": {"required", "min:6"},
    },
    Messages: validator.MapData{
      "email": {
        "required:邮箱不能为空",
        "max:邮箱长度不能大于255个字符",
        "email:邮箱格式错误",
      },
      "password": {
        "required:密码不能为空",
        "min:密码不能少于6位",
      },
    },
  }
}
