package controllers

import (
	"gin_weibo/config"
	"gin_weibo/routes/named"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirect : 路由重定向 use path
func Redirect(c *gin.Context, redirectPath string, withRoot bool) {
	path := redirectPath
	if withRoot {
		path = config.AppConfig.URL + redirectPath
	}

	redirect(c, path)
}

// RedirectRouter : 路由重定向 use router name
func RedirectRouter(c *gin.Context, routerName string, args ...interface{}) {
	redirect(c, named.G(routerName, args...))
}

// RedirectToLoginPage : 重定向到登录页面
func RedirectToLoginPage(c *gin.Context) {
	loginPath := named.G("login.create")

	if c.Request.Method == http.MethodPost {
		redirect(c, loginPath)
		return
	}

	redirect(c, loginPath+"?back="+c.Request.URL.Path)
}

// ------------------------ private
func redirect(c *gin.Context, redirectPath string) {
	// 千万注意，这个地方不能用 301(永久重定向)
	c.Redirect(http.StatusFound, redirectPath)
}
