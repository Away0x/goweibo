package wrapper

import (
  "goweibo/app/auth"
  "goweibo/app/models"
  "goweibo/core/context"
)

func User(handler func(*context.AppContext, *auth.TokenAuthInfo, *models.User) error) context.AppHandlerFunc {
  return TokenAuth(func(c *context.AppContext, a *auth.TokenAuthInfo) error {
    id, err := c.AWIntParam("")
    if err != nil {
      return err
    }

    if uint(id) == a.User.ID {
      return handler(c, a, a.User)
    }

    u, err := models.GetUserByID(uint(id))
    if err != nil {
      return err
    }

    return handler(c, a, u)
  })
}
