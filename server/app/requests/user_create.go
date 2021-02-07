package requests

import (
  "goweibo/app/models"
  "goweibo/core/errno"
  "goweibo/core/pkg/validator"
)

type CreateUser struct {
  Name                 string `valid:"name"`
  Email                string `valid:"email"`
  Password             string `valid:"password"`
}

func (u *CreateUser) Options() validator.Options {
  return validator.Options{
    Rules: validator.MapData{
      "name":     {"required", "max:50"},
      "email":    {"required", "max:255", "email", "not_exists:users"},
      "password": {"required", "min:6"},
    },
    Messages: validator.MapData{
      "name": {
        "required:名称不能为空",
        "max:名称长度不能大于 50 个字符",
      },
      "email": {
        "required:邮箱不能为空",
        "max:邮箱长度不能大于 255 个字符",
        "email:邮箱格式错误",
        "not_exists:邮箱已经被注册过了",
      },
      "password": {
        "required:密码不能为空",
        "min:密码长度不能小于 6 个字符",
      },
    },
  }
}

func (u *CreateUser) Create() (user *models.User, err error) {
  user = &models.User{
    Name: u.Name,
    Email: u.Email,
    Password: u.Password,
    Avatar: "",
  }
  if err = user.Create(); err != nil {
    err = errno.DatabaseErr.WithErr(err)
  }
  return
}
