package user

import (
	userModel "gin_weibo/app/models/user"

	"github.com/gin-gonic/gin"
)

// Followings 用户关注者列表
func Followings(c *gin.Context, currentUser *userModel.User) {

}

// Followers 用户粉丝列表
func Followers(c *gin.Context, currentUser *userModel.User) {

}
