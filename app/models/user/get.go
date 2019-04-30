package user

import (
	"crypto/md5"
	"encoding/hex"
	"gin_weibo/app/models"
	"gin_weibo/database"
	"strconv"
)

// Get -
func Get(id int) (*User, error) {
	user := &User{}
	d := database.DB.First(&user, id)
	return user, d.Error
}

// GetByEmail -
func GetByEmail(email string) (*User, error) {
	user := &User{}
	d := database.DB.Where("email = ?", email).First(&user)
	return user, d.Error
}

// GetByActivationToken -
func GetByActivationToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("activation_token = ?", token).First(&user)
	return user, d.Error
}

// GetByRememberToken -
func GetByRememberToken(token string) (*User, error) {
	user := &User{}
	d := database.DB.Where("remember_token = ?", token).First(&user)
	return user, d.Error
}

// List 获取用户列表
func List(offset, limit int) (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Offset(offset).Limit(limit).Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// All -
func All() (users []*User, err error) {
	users = make([]*User, 0)

	if err := database.DB.Order("id").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// AllCount 总用户数
func AllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}

// Gravatar 获取用户头像
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
	return u.IsAdmin == models.TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == models.TrueTinyint
}
