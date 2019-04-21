package policies

import (
	"gin_weibo/app/models"

	"github.com/gin-gonic/gin"
)

// UserPlolicyUpdate : 是否有更新目标 user 的权限
func UserPlolicyUpdate(c *gin.Context, currentUser *models.User, targetUserID int) bool {
	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}
