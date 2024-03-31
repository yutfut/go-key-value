package redis

import (
	"context"
	"fmt"
	"go-key-value/pkg/conf"

	"github.com/redis/go-redis/v9"
)

func RedisConn(config *conf.Conf) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}


	return rdb, nil
}