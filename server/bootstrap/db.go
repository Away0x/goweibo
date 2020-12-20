package bootstrap

import (
	"goweibo/core"
	"goweibo/database"
)

// SetupDB 初始化数据库
func SetupDB() {
	db, sqlDB := database.SetupDefaultDatabase()

	// 自动迁移
	db.AutoMigrate()

	core.NewDefaultConnection(db, sqlDB)
}
