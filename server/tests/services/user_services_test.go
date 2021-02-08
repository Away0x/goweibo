package services

import (
  "github.com/stretchr/testify/assert"
  "goweibo/app/services"
  "goweibo/tests"
  "testing"
)

func TestUserServicesList(t *testing.T) {
  a := assert.New(t)

  u1 := tests.CreateUserModel(t)
  u2 := tests.CreateUserModel(t)

  s := services.NewUserServices()
  us, err := s.List()
  a.NoError(err)

  a.Equal(len(us), 2)
  a.Equal(us[0].ID, u1.ID)
  a.Equal(us[1].ID, u2.ID)
}
