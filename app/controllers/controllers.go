package controllers

import (
	"net/http"

	"gin_weibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

type (
	renderObj = map[string]interface{}
)

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(renderObj)
	flashStore := flash.Read(c)
	oldValueStore := flash.ReadOldFromValue(c)

	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	obj[flash.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	for k, v := range data {
		obj[k] = v
	}

	c.HTML(http.StatusOK, tplPath, obj)
}

// Redirect : 路由重定向
func Redirect(c *gin.Context, redirectRoute string) {
	c.Redirect(http.StatusMovedPermanently, redirectRoute)
}
