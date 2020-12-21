package bootstrap

import (
	"goweibo/config"
	"goweibo/core"
)

// SetupConfig 初始化配置
func SetupConfig(configFilePath, configFileType string) {
	config.Setup(configFilePath, configFileType)
	core.NewAppConfig()

	config.WriteConfig(core.GetConfig().String("APP.TEMP_DIR") + "/config.json")
}
