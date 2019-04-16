package config

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
}
