package database

import (
	"database/sql"
	"fmt"
	"goweibo/core"
	"goweibo/core/pkg/db"

	"goweibo/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDefaultDatabase 初始化默认的数据库
func SetupDefaultDatabase() (*gorm.DB, *sql.DB) {
	dsn := db.BuildDatabaseDSN(core.GetConfig().DefaultString("DB.DEFAULT.CONNECTION", "mysql"), db.DatabaseConfig{
		UserName: core.GetConfig().String("DB.DEFAULT.USERNAME"),
		Password: core.GetConfig().String("DB.DEFAULT.PASSWORD"),
		Host:     core.GetConfig().String("DB.DEFAULT.HOST"),
		Port:     core.GetConfig().String("DB.DEFAULT.PORT"),
		DBName:   core.GetConfig().String("DB.DEFAULT.DATABASE"),
		Options:  core.GetConfig().String("DB.DEFAULT.OPTIONS"),
	}, func(s string) string {
		return core.GetConfig().String("DB.DEFAULT.DATABASE") + "_" + string(core.GetConfig().AppRunMode())
	})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(getGormLoggerLevel()),
	})
	if err != nil {
		panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("[SetupDefaultDatabase#newConnection error]: " + err.Error() + " " + dsn)
	}

	sqlDB.SetMaxOpenConns(core.GetConfig().Int("DB.DEFAULT.MAX_OPEN_CONNECTIONS"))
	sqlDB.SetMaxIdleConns(core.GetConfig().Int("DB.DEFAULT.MAX_IDLE_CONNECTIONS"))

	fmt.Printf("\nDefault atabase connection successful: %s\n", dsn)
	return db, sqlDB
}

// RegisterAutoMigrateModle 注册需要自动迁移的 model
func RegisterAutoMigrateModle() []interface{} {
	return []interface{}{
    models.User{},
	}
}

func getGormLoggerLevel() logger.LogLevel {
	if core.GetConfig().IsDev() {
		return logger.Info
	}

	return logger.Silent
}
