package user

import (
  "gorm.io/gorm"
  "goweibo/core/pkg/password"
)

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
  if !passwordEncrypted(u.Password) {
    pwd, err := password.Encrypt(u.Password)
    if err == nil {
      u.Password = pwd
    }
  }

  return err
}

// 判断密码是否加密过了
func passwordEncrypted(pwd string) (status bool) {
  return len(pwd) == 60 // 长度等于 60 说明加密过了
}
