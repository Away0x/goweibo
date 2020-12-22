package config

import (
	"time"

	"github.com/spf13/viper"
)

const (
	defaultTempDir = "storage"
	defaultAppPort = ":9999"
	defaultAppName = "app"
)

var now = time.Now()

// 默认配置
var defaultConfigMap = map[string]interface{}{
	// app
	"APP.NAME":          defaultAppName,
	"APP.VERSION":       "1.0.0",
	"APP.RUNMODE":       "production", // 环境
	"APP.PORT":          defaultAppPort,
	"APP.URL":           "http://localhost" + defaultAppPort,
	"APP.KEY":           "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
	"APP.TEMP_DIR":      defaultTempDir,    // 临时文件存储位置 (log ...)
	"APP.PUBLIC_DIR":    "public",          // public 文件夹
	"APP.UPLOAD_DIR":    "public/uploads",  // 文件上传文件夹
	"APP.RESOURCES_DIR": "resources",       // resources 文件夹
	"APP.TEMPLATE_DIR":  "resources/views", // 模板文件存放文件夹
	"APP.GZIP":          true,              // 是否开启 gzip
	"APP.TEMPLATE_EXT":  "tpl",             // 模版文件后缀

	// db
	"DB.DEFAULT.CONNECTION":           "mysql",
	"DB.DEFAULT.HOST":                 "127.0.0.1",
	"DB.DEFAULT.PORT":                 "3306",
	"DB.DEFAULT.DATABASE":             defaultAppName,
	"DB.DEFAULT.USERNAME":             "root",
	"DB.DEFAULT.PASSWORD":             "",
	"DB.DEFAULT.OPTIONS":              "charset=utf8&parseTime=true&loc=Local",
	"DB.DEFAULT.MAX_OPEN_CONNECTIONS": 100,
	"DB.DEFAULT.MAX_IDLE_CONNECTIONS": 20,
	"DB.DEFAULT.AUTO_MIGRATE":         true,

	// log
	"LOG.PREFIX":     "[ZAP_LOGGER]",
	"LOG.FOLDER":     defaultTempDir + "/logs/zap",
	"LOG.LEVEL":      "debug", // 日志级别: debug, info, warn, error, dpanic, panic, fatal
	"LOG.MAXSIZE":    10,      // 在进行切割之前，日志文件的最大大小（以MB为单位）
	"LOG.MAXBACKUPS": 5,       // 保留旧文件的最大个数
	"LOG.MAXAGES":    30,      // 保留旧文件的最大天数
}

// 设置配置默认值
func setupDefaultConfig() {
	for k, v := range defaultConfigMap {
		viper.SetDefault(k, v)
	}
}
