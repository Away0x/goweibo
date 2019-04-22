package models

import (
	"crypto/md5"
	"encoding/hex"
	"gin_weibo/database"
	"gin_weibo/pkg/auth"
	"strconv"
	"time"
)

// User 用户模型
type User struct {
	BaseModel
	Name            string    `gorm:"column:name;type:varchar(255);not null"`
	Email           string    `gorm:"column:email;type:varchar(255);unique;not null"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255);not null"`
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
func (u *User) Get(id int) error {
	d := database.DB.First(&u, id)
	return d.Error
}

// GetByEmail 根据 email 来获取用户
func (u *User) GetByEmail(email string) error {
	d := database.DB.Where("email = ?", email).First(&u)
	return d.Error
}

// All 获取所有用户
func (User) All() ([]*User, error) {
	users := make([]*User, 0)
	d := database.DB.Find(&users)
	return users, d.Error
}

// List 获取用户列表
func (User) List(offset, limit int) ([]*User, error) {
	users := make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// AllCount 总用户数
func (u *User) AllCount() int {
	count := 0
	database.DB.Table(u.TableName()).Count(&count)
	return count
}

// Create 创建用户
func (u *User) Create() error {
	if err := u.Encrypt(); err != nil {
		return err
	}
	return database.DB.Create(&u).Error
}

// Update 更新用户
func (u *User) Update(needEncryotPwd bool) error {
	if needEncryotPwd {
		if err := u.Encrypt(); err != nil {
			return err
		}
	}

	return database.DB.Save(&u).Error
}

// Delete : 删除用户
func (u *User) Delete(id int) error {
	u.BaseModel.ID = uint(id)
	return database.DB.Delete(&u).Error
}

// Encrypt 对密码进行加密
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// Compare 验证用户密码
func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)
	return
}

// Gravatar 生成用户头像
func (u *User) Gravatar() string {
	if u.Avatar != "" {
		return u.Avatar
	}

	hash := md5.Sum([]byte(u.Email))
	return "http://www.gravatar.com/avatar/" + hex.EncodeToString(hash[:])
}

// GetIDstring 获取字符串形式的 id
func (u *User) GetIDstring() string {
	return strconv.Itoa(int(u.ID))
}

// IsAdmin 是否为管理员
func (u *User) IsAdminRole() bool {
	return u.IsAdmin == TrueTinyint
}
