package models

import (
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
	ID        uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// IDString id to string
func (m *BaseModel) IDString() string {
	return strconv.Itoa(int(m.ID))
}

// DB 获取默认数据库
func DB() *gorm.DB {
	return core.GetDefaultConnectionEngine()
}

func CreateModel(value interface{}) (err error) {
  err = DB().Create(value).Error
  return
}

func UpdateModel(value interface{}) (err error) {
  err = DB().Save(value).Error
  return
}

func DeleteModel(value interface{}) (err error) {
  err = DB().Delete(value).Error
  return
}

// TinyBool tinyint => bool
func TinyBool(i uint) bool {
	return i == TrueTinyint
}
