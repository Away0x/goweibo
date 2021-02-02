package jwttoken

import (
  "fmt"
  "github.com/labstack/echo/v4"
)

const (
  tokenParamsKeyName          = "token"         // token 存在于 query formValue 中的 key name
  tokenHeaderKeyName          = "Authorization" // header key
  tokenInHeaderIdentification = "Bearer"        // header value split

  contextTokenKeyName = "_jwt_token_key"
)

func GetToken(c echo.Context) (string, error) {
  if token, ok := getTokenFromContext(c); ok {
    return token, nil
  }

  if token, ok := getTokenFromHeader(c); ok {
    c.Set(contextTokenKeyName, token)
    return token, nil
  }

  if token, ok := getTokenFromParams(c); ok {
    c.Set(contextTokenKeyName, token)
    return token, nil
  }

  return "", tokenNotFoundErr
}

func getTokenFromHeader(c echo.Context) (string, bool) {
  header := c.Request().Header.Get(tokenHeaderKeyName)
  if header == "" {
    return "", false
  }

  var token string
  _, err := fmt.Sscanf(header, tokenInHeaderIdentification+" %s", &token)
  if token == "" || err != nil {
    return "", false
  }
  return token, true
}

func getTokenFromParams(c echo.Context) (string, bool) {
  token := c.QueryParam(tokenParamsKeyName)
  if token != "" {
    return token, true
  }

  token = c.FormValue(tokenParamsKeyName)
  if token != "" {
    return token, true
  }

  return "", false
}

func getTokenFromContext(c echo.Context) (string, bool) {
  t := c.Get(contextTokenKeyName)
  if t == nil {
    return "", false
  }

  if s, ok := t.(string); ok {
    return s, true
  }

  return "", false
}
