package redis

import (
	"fmt"
	"redis/pkg/conf"

	"github.com/redis/go-redis/v9"
)

func RedisConn(config *conf.Conf) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})
}