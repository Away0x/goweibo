package controllers

import (
	"fmt"
	"gin_weibo/app/helpers"
	"gin_weibo/config"
	"gin_weibo/pkg/flash"
	"html/template"
	"net/http"
	"strconv"

	"gin_weibo/app/auth"
	viewmodels "gin_weibo/app/view_models"

	"github.com/gin-gonic/gin"
)

type (
	renderObj = map[string]interface{}
)

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(renderObj)
	flashStore := flash.Read(c)
	oldValueStore := flash.ReadOldFormValue(c)
	validateMsgArr := flash.ReadValidateMessage(c)

	// flash 数据
	obj[flash.FlashInContextAndCookieKeyName] = flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[flash.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	obj[flash.ValidateContextAndCookieKeyName] = validateMsgArr
	// csrf
	if config.AppConfig.EnableCsrf {
		if csrfHtml, ok := csrfField(c); ok {
			obj["csrfField"] = csrfHtml
		}
	}

	// 获取当前登录的用户 (如果用户登录了的话，中间件中会通过 session 存储用户数据)
	if user, err := auth.GetUserFromContext(c); err == nil {
		obj[config.AppConfig.ContextUserDataKey] = viewmodels.NewUserViewModelSerializer(user, 140)
	}

	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}

	c.HTML(http.StatusOK, tplPath, obj)
}

// RenderError : 渲染错误页面
func RenderError(c *gin.Context, code int, msg string) {
	errorCode := code
	if code == 419 || code == 403 {
		errorCode = 403
	}

	c.HTML(code, "error/error.html", gin.H{
		"errorMsg":  msg,
		"errorCode": errorCode,
		"errorImg":  helpers.Static("/svg/" + strconv.Itoa(code) + ".svg"),
		"backUrl":   "/",
	})
}

// Render403 -
func Render403(c *gin.Context) {
	RenderError(c, http.StatusForbidden, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
}

// Render404 -
func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "页面没找到")
}

// private ---------------------
func csrfField(c *gin.Context) (template.HTML, bool) {
	token := c.Keys[config.AppConfig.CsrfParamName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", false
	}

	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, config.AppConfig.CsrfParamName, tokenStr)), true
}
