package policies

import (
	statusModel "gin_weibo/app/models/status"
	userModel "gin_weibo/app/models/user"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// StatusPolicyDestroy 是否有删除微博的权限
func StatusPolicyDestroy(c *gin.Context, currentUser *userModel.User, status *statusModel.Status) bool {
	if currentUser.ID != status.UserID {
		log.Infof("%s 没有权限删除微博 (ID: %d)", currentUser.Name, status.UserID)
		Unauthorized(c)
		return false
	}

	return true
}
