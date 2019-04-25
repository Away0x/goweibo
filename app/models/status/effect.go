package status

import (
	"gin_weibo/database"

	"github.com/lexkong/log"
)

// Create -
func (s *Status) Create() (err error) {
	if err = database.DB.Create(&s).Error; err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	return nil
}
