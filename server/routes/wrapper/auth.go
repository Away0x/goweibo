package wrapper

import (
  "goweibo/app/auth"
  "goweibo/core/context"
)

func TokenAuth(handler func(*context.AppContext, *auth.TokenAuthInfo) error) context.AppHandlerFunc {
  return func(c *context.AppContext) error {
    t, err := auth.GetToken(c)
    if err != nil {
      return err
    }

    u, err := auth.GetUser(c)
    if err != nil {
      return err
    }

    a := &auth.TokenAuthInfo{User: u, Token: t}
    return handler(c, a)
  }
}
