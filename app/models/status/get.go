package status

import (
	userModel "gin_weibo/app/models/user"
	"gin_weibo/database"
)

// Get -
func Get(id int) (*Status, error) {
	s := &Status{}
	d := database.DB.First(&s, id)
	return s, d.Error
}

// GetUser 通过 status_id 获取该微博的所有者
func GetUser(statusID int) (*userModel.User, error) {
	s, err := Get(statusID)
	if err != nil {
		return nil, err
	}

	u, err := userModel.Get(int(s.UserID))
	if err != nil {
		return nil, err
	}

	return u, nil
}

// GetUserStatus 获取该用户的所有微博
func GetUserStatus(userID, offset, limit int) ([]*Status, error) {
	status := make([]*Status, 0)

	err := database.DB.Where("user_id = ?", userID).Offset(
		offset).Limit(limit).Order("id desc").Find(&status).Error

	if err != nil {
		return status, err
	}

	return status, nil
}

// GetUserAllStatusCount 获取该用户的所有微博 的 count
func GetUserAllStatusCount(userID int) (count int, err error) {
	err = database.DB.Model(&Status{}).Where("user_id = ?", userID).Count(&count).Error
	return
}
