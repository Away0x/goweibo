package core

import (
	"goweibo/config"

	"github.com/spf13/viper"
)

// AppConfig app 配置
type AppConfig struct{}

// NewAppConfig 初始化项目配置
func NewAppConfig() {
	appConfig = &AppConfig{}
}

// String 获取 string 配置值
func (*AppConfig) String(key string) string {
	return viper.GetString(key)
}

// DefaultString 获取 string 配置值，可设置默认值
func (*AppConfig) DefaultString(key string, defaultVal string) string {
	v := viper.GetString(key)
	if v == "" {
		return defaultVal
	}

	return v
}

// Int 获取 int 配置值
func (*AppConfig) Int(key string) int {
	return viper.GetInt(key)
}

// DefaultInt 获取 int 配置值，可设置默认值
func (*AppConfig) DefaultInt(key string, defaultVal int) int {
	v := viper.GetInt(key)
	if v == 0 {
		return defaultVal
	}

	return v
}

// Bool 获取 bool 配置值
func (*AppConfig) Bool(key string) bool {
	return viper.GetBool(key)
}

// IsDev 是否为开发模式
func (c *AppConfig) IsDev() bool {
	return c.AppRunMode() == config.RunmodeDevelopment
}

// AppRunMode 获取当前应用的启动模式
func (c *AppConfig) AppRunMode() config.Runmode {
	mode := config.Runmode(c.String("APP.RUNMODE"))

	switch mode {
	case config.RunmodeProduction:
		return config.RunmodeProduction
	case config.RunmodeStaging:
		return config.RunmodeStaging
	case config.RunmodeDevelopment:
		return config.RunmodeDevelopment
	case config.RunmodeTest:
		return config.RunmodeTest
	default:
		return config.RunmodeDevelopment
	}
}
