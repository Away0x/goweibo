package models

import (
	"crypto/md5"
	"encoding/hex"
	"gin_weibo/database"
	"gin_weibo/pkg/auth"
	"strconv"
	"time"

	"github.com/lexkong/log"
)

// User 用户模型
type User struct {
	BaseModel
	Name            string    `gorm:"column:name;type:varchar(255);not null"`
	Email           string    `gorm:"column:email;type:varchar(255);unique;not null"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255);not null"`
	EmailVerifiedAt time.Time `gorm:"column:email_verified_at"`
	Password        string    `gorm:"column:password;type:varchar(255);not null"`
	IsAdmin         uint      `gorm:"column:is_admin;type:tinyint(1)"`
	ActivationTOken string    `gorm:"column:activation_token;type:varchar(255)"`
	Activated       uint      `gorm:"column:activated;type:tinyint(1);not null"`
}

// TableName 表名
func (User) TableName() string {
	return "users"
}

// Get 获取一个用户
func (u *User) Get(id int) (err error) {
	if err = database.DB.First(&u, id).Error; err != nil {
		log.Warnf("用户获取失败: %v", err)
		return err
	}

	return nil
}

// GetByEmail 根据 email 来获取用户
func (u *User) GetByEmail(email string) (err error) {
	if err = database.DB.Where("email = ?", email).First(&u).Error; err != nil {
		log.Warnf("用户获取失败: %v", err)
		return err
	}

	return nil
}

// All 获取所有用户
func (User) All() (users []*User, err error) {
	users = make([]*User, 0)

	if err = database.DB.Find(&users).Error; err != nil {
		log.Warnf("用户获取失败: %v", err)
		return users, err
	}

	return users, nil
}

// List 获取用户列表
func (User) List(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		log.Warnf("用户获取失败: %v", err)
		return users, err
	}

	return users, nil
}

// AllCount 总用户数
func (u *User) AllCount() (count int) {
	count = 0
	database.DB.Table(u.TableName()).Count(&count)
	return count
}

// Create 创建用户
func (u *User) Create() (err error) {
	if err = u.Encrypt(); err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	if err = database.DB.Create(&u).Error; err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	return nil
}

// Update 更新用户
func (u *User) Update(needEncryotPwd bool) (err error) {
	if needEncryotPwd {
		if err = u.Encrypt(); err != nil {
			log.Warnf("用户更新失败: %v", err)
			return err
		}
	}

	if err = database.DB.Save(&u).Error; err != nil {
		log.Warnf("用户更新失败: %v", err)
		return err
	}

	return nil
}

// Delete : 删除用户
func (u *User) Delete(id int) (err error) {
	u.BaseModel.ID = uint(id)

	if err = database.DB.Delete(&u).Error; err != nil {
		log.Warnf("用户删除失败: %v", err)
		return err
	}

	return nil
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

// IsAdminRole 是否为管理员
func (u *User) IsAdminRole() bool {
	return u.IsAdmin == TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == TrueTinyint
}
