package models

import (
  "github.com/Pallinder/go-randomdata"
  "github.com/stretchr/testify/assert"
  "goweibo/app/models"
  "goweibo/tests"
  "testing"
)

func TestCreateStatusModel(t *testing.T) {
  a := assert.New(t)
  u := tests.CreateUserModel(t)

  s1 := &models.Status{UserID: u.ID, Content: randomdata.RandStringRunes(6)}
  a.NoError(models.CreateModel(s1))
  s2 := &models.Status{UserID: u.ID, Content: randomdata.RandStringRunes(6)}
  a.NoError(models.CreateModel(s2))

  ss := make([]*models.Status, 0)
  models.DB().Where("user_id = ?", u.ID).Find(&ss)
  a.Equal(len(ss), 2)
  a.Equal(ss[0].UserID, u.ID)

  u = new(models.User)
  a.NoError(models.DB().Preload("Statuses").First(u).Error)
  a.Equal(len(u.Statuses), 2)
  a.Equal(u.Statuses[0].UserID, u.ID)
  a.Equal(u.Statuses[0].ID, s1.ID)
}
