package redisrepository

import (
	"context"

	"go-key-value/internal/interfaces"
	"go-key-value/internal/models"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) interfaces.KeyValueRepositoryInterface {
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

func (r *RedisRepository) Del(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	return data, r.redis.Del(ctx, data.Key).Err()
}