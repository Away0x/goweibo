package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"

	"gin_weibo/config"
)

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	db, err := gorm.Open(config.DBConfig.Connection, config.DBConfig.URL)
	if err != nil {
		log.Fatal("Database connection failed. Database url: "+config.DBConfig.URL+" error: ", err)
	} else {
		fmt.Print("\n\n------------------------------------------ GORM OPEN SUCCESS! -----------------------------------------------\n\n")
	}

	db.LogMode(config.DBConfig.Debug)
	DB = db

	return db
}
