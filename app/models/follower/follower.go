package follower

const (
	tableName = "followers"
)

// Follower 粉丝
type Follower struct {
	ID         uint `gorm:"column:id;primary_key;AUTO_INCREMENT;not null"`
	UserID     uint `gorm:"column:user_id;not null" sql:"index"`     // 多对多，关联 User Model (关注者)
	FollowerID uint `gorm:"column:follower_id;not null" sql:"index"` // 多对多，关联 User Model (粉丝)
}

// TableName 表名
func (Follower) TableName() string {
	return tableName
}
