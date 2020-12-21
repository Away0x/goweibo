package core

import (
	"database/sql"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

// GormConnection connection
type GormConnection struct {
	Engine *gorm.DB
	DB     *sql.DB
}

// NewDefaultConnection setup default connection
func NewDefaultConnection(e *gorm.DB, d *sql.DB) {
	defaultConnection = &GormConnection{
		Engine: e,
		DB:     d,
	}
}

// NewRedis setup redis
func NewRedis(r *redis.Client) {
	redisClient = r
}
