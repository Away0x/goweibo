package models_test

import (
  "github.com/stretchr/testify/assert"
  "goweibo/app/models"
  "testing"
)

func TestCreateUser(t *testing.T) {
  u := &models.User{Name: "test1", Email: "1@qq.com", Password: "123"}
  err := models.CreateModel(u)
  assert.Equal(t, err, nil)

  _, err = models.GetUserByID(u.ID)
  assert.Equal(t, err, nil)
}
