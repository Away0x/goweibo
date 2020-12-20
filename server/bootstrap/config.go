package bootstrap

import (
	"goweibo/config"
	"goweibo/core"
)

// SetupConfig 初始化配置
func SetupConfig(configFilePath, configFileType string) {
	c := config.Setup(configFilePath, configFileType)
	core.NewAppConfig(c)

	config.WriteConfig(c.String("APP.TEMP_DIR") + "/config.json")
}
