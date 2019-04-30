package user

import (
	"gin_weibo/pkg/auth"
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
