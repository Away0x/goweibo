package policies

import (
	userModel "gin_weibo/app/models/user"

	"github.com/gin-gonic/gin"
)

// UserPolicyUpdate : 是否有更新目标 user 的权限
func UserPolicyUpdate(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID != uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}

// UserPolicyDestroy : 是否有删除用户的权限 (只有当前用户拥有管理员权限且删除的用户不是自己时)
func UserPolicyDestroy(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID == uint(targetUserID) || !currentUser.IsAdminRole() {
		Unauthorized(c)
		return false
	}

	return true
}

// UserPolicyFollow : 是否有关注用户的权限
func UserPolicyFollow(c *gin.Context, currentUser *userModel.User, targetUserID int) bool {
	if currentUser.ID == uint(targetUserID) {
		Unauthorized(c)
		return false
	}

	return true
}
