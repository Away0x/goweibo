package api

import (
  "goweibo/app/auth"
  "goweibo/app/models"
  "goweibo/app/models/user"
  "goweibo/core/context"
  "goweibo/core/errno"
)

type IUserController interface {
  Index(*context.AppContext, *auth.TokenAuthInfo) error
  Show(*context.AppContext, *auth.TokenAuthInfo, *user.User) error
}

type UserController struct {}

func NewUserController() IUserController {
  return &UserController{}
}

func (*UserController) Index(c *context.AppContext, a *auth.TokenAuthInfo) error {
  users := new([]*user.User)
  if err := models.DB().Find(users).Error; err != nil {
    return errno.DatabaseErr.WithErr(err)
  }
  return c.AWSuccessJSON(users)
}

// Show 获取用户信息
// @Summary 获取用户信息
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} context.CommonResponse{data=user.User}
// @Security ApiKeyAuth
// @Router /user/show [get]
func (*UserController) Show(c *context.AppContext, a *auth.TokenAuthInfo, u *user.User) error {
  return c.AWSuccessJSON(u)
}
