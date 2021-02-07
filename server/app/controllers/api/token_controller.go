package api

import (
  "goweibo/app/auth"
  userModel "goweibo/app/models/user"
  "goweibo/app/requests"
  "goweibo/core/context"
  "goweibo/core/errno"
)

type ITokenController interface {
  Store(*context.AppContext) error
  Refresh(*context.AppContext, *auth.TokenAuthInfo) error
}

type TokenController struct {}

func NewTokenController() ITokenController {
  return &TokenController{}
}

// Store create token
// @Summary create token
// @Tags Token
// @Accept json
// @Produce json
// @Param json body requests.UserLogin true "登录信息"
// @Success 200 {object} context.CommonResponse{data=context.TokenResp}
// @Router /token/store [post]
func (*TokenController) Store(c *context.AppContext) (err error) {
  req := new(requests.UserLogin)
  if err = c.AWBindValidatorStruct(req); err != nil {
    return err
  }

  u, err := userModel.GetByEmail(req.Email)
  if err != nil {
    return errno.DatabaseErr.WithErr(err)
  }

  result, err := c.AWTokenSign(u.ID)
  if err != nil {
    return err
  }

  return c.AWSuccessJSON(result)
}

// Refresh 刷新 token
// @Summary refresh token
// @Tags Token
// @Accept json
// @Produce json
// @Success 200 {object} context.CommonResponse{data=context.TokenResp}
// @Security ApiKeyAuth
// @Router /token/refresh [put]
func (*TokenController) Refresh(c *context.AppContext, a *auth.TokenAuthInfo) error {
  r, err := c.AWTokenRefresh(a.Token)
  if err != nil {
    return err
  }

  return c.AWSuccessJSON(r)
}
