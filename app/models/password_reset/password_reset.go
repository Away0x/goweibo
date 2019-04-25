package passwordreset

import (
	"time"
)

// PasswordReset 重置密码模型
type PasswordReset struct {
	Email     string    `gorm:"column:email;type:varchar(255);not null" sql:"index"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" sql:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// TableName 表名
func (PasswordReset) TableName() string {
	return "password_resets"
}