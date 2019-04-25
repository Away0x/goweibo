package controllers

import (
	"fmt"
	"gin_weibo/app/helpers"
	"gin_weibo/config"
	"gin_weibo/pkg/flash"
	"gin_weibo/routes/named"
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

const (
	csrfInputHTML = "csrfField"
	csrfTokenName = "csrfToken"
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
		if csrfHTML, csrfToken, ok := csrfField(c); ok {
			obj[csrfInputHTML] = csrfHTML
			obj[csrfTokenName] = csrfToken
		}
	}

	// 获取当前登录的用户 (如果用户登录了的话，中间件中会通过 session 存储用户数据)
	if user, err := auth.GetCurrentUserFromContext(c); err == nil {
		obj[config.AppConfig.ContextCurrentUserDataKey] = viewmodels.NewUserViewModelSerializer(user)
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
		"backUrl":   named.G("root"),
	})
}

// Render403 -
func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg)
}

// Render404 -
func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}

// RenderUnauthorized -
func RenderUnauthorized(c *gin.Context) {
	Render403(c, "很抱歉，您没有权限访问该页面")
}

// private ---------------------
func csrfField(c *gin.Context) (template.HTML, string, bool) {
	token := c.Keys[config.AppConfig.CsrfParamName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", "", false
	}

	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, config.AppConfig.CsrfParamName, tokenStr)), tokenStr, true
}
