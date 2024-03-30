package redisrepository

import (
	"context"

	"redis/internal/models"

	"github.com/redis/go-redis/v9"
)

type RedisInterface interface {
	Set(ctx context.Context, data *models.Redis) (*models.Redis, error)
	Get(ctx context.Context, data *models.Redis) (*models.Redis, error)
}

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{
		redis: redis,
	}
}

func (h *RedisRepository) Set(ctx context.Context, data *models.Redis) (*models.Redis, error) {
	return data, h.redis.Set(ctx, data.Key, data.Value, 0).Err()
}

func (h *RedisRepository) Get(ctx context.Context, data *models.Redis) (*models.Redis, error) {
	result, err := h.redis.Get(ctx, data.Key).Result()
	data.Value = result
	return data, err
}