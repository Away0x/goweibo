package viewmodels

import (
	userModel "gin_weibo/app/models/user"
)

// UserViewModel 用户
type UserViewModel struct {
	ID      int
	Name    string
	Email   string
	Avatar  string
	IsAdmin bool
}

// NewUserViewModelSerializer 用户数据展示
func NewUserViewModelSerializer(u *userModel.User) *UserViewModel {
	return &UserViewModel{
		ID:      int(u.ID),
		Name:    u.Name,
		Email:   u.Email,
		Avatar:  u.Gravatar(),
		IsAdmin: u.IsAdminRole(),
	}
}
