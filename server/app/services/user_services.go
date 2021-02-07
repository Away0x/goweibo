package services

import (
  "goweibo/app/models"
  "goweibo/app/models/user"
)

type IUserServices interface {
  List() ([]*user.User, error)
}

type UserServices struct {}

func NewUserServices() *UserServices {
  return &UserServices{}
}

func (*UserServices) List() (users []*user.User, err error) {
  users = make([]*user.User, 0)
  err = models.DB().Find(&users).Error
  return
}
