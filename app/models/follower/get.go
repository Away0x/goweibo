package follower

import (
	"fmt"
	userModel "gin_weibo/app/models/user"
	"gin_weibo/database"
)

// Followers 获取粉丝列表
func Followers(userID, offset, limit int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", tableName)
	if limit == 0 {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Order("id").Find(&followers)
		return followers, d.Error
	} else {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
		return followers, d.Error
	}
}

// Followings 获取用户关注人列表
func Followings(userID, offset, limit int) (followers []*userModel.User, err error) {
	followers = make([]*userModel.User, 0)
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", tableName)
	if limit == 0 {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Order("id").Find(&followers)
		return followers, d.Error
	} else {
		d := database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Offset(offset).Limit(limit).Order("id").Find(&followers)
		return followers, d.Error
	}
}

// FollowingsIDList 获取用户关注人 ID 列表
func FollowingsIDList(userID int) (followerIDList []uint) {
	followers, _ := Followings(userID, 0, 0)
	followerIDList = make([]uint, 0)
	for _, v := range followers {
		followerIDList = append(followerIDList, v.ID)
	}
	return
}

// FollowingsCount 关注数
func FollowingsCount(userID int) (count int, err error) {
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.user_id", tableName)
	err = database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.follower_id = ?", userID).Count(&count).Error
	return
}

// FollowersCount 粉丝数
func FollowersCount(userID int) (count int, err error) {
	joinSQL := fmt.Sprintf("inner join %s on users.id = followers.follower_id", tableName)
	err = database.DB.Model(&userModel.User{}).Joins(joinSQL).Where("followers.user_id = ?", userID).Count(&count).Error
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
