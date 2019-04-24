package models

import (
	"gin_weibo/database"
)

// Status 微博
type Status struct {
	BaseModel
	Content string `gorm:"column:context;type:text;not null"`
	UserID  uint   `gorm:"column:user_id;not null" sql:"index"` // 一对多，关联 User Model
}

// TableName 表名
func (Status) TableName() string {
	return "statuses"
}

// Get -
func (s *Status) Get(id int) (err error) {
	if err = database.DB.First(&s, id).Error; err != nil {
		return err
	}

	return nil
}

// GetUser 获取该微博的 user
func (s *Status) GetUser(statusID int) (user *User, err error) {
	if err = database.DB.First(&s, statusID).Error; err != nil {
		return nil, err
	}

	user = &User{}
	if err = user.Get(int(s.UserID)); err != nil {
		return nil, err
	}

	return user, nil
}

// GetByUserID 获取该用户的所有微博
func (s *Status) GetByUserID(userID int) (status []*Status, err error) {
	status = make([]*Status, 0)
	if err = database.DB.Where("user_id = ?", userID).Find(&status).Error; err != nil {
		return status, nil
	}

	return status, nil
}
