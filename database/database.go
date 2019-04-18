package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"

	"gin_weibo/config"
)

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatalf("Database connection failed. Database url: %s error: %v", config.DBConfig.URL, err)
	} else {
		log.Println("gorm open success!")
	}

	db.LogMode(config.DBConfig.Debug)
	DB = db

	return db
}
