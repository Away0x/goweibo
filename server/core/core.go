package core

import (
	"database/sql"
	"goweibo/config"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	application       *Application
	defaultConnection *GormConnection
	appConfig         *config.ConfigUtil
)

// GetApplication 获取全局应用实例
func GetApplication() *Application {
	if application == nil {
		panic("application is not initialized")
	}

	return application
}

// GetApplicationEngine 获取全局应用实例
func GetApplicationEngine() *echo.Echo {
	return GetApplication().Engine
}

// GetDefaultConnection 获取全局默认数据库实例
func GetDefaultConnection() *sql.DB {
	if defaultConnection == nil || defaultConnection.DB == nil {
		panic("default connnetion is not initialized")
	}

	return defaultConnection.DB
}

// GetDefaultConnectionEngine 获取全局默认数据库实例
func GetDefaultConnectionEngine() *gorm.DB {
	if defaultConnection == nil || defaultConnection.Engine == nil {
		panic("default connnetion is not initialized")
	}

	return defaultConnection.Engine
}

// GetConfig 获取全局配置
func GetConfig() *config.ConfigUtil {
	if appConfig == nil {
		panic("config is not initialized")
	}

	return appConfig
}
