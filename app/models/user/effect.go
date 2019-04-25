package user

import (
	"github.com/lexkong/log"
	"gin_weibo/pkg/utils"
	"gin_weibo/database"
)

// Create -
func (u *User) Create() (err error) {
  if err = u.Encrypt(); err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	// 生成用户 remember_token
	if u.RememberToken == "" {
		u.RememberToken = string(utils.RandomCreateBytes(10))
	}
	// 生成用户激活 token
	if u.ActivationToken == "" {
		u.ActivationToken = string(utils.RandomCreateBytes(30))
	}

	if err = database.DB.Create(&u).Error; err != nil {
		log.Warnf("用户创建失败: %v", err)
		return err
	}

	return nil
}

// Update 更新用户
func (u *User) Update(needEncryotPwd bool) (err error) {
	if needEncryotPwd {
		if err = u.Encrypt(); err != nil {
			log.Warnf("用户更新失败: %v", err)
			return err
		}
	}

	if err = database.DB.Save(&u).Error; err != nil {
		log.Warnf("用户更新失败: %v", err)
		return err
	}

	return nil
}

// Delete -
func Delete(id int) (err error) {
  user := &User{}
  user.BaseModel.ID = uint(id)

  // Unscoped: 永久删除而不是软删除 (由于该操作是管理员操作的，所以不使用软删除)
	if err = database.DB.Unscoped().Delete(&user).Error; err != nil {
		log.Warnf("用户删除失败: %v", err)
		return err
	}

	return nil
}
