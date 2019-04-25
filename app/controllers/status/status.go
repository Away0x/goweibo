package status

import (
	userModel "gin_weibo/app/models/user"

	"github.com/gin-gonic/gin"
)

// Store 创建微博
func Store(g *gin.Context, currentUser *userModel.User) {}

// Destroy 删除微博
func Destroy(g *gin.Context, currentUser *userModel.User) {}
