package session

import (
	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

// SetSession : 设置 session
func SetSession(c *gin.Context, key string, val string) {
	session := sessions.Get(c)
	session.Set(key, val)
	session.Save()
}

// GetSession : 获取 session
func GetSession(c *gin.Context, key string) string {
	session := sessions.Get(c)
	if s, ok := session.Get(key).(string); !ok {
		return ""
	} else {
		return s
	}
}

// DeleteSession : 删除 session
func DeleteSession(c *gin.Context, key string) {
	session := sessions.Get(c)
	session.Delete(key)
	session.Save()
}
