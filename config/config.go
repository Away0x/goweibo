package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	// 配置文件路径
	configFilePath = "./config.yaml"
	// 配置文件格式
	configFileType = "yaml"
)

var (
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

	// 初始化默认配置
	setDefaultConfig()

	// 初始化 app 配置
	AppConfig = &appConfig{
		Name:    viper.GetString("APP.NAME"),
		RunMode: viper.GetString("APP.RUNMODE"),
		Addr:    viper.GetString("APP.ADDR"),
		URL:     viper.GetString("APP.URL"),
		Key:     viper.GetString("APP.KEY"),
	}

	// 初始化数据库配置
	DBConfig = &dbConfig{
		Connection: viper.GetString("DB.CONNECTION"),
		Host:       viper.GetString("DB.HOST"),
		Port:       viper.GetInt("DB.PORT"),
		Database:   viper.GetString("DB.DATABASE"),
		Username:   viper.GetString("DB.USERNAME"),
		Password:   viper.GetString("DB.PASSWORD"),
	}

	// 热更新配置文件
	watchConfig()

	return nil
}

// setDefaultConfig 默认配置
func setDefaultConfig() {
	// app 默认配置
	viper.SetDefault("APP.NAME", "gin_weibo")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.URL", "")
	viper.SetDefault("APP.KEY", "base64:O+VQ74YEigLPDzLKnh2HW/yjCdU2ON9v7xuKBgSOEAo=")

	// 数据库 默认配置
	viper.SetDefault("DB.CONNECTION", "mysql")
	viper.SetDefault("DB.HOST", "127.0.0.1")
	viper.SetDefault("DB.PORT", 3306)
	viper.SetDefault("DB.DATABASE", viper.GetString("APP.NAME"))
	viper.SetDefault("DB.USERNAME", "gin")
	viper.SetDefault("DB.PASSWORD", "")
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		// 配置文件更新了
		log.Printf("Config file changed: %s", ev.Name)
	})
}
