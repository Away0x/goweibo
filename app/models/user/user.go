package user

import (
  "time"
  "gin_weibo/app/models"
)


// User 用户模型
type User struct {
	models.BaseModel
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Email    string `gorm:"column:email;type:varchar(255);unique;not null"`
	Avatar   string `gorm:"column:avatar;type:varchar(255);not null"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	// 是否为管理员
	IsAdmin uint `gorm:"column:is_admin;type:tinyint(1)"`
	// 用户激活
	ActivationToken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"` // 激活时间
	// 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录
	RememberToken string `gorm:"column:remember_token;type:varchar(100)"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}
