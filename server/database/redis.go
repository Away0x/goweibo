package database

import (
	"fmt"
	"goweibo/core"

	"github.com/go-redis/redis"
)

// SetupRedis 初始化 redis
func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     core.GetConfig().String("REDIS.ADDR"),
		Password: core.GetConfig().String("REDIS.PASSWORD"),
		DB:       core.GetConfig().Int("REDIS.DATABASE"),
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic("[SetupRedis error]: " + err.Error())
	} else {
		fmt.Printf("redis connect ping response: %s", pong)
	}

	return client
}
