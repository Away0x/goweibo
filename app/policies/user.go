package policies

import (
	"gin_weibo/app/models"

	"github.com/gin-gonic/gin"
)

// UserPolicyUpdate : 是否有更新目标 user 的权限
func UserPolicyUpdate(c *gin.Context, currentUser *models.User, targetUserID int) bool {
	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}

// UserPolicyDestory : 是否有删除用户的权限 (只有当前用户拥有管理员权限且删除的用户不是自己时)
func UserPolicyDestory(c *gin.Context, currentUser *models.User, targetUserID int) bool {
	if currentUser.ID == uint(targetUserID) || !currentUser.IsAdminRole() {
		Unauthorized(c)
		return false
	}

	return true
}
