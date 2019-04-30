package follower

import (
	"fmt"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/database"
)

// Followers 获取粉丝列表
func Followers(userID int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", tableName)
	d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Find(&followers)
	return followers, d.Error
}

// Followings 获取用户关注人列表
func Followings(userID int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", tableName)
	d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Find(&followers)
	return followers, d.Error
}

// FollowingsIDList 获取用户关注人 ID 列表
func FollowingsIDList(userID int) (followerIDList []uint) {
	followers, _ := Followings(userID)
	followerIDList = make([]uint, len(followers))
	for _, v := range followers {
		followerIDList = append(followerIDList, v.ID)
	}
	return
}

// IsFollowing 已经关注了
func IsFollowing(currentUserID, userID int) bool {
	followerIDList := FollowingsIDList(currentUserID)
	id := uint(userID)
	for _, v := range followerIDList {
		if id == v {
			return true
		}
	}
	return false
}
