package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// 启动模式
	runmodeDebug   = "debug"
	runmodeRelease = "release"
	runmodeTest    = "test"

	// 配置文件路径
	configFilePath = "./config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

var (
	// ProjectConfig 项目固定配置
	ProjectConfig = &projectConfig{
		PublicPath: "public",
	}
	// AppConfig 应用配置
	AppConfig *appConfig
	// DBConfig 数据库配置
	DBConfig *dbConfig
)

// InitConfig 初始化配置
func InitConfig() error {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 初始化 app 配置
	AppConfig = newAppConfig()
	// 初始化数据库配置
	DBConfig = newDBConfig()

	// 热更新配置文件
	watchConfig()

	return nil
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		// 配置文件更新了
		log.Printf("Config file changed: %s", ev.Name)
	})
}
