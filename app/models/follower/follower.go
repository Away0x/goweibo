package follower

import (
	"gin_weibo/app/models"
)

// Follower 粉丝
type Follower struct {
	models.BaseModel
	UserID     uint `gorm:"column:user_id;not null" sql:"index"`     // 多对多，关联 User Model (关注者)
	FollowerID uint `gorm:"column:follower_id;not null" sql:"index"` // 多对多，关联 User Model (粉丝)
}

// TableName 表名
func (Follower) TableName() string {
	return "followers"
}

// // Followers 获取粉丝列表
// func Followers() (users []*userModel.User, err error) {}

// // Followings 获取用户关注人列表
// func Followings() (users []*userModel.User, err error) {}

// // DoFollow 关注
// func DoFollow() (err error) {}

// // DoUnFollow 取消关注
// func DoUnFollow() (err error) {}

// // IsFollowing 已经关注了
// func IsFollowing() bool {}
