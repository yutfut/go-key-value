package redisrepository

import (
	"context"

	"redis/internal/models"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Set(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
	Get(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
}

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) RedisInterface {
	return &RedisRepository{
		redis: redis,
	}
}

func (r *RedisRepository) Set(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	return data, r.redis.Set(ctx, data.Key, data.Value, 0).Err()
}

func (r *RedisRepository) Get(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	result, err := r.redis.Get(ctx, data.Key).Result()
	if err != nil {
		return nil, err
	}

	data.Value = result
	return data, err
}