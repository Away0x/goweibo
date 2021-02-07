package services

import (
  "goweibo/app/models"
)

type IUserServices interface {
  List() ([]*models.User, error)
}

type UserServices struct {}

func NewUserServices() *UserServices {
  return &UserServices{}
}

func (*UserServices) List() (users []*models.User, err error) {
  users = make([]*models.User, 0)
  err = models.DB().Find(&users).Error
  return
}
