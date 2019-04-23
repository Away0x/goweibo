package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_weibo/app/controllers"
	"gin_weibo/config"
	"gin_weibo/pkg/utils"
)

// Csrf : csrf middleware
func Csrf() gin.HandlerFunc {
	return func(c *gin.Context) {

		if config.AppConfig.EnableCsrf {
			// cookie 中获取 csrf token (如没有则设置)
			csrfToken := getCsrfTokenFromCookie(c)

			// POST 并且开启了 csrf
			if c.Request.Method == http.MethodPost {
				// params 中获取 csrf token
				paramCsrfToken := getCsrfTokenFromParamsOrHeader(c)

				if paramCsrfToken == "" || paramCsrfToken != csrfToken {
					controllers.Render403(c, "很抱歉！您的 Session 已过期，请刷新后再试一次。")
					c.Abort()
					return
				}
			}
		}

		c.Next()
	}
}

// 从 cookie 中获取 csrf token
func getCsrfTokenFromCookie(c *gin.Context) (token string) {
	keyName := config.AppConfig.CsrfParamName

	if s, err := c.Request.Cookie(keyName); err == nil {
		token = s.Value
	}

	if token == "" {
		token = string(utils.RandomCreateBytes(32))
		c.SetCookie(keyName, token, 0, "/", "", false, false)
	}
	c.Keys[keyName] = token

	return token
}

// 从 params 或 headers 中获取 csrf token
func getCsrfTokenFromParamsOrHeader(c *gin.Context) (token string) {
	req := c.Request

	if req.Form == nil {
		req.ParseForm()
	}

	// 从 params 中获取
	token = req.FormValue(config.AppConfig.CsrfParamName)
	if token == "" {
		// 从 headers 中获取
		token = req.Header.Get(config.AppConfig.CsrfHeaderName)
	}

	return token
}
