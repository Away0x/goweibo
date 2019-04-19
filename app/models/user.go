package models

import (
	"crypto/md5"
	"encoding/hex"
	"gin_weibo/database"
	"strconv"
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

// Gravatar 生成用户头像
func (u *User) Gravatar(size int) string {
	hash := md5.Sum([]byte(u.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:]) + "?s=" + strconv.Itoa(size)
}
