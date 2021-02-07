package user

import (
  "goweibo/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Avatar   string `gorm:"type:varchar(255);not null" json:"avatar"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

func (u User) Serialize() interface{} {
  return u
}
