package user

import (
	"gin_weibo/database"
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

	if err := database.DB.Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// AllCount 总用户数
func AllCount() (count int, err error) {
	err = database.DB.Model(&User{}).Count(&count).Error
	return
}
