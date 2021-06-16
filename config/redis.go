package config

import (
	"github.com/go-redis/redis/v8"
)
var RDB *redis.Client

func InitClient() (  *redis.Client) {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})



	return RDB
}