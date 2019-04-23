package config

import (
	"fmt"

	"github.com/lexkong/log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// RunmodeDebug -
	RunmodeDebug = "debug"
	// RunmodeRelease -
	RunmodeRelease = "release"
	// RunmodeTest -
	RunmodeTest = "test"

	// 配置文件路径
	configFilePath = "./config.yaml"
	// 日志文件路径
	logFilePath = "storage/logs/gin_weibo.log"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
	// MailConfig 邮件配置
	MailConfig *mailConfig
)

// InitConfig 初始化配置
func InitConfig() {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查: %v", err))
	}

	// 初始化日志
	initLog()
	// 初始化 app 配置
	AppConfig = newAppConfig()
	// 初始化数据库配置
	DBConfig = newDBConfig()
	// 初始化邮件配置
	MailConfig = newMailConfig()

	// 热更新配置文件
	watchConfig()
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		// 配置文件更新了
		log.Infof("Config file changed: %s", ev.Name)
	})
}
