package auth

import (
	"errors"
	"gin_weibo/app/models"
	"gin_weibo/config"
	"gin_weibo/pkg/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	rememberFormKey = "remember"
)

// Login 登录
func Login(c *gin.Context, u *models.User) {
	// 记住我
	rememberMe := c.PostForm(rememberFormKey) == "on"
	if rememberMe {

	}

	session.SetSession(c, config.AppConfig.AuthSessionKey, u.GetIDstring())
}

// Logout 登出
func Logout(c *gin.Context) {
	session.DeleteSession(c, config.AppConfig.AuthSessionKey)
}

// -------------- private --------------
// getCurrentUserFromSession : 从 session 中获取用户
func getCurrentUserFromSession(c *gin.Context) (*models.User, error) {
	user := new(models.User)
	idStr := session.GetSession(c, config.AppConfig.AuthSessionKey)
	if idStr == "" {
		return nil, errors.New("没有获取到 session")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}

	if err = user.Get(id); err != nil {
		return nil, err
	}

	return user, nil
}
