package core

import (
	"database/sql"
	"goweibo/config"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	application       *Application
	defaultConnection *GormConnection
	appConfig         *config.ConfigUtil
	redisClient       *redis.Client
	appLog            *zap.SugaredLogger
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

// GetRedis 获取全局默认 redis 实例
func GetRedis() *redis.Client {
	if redisClient == nil {
		panic("redis is not initialized")
	}

	return redisClient
}

// GetConfig 获取全局配置
func GetConfig() *config.ConfigUtil {
	if appConfig == nil {
		panic("config is not initialized")
	}

	return appConfig
}

// GetLog 获取全局默认日志实例
func GetLog() *zap.SugaredLogger {
	if appLog == nil {
		panic("log is not initialized")
	}

	return appLog
}
