package models

import "time"

// PasswordReset 重置密码模型
type PasswordReset struct {
  Email     string    `gorm:"type:varchar(255);not null;index"`
  Token     string    `gorm:"type:varchar(255);not null;index"`
  CreatedAt time.Time
}
