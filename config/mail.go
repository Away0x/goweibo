package config

import "github.com/spf13/viper"

// 邮件配置
type mailConfig struct {
	Driver   string // smtp or log (log 时邮件是写在 log 中的，便于调试)
	Host     string // 邮箱的服务器地址
	Port     int    // 邮箱的服务器端口
	User     string // 发送者的 name
	Password string // 授权码或密码
	FromName string // 用来作为邮件的发送者名称
}

func newMailConfig() *mailConfig {
	// 默认配置
	viper.SetDefault("MAIL.MAIL_DRIVER", "gin_weibo")
	viper.SetDefault("MAIL.MAIL_HOST", "")
	viper.SetDefault("MAIL.MAIL_PORT", 25)
	viper.SetDefault("MAIL.MAIL_USERNAME", "")
	viper.SetDefault("MAIL.MAIL_PASSWORD", "")
	viper.SetDefault("MAIL.MAIL_FROM_NAME", "gin_weibo")

	return &mailConfig{
		Driver:   viper.GetString("MAIL.MAIL_DRIVER"),
		Host:     viper.GetString("MAIL.MAIL_HOST"),
		Port:     viper.GetInt("MAIL.MAIL_PORT"),
		User:     viper.GetString("MAIL.MAIL_USERNAME"),
		Password: viper.GetString("MAIL.MAIL_PASSWORD"),
		FromName: viper.GetString("MAIL.MAIL_FROM_NAME"),
	}
}
