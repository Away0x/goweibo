package models

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "gorm.io/gorm"
  "goweibo/core/pkg/password"
  "goweibo/core/pkg/strutils"
  "time"
)

// User 用户模型
type User struct {
  BaseModel
  Name            string     `gorm:"type:varchar(255);not null" json:"name"`
  Email           string     `gorm:"type:varchar(255);unique;not null" json:"email"`
  Avatar          string     `gorm:"type:varchar(255);not null" json:"avatar"`
  Password        string     `gorm:"type:varchar(255);not null" json:"-"`
  ActivationToken string     `gorm:"type:varchar(255);comment:用于用户激活" json:"-"`
  RememberToken   string     `gorm:"type:varchar(100);comment:用于实现记住我功能" json:"-"`
  IsAdmin         uint       `gorm:"type:tinyint(1);comment:是否为管理员" json:"is_admin"`
  Activated       uint       `gorm:"type:tinyint(1);not null;comment:用户是否激活" json:"activated"`
  EmailVerifiedAt *time.Time `gorm:"comment:邮件激活时间" json:"-"`

  Statuses []Status
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
  if !password.Encrypted(u.Password) {
    pwd, err := password.Encrypt(u.Password)
    if err == nil {
      u.Password = pwd
    }
  }

  if u.RememberToken == "" {
    u.RememberToken = string(strutils.RandomCreateBytes(10))
  }

  if u.ActivationToken == "" {
    u.ActivationToken = string(strutils.RandomCreateBytes(30))
  }

  return err
}

func GetUserByID(id uint) (user *User, err error) {
  user = new(User)
  err = DB().First(user, id).Error
  return
}

func GetUserByEmail(email string) (user *User, err error) {
  user = new(User)
  err = DB().Where("email = ?", email).First(user).Error
  return
}

func (u *User) ComparePassword(pwd string) (err error) {
  err = password.Compare(u.Password, pwd)
  return
}

func (u *User) Gravatar() string {
  if u.Avatar != "" {
    return u.Avatar
  }

  hash := md5.Sum([]byte(u.Email))
  return fmt.Sprintf("http://www.gravatar.com/avatar/%s?d=identicon", hex.EncodeToString(hash[:]))
}

// IsAdminRole 是否为管理员
func (u *User) IsAdminRole() bool {
  return u.IsAdmin == TrueTinyint
}

// IsActivated 是否已激活
func (u *User) IsActivated() bool {
  return u.Activated == TrueTinyint
}
