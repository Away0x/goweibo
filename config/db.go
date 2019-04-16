package config

// 数据库配置
type dbConfig struct {
	Connection string
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string
}
