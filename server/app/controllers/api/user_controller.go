package api

import (
  "goweibo/app/auth"
  "goweibo/app/models/user"
  "goweibo/app/services"
  "goweibo/core/context"
  "goweibo/core/errno"
)

type IUserController interface {
  Index(*context.AppContext, *auth.TokenAuthInfo) error
  Show(*context.AppContext, *auth.TokenAuthInfo, *user.User) error
}

type UserController struct {
  UserServices services.IUserServices
}

func NewUserController(s services.IUserServices) IUserController {
  return &UserController{UserServices: s}
}

// Index 获取用户列表
// @Summary 获取用户列表
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} context.CommonResponse{data=[]user.User}
// @Security ApiKeyAuth
// @Router /user [get]
func (u *UserController) Index(c *context.AppContext, a *auth.TokenAuthInfo) error {
  users, err := u.UserServices.List()
  if err != nil {
    return errno.DatabaseErr.WithErr(err)
  }
  return c.AWSuccessJSON(users)
}

// Show 获取用户信息
// @Summary 获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "用户 id"
// @Success 200 {object} context.CommonResponse{data=user.User}
// @Security ApiKeyAuth
// @Router /user/{id} [get]
func (*UserController) Show(c *context.AppContext, a *auth.TokenAuthInfo, u *user.User) error {
  return c.AWSuccessJSON(u.Serialize())
}
