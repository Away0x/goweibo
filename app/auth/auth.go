package auth

import (
	"errors"
	"gin_weibo/app/models"
	"gin_weibo/config"
	"gin_weibo/pkg/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SaveUserToContext : 保存用户数据到 context 中
func SaveUserToContext(c *gin.Context) {
	user, err := GetUserFromSession(c)
	if err != nil {
		return
	}

	c.Keys[config.AppConfig.ContextUserDataKey] = user
}

// GetUserFromSession : 从 session 中获取用户
func GetUserFromSession(c *gin.Context) (*models.User, error) {
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

// GetUserFromContext : 从 context 中获取用户模型
func GetUserFromContext(c *gin.Context) (*models.User, error) {
	err := errors.New("没有获取到用户数据")
	userDataFromContext := c.Keys[config.AppConfig.ContextUserDataKey]
	if userDataFromContext == nil {
		return nil, err
	}

	if user, ok := userDataFromContext.(*models.User); !ok {
		return nil, err
	} else {
		return user, nil
	}
}

// 登录
func Login(c *gin.Context, u *models.User) {
	session.SetSession(c, config.AppConfig.AuthSessionKey, u.GetIDstring())
}

// 登出
func Logout(c *gin.Context) {
	session.DeleteSession(c, config.AppConfig.AuthSessionKey)
}
