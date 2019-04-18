package models

import (
	"gin_weibo/database"
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Name            string    `gorm:"column:name;type:varchar(255);not null"`
	Email           string    `gorm:"column:email;type:varchar(255);unique;not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	Password        string    `gorm:"column:password;type:varchar(255);not null"`
	RememberToken   string    `gorm:"column:remember_token;type:varchar(100)"`
	IsAdmin         uint      `gorm:"column:is_admin;type:tinyint(1)"`
	ActivationTOken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// Get 获取一个用户
func (User) Get(id int) (*User, error) {
	u := &User{}
	d := database.DB.First(&u, id)
	return u, d.Error
}
