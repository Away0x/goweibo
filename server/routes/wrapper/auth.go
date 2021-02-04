package wrapper

import (
  "goweibo/app/models/user"
  "goweibo/core/context"
  "goweibo/core/errno"
  "goweibo/core/pkg/jwttoken"
)

type AuthInfo struct {
  User *user.User
  Token string
}

func TokenAuth(handler func(*context.AppContext, *AuthInfo) error) context.AppHandlerFunc {
  return func(c *context.AppContext) error {
    t, err := jwttoken.GetToken(c.Context)
    if err != nil {
      return errno.TokenErr.WithErr(err)
    }

    claims, err := jwttoken.VerifyToken(t)
    if err != nil {
      return errno.TokenErr.WithErr(err)
    }

    u, err := user.GetUser(claims.UserID)
    if err != nil {
      return errno.DatabaseErr.WithErr(err)
    }

    a := &AuthInfo{
      User: u,
      Token: t,
    }

    return handler(c, a)
  }
}
