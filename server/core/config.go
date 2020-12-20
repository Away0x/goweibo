package core

import "goweibo/config"

// NewAppConfig 初始化项目配置
func NewAppConfig(c *config.ConfigUtil) {
	appConfig = c
}
