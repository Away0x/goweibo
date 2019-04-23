package controllers

import (
	"gin_weibo/app/auth"
	"gin_weibo/app/models"
	"gin_weibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

// MustGetCurrentUser : 获取不到用户则跳转登录页
func MustGetCurrentUser(c *gin.Context) *models.User {
	currentUser, err := auth.GetCurrentUserFromContext(c)

	if currentUser == nil || err != nil {
		RedirectToLoginPage(c)
		return nil
	}

	return currentUser
}

// MustIsGuset : 已登录用户不能进行的操作
func MustIsGuset(c *gin.Context) bool {
	// 用户已经登录了则跳转到 root page
	currentUser, err := auth.GetCurrentUserFromContext(c)
	if currentUser != nil || err == nil {
		flash.NewInfoFlash(c, "您已登录，无需再次操作。")
		RedirectToRootPage(c)
		return false
	}

	return true
}
