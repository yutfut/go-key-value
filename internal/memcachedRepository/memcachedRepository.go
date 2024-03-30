package memcachedRepository

import (
	"context"

	"redis/internal/models"
	
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedInterface interface {
	Set(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
	Get(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
}

type MemcachedRepository struct {
	memcached *memcache.Client
}

func NewMemcachedRepository(memcached *memcache.Client) MemcachedInterface {
	return &MemcachedRepository{
		memcached: memcached,
	}
}

func (m *MemcachedRepository) Set(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	return data, m.memcached.Set(&memcache.Item{Key: data.Key, Value: []byte(data.Value)})
}

func (m *MemcachedRepository) Get(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	result, err := m.memcached.Get(data.Key)
	if err != nil {
		return nil, err
	}

	data.Value = string(result.Value)
	return data, nil
}