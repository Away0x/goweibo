package auth

import (
  "github.com/labstack/echo/v4"
  userModel "goweibo/app/models/user"
  "goweibo/core/context"
  "goweibo/core/errno"
  "goweibo/core/pkg/jwttoken"
)

const (
  contextUserKeyName = "__user_key__"
)

func getUserFromContext(c echo.Context) (*userModel.User, bool) {
  u := c.Get(contextUserKeyName)
  if u == nil {
    return nil, false
  }

  if uu, ok := u.(*userModel.User); ok {
    return uu, true
  }

  return nil, false
}

func GetToken(c *context.AppContext) (t string, err error) {
  t, err = jwttoken.GetToken(c.Context)
  if err != nil {
    return "", errno.TokenErr.WithErr(err)
  }
  return
}

func GetUser(c *context.AppContext) (u *userModel.User, err error) {
  if u, ok := getUserFromContext(c.Context); ok {
    return u, nil
  }

  t, err := GetToken(c)
  if err != nil {
    return nil, err
  }

  claims, err := jwttoken.VerifyToken(t)
  if err != nil {
    return nil, errno.TokenErr.WithErr(err)
  }

  u, err = userModel.Get(claims.UserID)
  if err != nil {
    return nil, errno.DatabaseErr.WithErr(err)
  }

  c.Set(contextUserKeyName, u)
  return
}

func MustGetToken(c *context.AppContext) string {
  t, err := GetToken(c)
  if err != nil {
    panic(err)
  }
  return t
}

func MustGetUser(c *context.AppContext) *userModel.User {
  u, err := GetUser(c)
  if err != nil {
    panic(err)
  }
  return u
}
