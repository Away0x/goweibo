package bootstrap

import (
	"goweibo/config"
	"goweibo/core"
  "goweibo/core/pkg/jwttoken"
)

// SetupConfig 初始化配置
func SetupConfig(configFilePath, configFileType string) {
	config.Setup(configFilePath, configFileType)
	core.NewAppConfig()
  config.WriteConfig(core.GetConfig().String("APP.TEMP_DIR") + "/config.json")

	// config jwt token
	jwttoken.SetupToken(&jwttoken.Config{
    SecretKey: core.GetConfig().String("APP.KEY"),
    AccessTokenLifeTime: core.GetConfig().Duration("TOKEN.ACCESS_TOKEN_LIFETIME"),
    RefreshTokenLifeTime: core.GetConfig().Duration("TOKEN.REFRESH_TOKEN_LIFETIME"),
  })
}
