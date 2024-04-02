package memcachedRepository

import (
	"context"

	"go-key-value/internal/models"
	"go-key-value/internal/interfaces"
	
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedRepository struct {
	memcached *memcache.Client
}

func NewMemcachedRepository(memcached *memcache.Client) interfaces.KeyValueRepositoryInterface {
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

func (m *MemcachedRepository) Del(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error) {
	return data, m.memcached.Delete(data.Key)
}