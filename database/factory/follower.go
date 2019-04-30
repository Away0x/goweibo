package factory

import (
	followerModel "gin_weibo/app/models/follower"
	userModel "gin_weibo/app/models/user"
)

// FollowerTableSeeder -
func FollowerTableSeeder(needCleanTable bool) {
	if needCleanTable {
		DropAndCreateTable(&followerModel.Follower{})
	}

	users, err := userModel.All()
	if err != nil {
		panic("follower mock error!")
	}
	user := users[0]
	userID := user.ID

	// 获取去除掉 ID 为 1 的所有用户 ID 数组
	followers := users[1:]
	followerIDs := make([]uint, 0)
	for _, v := range followers {
		followerIDs = append(followerIDs, v.ID)
	}

	// 关注除了 1 号用户以外的所有用户
	followerModel.DoFollow(userID, followerIDs...)
	// 除了 1 号用户以外的所有用户都来关注 1 号用户
	for _, v := range followerIDs {
		followerModel.DoFollow(v, userID)
	}
}
