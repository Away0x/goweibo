package requests

import (
  "github.com/thedevsaddam/govalidator"
  "goweibo/core/pkg/validator"
)

type UserLogin struct {
  validator.BaseValidator
  Email string `json:"email"`
  Password string `json:"password"`
}

func (u *UserLogin) Options() govalidator.Options {
  return govalidator.Options{
    Data: u,
    Rules: govalidator.MapData{
      "email":    []string{"required", "max:255", "email"},
      "password": []string{"required", "min:6"},
    },
    Messages: govalidator.MapData{
      "email": []string{
        "required:邮箱不能为空",
        "max:邮箱长度不能大于255个字符",
        "email:邮箱格式错误",
      },
      "password": []string{
        "required:密码不能为空",
        "min:密码不能少于6位",
      },
    },
  }
}



