package wrapper

import (
  "goweibo/app/auth"
  "goweibo/app/models/user"
  "goweibo/core/context"
)

func User(handler func(*context.AppContext, *auth.TokenAuthInfo, *user.User) error) context.AppHandlerFunc {
  return TokenAuth(func(c *context.AppContext, a *auth.TokenAuthInfo) error {
    id, err := c.AWIntParam("")
    if err != nil {
      return err
    }

    u, err := user.Get(uint(id))
    if err != nil {
      return err
    }

    return handler(c, a, u)
  })
}
