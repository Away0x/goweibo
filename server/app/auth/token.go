package auth

import (
  "github.com/labstack/echo/v4"
  "goweibo/app/models"
  "goweibo/core/context"
  "goweibo/core/errno"
  "goweibo/core/pkg/jwttoken"
)

const (
  contextUserKeyName = "__user_key__"
)

func getUserFromContext(c echo.Context) (*models.User, bool) {
  u := c.Get(contextUserKeyName)
  if u == nil {
    return nil, false
  }

  if uu, ok := u.(*models.User); ok {
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

func GetUser(c *context.AppContext) (u *models.User, err error) {
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

  u, err = models.GetUserByID(claims.UserID)
  if err != nil {
    return nil, errno.DatabaseErr.WithErr(err)
  }

  c.Set(contextUserKeyName, u)
  return
}

func GetTokenAndUser(c *context.AppContext) (t string, u *models.User, err error) {
  if t, err = GetToken(c); err != nil {
    return
  }

  if u, err = GetUser(c); err != nil {
    return
  }

  return
}

func MustGetToken(c *context.AppContext) string {
  t, err := GetToken(c)
  if err != nil {
    panic(err)
  }
  return t
}

func MustGetUser(c *context.AppContext) *models.User {
  u, err := GetUser(c)
  if err != nil {
    panic(err)
  }
  return u
}
