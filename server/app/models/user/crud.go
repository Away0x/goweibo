package user

import "goweibo/app/models"

func GetUser(id uint) (user *User, err error) {
  user = new(User)
  if err = models.DB().First(user, id).Error; err != nil {
    return nil, err
  }

  return
}
