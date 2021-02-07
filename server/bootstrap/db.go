package bootstrap

import (
  "goweibo/core"
  "goweibo/database"
)

// SetupDB 初始化数据库
func SetupDB() {
	db, sqlDB := database.SetupDefaultDatabase()

	// 自动迁移
	if core.GetConfig().Bool("DB.DEFAULT.AUTO_MIGRATE") {
		db.AutoMigrate(database.RegisterAutoMigrateModle()...)
	}

	core.NewDefaultConnection(db, sqlDB)
}

// SetupRedis 初始化 redis
func SetupRedis() {
	client := database.SetupRedis()
	core.NewRedis(client)
}
