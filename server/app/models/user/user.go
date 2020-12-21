package user

import (
	"goweibo/app/models"
	"time"
)

// User 用户模型
type User struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Avatar   string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	// 是否为管理员
	IsAdmin uint `gorm:"type:tinyint(1)"`
	// 用户激活
	ActivationToken string    `gorm:"type:varchar(255)"`
	Activated       uint      `gorm:"ctype:tinyint(1);not null"`
	EmailVerifiedAt time.Time // 激活时间
	// 用于实现记住我功能，存入 cookie 中，下次带上时，即可直接登录
	RememberToken string `gorm:"type:varchar(100)"`
}
