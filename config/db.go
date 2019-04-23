package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 数据库配置
type dbConfig struct {
	Connection string
	Host       string
	Port       int
	Database   string
	Username   string
	Password   string

	URL string

	Debug bool
}

func newDBConfig() *dbConfig {
	// 默认配置
	viper.SetDefault("DB.CONNECTION", "mysql")
	viper.SetDefault("DB.HOST", "127.0.0.1")
	viper.SetDefault("DB.PORT", 3306)
	viper.SetDefault("DB.DATABASE", viper.GetString("APP.NAME"))
	viper.SetDefault("DB.USERNAME", "gin")
	viper.SetDefault("DB.PASSWORD", "")

	username := viper.GetString("DB.USERNAME")
	password := viper.GetString("DB.PASSWORD")
	host := viper.GetString("DB.HOST")
	port := viper.GetInt("DB.PORT")
	database := viper.GetString("DB.DATABASE")
	url := createDBURL(username, password, host, port, database)

	return &dbConfig{
		Connection: viper.GetString("DB.CONNECTION"),
		Host:       host,
		Port:       port,
		Database:   database,
		Username:   username,
		Password:   password,
		URL:        url,
		Debug:      AppConfig.RunMode == RunmodeDebug,
	}
}

func createDBURL(uname string, pwd string, host string, port int, name string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		uname, pwd,
		host, port,
		name, true, "Local")
}
