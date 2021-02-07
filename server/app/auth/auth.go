package auth

import (
  "goweibo/app/models"
)

type TokenAuthInfo struct {
  User  *models.User
  Token string
}
