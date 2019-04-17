package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/jinzhu/gorm"
)

// DB gorm
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB(dialect string, url string, hasLog bool) *gorm.DB {
	db, err := gorm.Open(dialect, url)
	if err != nil {
		log.Fatalf("Database connection failed. Database url: %s error: %v", url, err)
	} else {
		log.Println("gorm open success!")
	}

	db.LogMode(hasLog)
	DB = db

	return db
}
