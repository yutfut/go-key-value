package redis

import "github.com/redis/go-redis/v9"

func RedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "password",
		DB:       0,
	})
}