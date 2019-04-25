package user

import (
	"crypto/md5"
	"strconv"
	"gin_weibo/pkg/auth"
	"gin_weibo/app/models"
	"encoding/hex"
)

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
	return u.IsAdmin == models.TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
	return u.Activated == models.TrueTinyint
}
