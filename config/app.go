package config

import (
	"github.com/spf13/viper"
)

// 应用程序配置
type appConfig struct {
	// 应用名称
	Name string
	// 运行模式: debug, release, test
	RunMode string
	// 运行 addr
	Addr string
	// 完整 url
	URL string
	// secret key
	Key string

	// 静态资源存放路径
	PublicPath string
	// 模板等前端源码文件存放路径
	ResourcesPath string
	// 模板文件存放的路径
	ViewsPath string

	// 是否开启 csrf
	EnableCsrf bool
	// csrf param name
	CsrfParamName string
	// csrf header
	CsrfHeaderName string

	// auth session key
	AuthSessionKey string
	// Context 中当前用户数据的 key
	ContextCurrentUserDataKey string
}

func newAppConfig() *appConfig {
	// 默认配置
	viper.SetDefault("APP.NAME", "gin_weibo")
	viper.SetDefault("APP.RUNMODE", "release")
	viper.SetDefault("APP.ADDR", ":8080")
	viper.SetDefault("APP.URL", "")
	viper.SetDefault("APP.KEY", "base64:O+VQ74YEigLPDzLKnh2HW/yjCdU2ON9v7xuKBgSOEAo=")
	viper.SetDefault("APP.ENABLE_CSRF", true)

	return &appConfig{
		Name:    viper.GetString("APP.NAME"),
		RunMode: viper.GetString("APP.RUNMODE"),
		Addr:    viper.GetString("APP.ADDR"),
		URL:     viper.GetString("APP.URL"),
		Key:     viper.GetString("APP.KEY"),

		PublicPath:    "public",
		ResourcesPath: "resources",
		ViewsPath:     "resources/views",

		EnableCsrf:     viper.GetBool("APP.ENABLE_CSRF"),
		CsrfParamName:  "_csrf",
		CsrfHeaderName: "X-CsrfToken",

		AuthSessionKey:            "gin_session",
		ContextCurrentUserDataKey: "currentUserData",
	}
}
