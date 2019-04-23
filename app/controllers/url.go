package controllers

import (
	"gin_weibo/app/models"
	"gin_weibo/config"
)

// CreateSignupConfirmURL 生成完整的用户激活的接口地址
func CreateSignupConfirmURL(u *models.User) string {
	return config.AppConfig.URL + "/signup/confirm/" + u.ActivationToken
}
