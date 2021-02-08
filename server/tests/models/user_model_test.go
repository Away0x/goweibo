package models_test

import (
  "github.com/stretchr/testify/assert"
  "goweibo/app/models"
  "goweibo/tests"
  "testing"
)

func TestCreateUserModel(t *testing.T) {
  a := assert.New(t)

  u := tests.CreateUserModel(t)
  _, err := models.GetUserByID(u.ID)
  a.NoError(err)
}
