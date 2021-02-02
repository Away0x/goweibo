package user

import (
  "goweibo/app/models"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Avatar   string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
}
