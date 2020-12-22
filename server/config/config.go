package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type (
	// Runmode 开发模式
	Runmode string
)

const (
	// RunmodeProduction 生产环境
	RunmodeProduction Runmode = "production"
	// RunmodeStaging 准生产环境
	RunmodeStaging Runmode = "staging"
	// RunmodeDevelopment 调试、开发环境
	RunmodeDevelopment Runmode = "development"
	// RunmodeTest 测试环境
	RunmodeTest Runmode = "test"
)

// Setup 初始化配置
// configFilePath: 配置文件路径
// configFileType: 配置文件类型
func Setup(configFilePath, configFileType string) {
	// 初始化 viper 配置
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType(configFileType)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("读取配置文件失败，请检查 %s 配置文件是否存在: %v", configFilePath, err))
	}

	// 设置配置默认值
	setupDefaultConfig()

	// 环境变量 (设置环境变量: export APPNAME_APP_RUNMODE=development)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(viper.GetString("APP.NAME")) // 环境变量前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 监听配置文件变化
	watchConfig()
}

// WriteConfig 写配置到文件
func WriteConfig(filename string) {
	viper.WriteConfigAs(filename)
}

// 监控配置文件变化
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(ev fsnotify.Event) {
		log.Infof("Config file changed: %s", ev.Name)
	})
}
