package interfaces

import (
	"context"

	"go-key-value/internal/models"
)

type KeyValueRepositoryInterface interface {
	Set(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
	Get(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
	Del(ctx context.Context, data *models.KeyValue) (*models.KeyValue, error)
}