package controllers

import (
	"gin_weibo/app/models"
	"gin_weibo/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Redirect : 路由重定向
func Redirect(c *gin.Context, redirectRoute string) {
	c.Redirect(http.StatusMovedPermanently, config.AppConfig.URL+redirectRoute)
}

// 重定向到用户展示页面
func RedirectToUserShowPage(c *gin.Context, user *models.User) {
	Redirect(c, "/users/show/"+strconv.Itoa(int(user.ID)))
}

// 重定向到登录页面
func RedirectToLoginPage(c *gin.Context) {
	Redirect(c, "/login")
}

// 重定向到用户创建页面 (注册页面)
func RedirectToUserCreatePage(c *gin.Context) {
	Redirect(c, "/users/create")
}
