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
