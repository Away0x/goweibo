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
	// ConfigUtil 配置文件工具方法
	ConfigUtil struct{}
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
func Setup(configFilePath, configFileType string) *ConfigUtil {
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

	return &ConfigUtil{}
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

// String 获取 string 配置值
func (*ConfigUtil) String(key string) string {
	return viper.GetString(key)
}

// DefaultString 获取 string 配置值，可设置默认值
func (*ConfigUtil) DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}

	return v
}

// Int 获取 int 配置值
func (*ConfigUtil) Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt 获取 int 配置值，可设置默认值
func (*ConfigUtil) DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool 获取 bool 配置值
func (*ConfigUtil) Bool(key string) bool {
	return viper.GetBool(key)
}

// IsDev 是否为开发模式
func (c *ConfigUtil) IsDev() bool {
	return c.AppRunMode() == RunmodeDevelopment
}

// AppRunMode 获取当前应用的启动模式
func (c *ConfigUtil) AppRunMode() Runmode {
	mode := Runmode(c.String("APP.RUNMODE"))

	switch mode {
	case RunmodeProduction:
		return RunmodeProduction
	case RunmodeStaging:
		return RunmodeStaging
	case RunmodeDevelopment:
		return RunmodeDevelopment
	case RunmodeTest:
		return RunmodeTest
	default:
		return RunmodeDevelopment
	}
}
