package context

import (
  "goweibo/core/errno"
  "goweibo/core/pkg/jwttoken"
)

type TokenResp struct {
  AccessToken  *jwttoken.AppJWTInfo `json:"access_token"`
  RefreshToken *jwttoken.AppJWTInfo `json:"refresh_token,omitempty"`
}

// AWTokenSign 签发 token
func (c *AppContext) AWTokenSign(userID uint) (*TokenResp, error) {
  a, r, err := jwttoken.CreateToken(userID)
  if err != nil {
    return nil, errno.TokenErr.WithErr(err)
  }

  return &TokenResp{AccessToken: a, RefreshToken: r}, nil
}

// AWTokenRefresh 刷新 token
func (c *AppContext) AWTokenRefresh(t string) (*TokenResp, error) {
  t, err := jwttoken.GetToken(c.Context)
  if err != nil {
    return nil, errno.TokenErr.WithErr(err)
  }

  td, err := jwttoken.RefreshToken(t)
  if err != nil {
    return nil, errno.TokenErr.WithErr(err)
  }

  return &TokenResp{AccessToken: td}, nil
}
