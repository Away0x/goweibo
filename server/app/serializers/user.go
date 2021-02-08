package serializers

import (
  "goweibo/app/models"
  "goweibo/core/pkg/timeutils"
)

type (
  User struct {
    Name      string `json:"name"`
    Email     string `json:"email"`
    Avatar    string `json:"avatar"`
    UpdatedAt string `json:"updated_at"`
    IsAdmin   bool   `json:"is_admin"`
  }
  UserList []User
)

func NewUserSerializer(m *models.User) User {
  return User{
    Name:      m.Name,
    Email:     m.Email,
    Avatar:    m.Gravatar(),
    IsAdmin:   m.IsAdminRole(),
    UpdatedAt: timeutils.FormatDate(m.UpdatedAt),
  }
}

func NewUserListSerializer(ms []*models.User) UserList {
  us := make(UserList, 0)
  for _, m := range ms {
    us = append(us, NewUserSerializer(m))
  }
  return us
}
