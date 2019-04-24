package models

import (
	"gin_weibo/database"
	"gin_weibo/pkg/utils"
	"time"

	"github.com/lexkong/log"
)

// PasswordReset 重置密码模型
type PasswordReset struct {
	Email     string    `gorm:"column:email;type:varchar(255);not null" sql:"index"`
	Token     string    `gorm:"column:token;type:varchar(255);not null" sql:"index"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

// TableName 表名
func (PasswordReset) TableName() string {
	return "password_resets"
}

// GetByEmail -
func (p *PasswordReset) GetByEmail(email string) (err error) {
	if err = database.DB.Where("email = ?", email).First(&p).Error; err != nil {
		return err
	}

	return nil
}

// GetByToken -
func (p *PasswordReset) GetByToken(token string) (err error) {
	if err = database.DB.Where("token = ?", token).First(&p).Error; err != nil {
		return err
	}

	return nil
}

// Create 创建重置密码的数据
func (p *PasswordReset) Create() (token string, err error) {
	token = string(utils.RandomCreateBytes(30))

	// 如已存在则先删除 (可以判断下，不能创建太频繁)
	if err = p.GetByEmail(p.Email); err == nil {
		if err = p.Delete(p.Email); err != nil {
			return "", err
		}
	}

	// 创建
	p.Token = token
	if err = database.DB.Create(&p).Error; err != nil {
		log.Warnf("%s 创建失败: %v", p.TableName(), err)
		return "", err
	}

	return token, nil
}

// Delete : 删除
func (p *PasswordReset) Delete(email string) (err error) {
	if err = database.DB.Where("email = ?", email).Delete(&p).Error; err != nil {
		log.Warnf("%s 删除失败: %v", p.TableName(), err)
		return err
	}

	return nil
}

// DeleteByToken : 删除
func (p *PasswordReset) DeleteByToken(token string) (err error) {
	if err = database.DB.Where("token = ?", token).Delete(&p).Error; err != nil {
		log.Warnf("%s 删除失败: %v", p.TableName(), err)
		return err
	}

	return nil
}
