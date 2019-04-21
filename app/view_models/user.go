package viewmodels

import (
	"gin_weibo/app/models"
)

// UserViewModel 用户
type UserViewModel struct {
	ID     int
	Name   string
	Email  string
	Avatar string
}

// NewUserViewModelSerializer 用户数据展示
func NewUserViewModelSerializer(u *models.User) UserViewModel {
	return UserViewModel{
		ID:     int(u.ID),
		Name:   u.Name,
		Email:  u.Email,
		Avatar: u.Gravatar(),
	}
}
