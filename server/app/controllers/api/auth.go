package api

import (
  "goweibo/app/models/user"
  "goweibo/core/context"
  "goweibo/routes/wrapper"
  "strconv"
)

func AuthLogin(c *context.AppContext) error {
  id := c.QueryParam("id")
  iid, _ := strconv.Atoi(id)

  u, err := user.GetUser(uint(iid))
  if err != nil {
    return err
  }

  result, err := c.AWTokenSign(u.ID)
  if err != nil {
    return err
  }

  return c.AWSuccessJSON(result)
}

func AuthRefreshToken(c *context.AppContext, a *wrapper.AuthInfo) error {
  r, err := c.AWTokenRefresh(a.Token)
  if err != nil {
    return err
  }

  return c.AWSuccessJSON(r)
}

func AuthGetUserInfo(c *context.AppContext, a *wrapper.AuthInfo) error {
  return c.AWSuccessJSON(a)
}
