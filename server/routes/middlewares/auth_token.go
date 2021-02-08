package middlewares

import (
  "github.com/labstack/echo/v4"
  "goweibo/app/auth"
  "goweibo/core/context"
)

func AuthToken(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    _, _, err := auth.GetTokenAndUser(context.NewAppContext(c))
    if err != nil {
      return err
    }

    // 验证用户是否激活
    //if !u.IsActivated() {
    //  return errno.AuthErr.WithMessage("用户未激活")
    //}

    return next(c)
  }
}
