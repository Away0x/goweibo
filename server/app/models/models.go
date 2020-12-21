package models

import (
	"database/sql"
	"goweibo/core"
	"strconv"
	"time"

	"gorm.io/gorm"
)

const (
	// TrueTinyint true
	TrueTinyint uint = 1
	// FalseTinyint false
	FalseTinyint uint = 0
)

// BaseModel model 基类
type BaseModel struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// IDString id to string
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// DB 获取默认数据库
func DB() *sql.DB {
	return core.GetDefaultConnection()
}

// TinyBool tinyint => bool
func TinyBool(i uint) bool {
	return i == TrueTinyint
}
