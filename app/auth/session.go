package auth

import (
	"errors"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/config"
	"gin_weibo/pkg/session"
	"gin_weibo/pkg/utils"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	rememberFormKey    = "remember"    // 表单提交时的 form name
	rememberCookieName = "remember_me" // 存在 cookie 中的 key name
	rememberMaxAge     = 88888888      // 过期时间
)

// Login 登录
func Login(c *gin.Context, u *userModel.User) {
	session.SetSession(c, config.AppConfig.AuthSessionKey, u.GetIDstring())
	// 记住我
	setRememberTokenInCookie(c, u)
}

// Logout 登出
func Logout(c *gin.Context) {
	session.DeleteSession(c, config.AppConfig.AuthSessionKey)
	delRememberToken(c)
}

// -------------- private --------------
// getCurrentUserFromSession : 从 session 中获取用户
func getCurrentUserFromSession(c *gin.Context) (*userModel.User, error) {
	// 从 cookie 中获取 remember me token (如有则自动登录)
	rememberMeToken := getRememberTokenFromCookie(c)
	if rememberMeToken != "" {
		if user, err := userModel.GetByRememberToken(rememberMeToken); err == nil {
			Login(c, user)
			return user, nil
		}
		delRememberToken(c)
	}

	// 从 session 中获取用户 id
	idStr := session.GetSession(c, config.AppConfig.AuthSessionKey)
	if idStr == "" {
		return nil, errors.New("没有获取到 session")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	user, err := userModel.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// -------------- 记住我功能 utils --------------
func setRememberTokenInCookie(c *gin.Context, u *userModel.User) {
	// 记住我 (如果 登录的 PostForm 中有着 remember="on" 说明开启记住我功能)
	rememberMe := c.PostForm(rememberFormKey) == "on"
	if !rememberMe {
		return
	}

	// 更新用户的 RememberToken
	newToken := string(utils.RandomCreateBytes(10))
	u.RememberToken = newToken
	if err := u.Update(false); err != nil {
		return
	}

	// 写入 cookie
	c.SetCookie(rememberCookieName, u.RememberToken, rememberMaxAge, "/", "", false, true)
}

func getRememberTokenFromCookie(c *gin.Context) string {
	if cookie, err := c.Request.Cookie(rememberCookieName); err == nil {
		if v, err := url.QueryUnescape(cookie.Value); err == nil {
			return v
		}
	}

	return ""
}

func delRememberToken(c *gin.Context) {
	c.SetCookie(rememberCookieName, "", -1, "/", "", false, true)
}
