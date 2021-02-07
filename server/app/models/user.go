package models

import (
  "gorm.io/gorm"
  "goweibo/core/pkg/password"
)

// User 用户模型
type User struct {
  BaseModel
  Name     string `gorm:"type:varchar(255);not null" json:"name"`
  Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
  Avatar   string `gorm:"type:varchar(255);not null" json:"avatar"`
  Password string `gorm:"type:varchar(255);not null" json:"-"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
  if !password.Encrypted(u.Password) {
    pwd, err := password.Encrypt(u.Password)
    if err == nil {
      u.Password = pwd
    }
  }

  return err
}

func GetUserByID(id uint) (user *User, err error) {
  user = new(User)
  d := DB().First(user, id)
  return user, d.Error
}

func GetUserByEmail(email string) (user *User, err error) {
  user = new(User)
  d := DB().Where("email = ?", email).First(user)
  return user, d.Error
}

func (u *User) Create() (err error) {
  err = DB().Create(&u).Error
  return
}
