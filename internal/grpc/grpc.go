package grpc

import (
	"context"

	"go-key-value/internal/interfaces"
	"go-key-value/internal/models"

	proto "go-key-value/pkg/keyvalue"
)

type server struct {
	proto.UnimplementedKeyvalueServer
	redis interfaces.KeyValueRepositoryInterface
}

func NewAuthHandler(redis interfaces.KeyValueRepositoryInterface) *server {
	return &server{
		redis: redis,
	}
}

func (s *server) Set(ctx context.Context, in *proto.KeyValue) (*proto.KeyValue, error) {
	requestData, err := s.redis.Set(ctx, &models.KeyValue{
		Key: in.Key,
		Value: in.Value,
	})
	if err != nil {
		return &proto.KeyValue{}, err
	}
	return &proto.KeyValue{
		Key: requestData.Key,
		Value: requestData.Value,
	}, nil
}

func (s *server) Get(ctx context.Context, in *proto.KeyValue) (*proto.KeyValue, error) {
	requestData, err := s.redis.Get(ctx, &models.KeyValue{Key: in.Key})
	if err != nil {
		return &proto.KeyValue{}, err
	}
	return &proto.KeyValue{
		Key: requestData.Key,
		Value: requestData.Value,
	}, nil
}

func (s *server) Del(ctx context.Context, in *proto.KeyValue) (*proto.KeyValue, error) {
	requestData, err := s.redis.Del(ctx, &models.KeyValue{Key: in.Key})
	if err != nil {
		return &proto.KeyValue{}, err
	}
	return &proto.KeyValue{
		Key: requestData.Key,
		Value: requestData.Value,
	}, nil
}