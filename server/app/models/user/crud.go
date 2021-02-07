package user

import "goweibo/app/models"

func Get(id uint) (user *User, err error) {
  user = new(User)
  d := models.DB().First(user, id)
  return user, d.Error
}

func GetByEmail(email string) (user *User, err error) {
  user = new(User)
  d := models.DB().Where("email = ?", email).First(user)
  return user, d.Error
}
