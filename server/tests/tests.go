package tests

import (
  "github.com/Pallinder/go-randomdata"
  "github.com/stretchr/testify/assert"
  "goweibo/app/models"
  "testing"
)

const (
  DefaultUserPassword = "123456"
)

func ResetDatabase() error {
  return ClearDatabase(
    &models.User{},
    &models.Status{},
  )
}

func ClearDatabase(tables ...interface{}) (err error) {
  for _, t := range tables {
    err = models.DB().Migrator().DropTable(t)
    err = models.DB().Migrator().CreateTable(t)
  }

  return
}

func CreateUserModel(t *testing.T) *models.User {
  u := &models.User{
    Name: randomdata.FullName(randomdata.Male),
    Email: randomdata.Email(),
    Password: DefaultUserPassword,
  }
  assert.NoError(t, models.CreateModel(u))
  return u
}
