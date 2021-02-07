package auth

import "goweibo/app/models/user"

type TokenAuthInfo struct {
  User  *user.User
  Token string
}
