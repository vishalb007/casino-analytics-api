package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"casino-analytics/internal/repository"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo  *repository.TransactionRepo
	cache *redis.Client
}

func NewAnalyticsService(r *repository.TransactionRepo, c *redis.Client) *Service {
	return &Service{repo: r, cache: c}
}

func (s *Service) GetGGR(ctx context.Context, from, to time.Time) (interface{}, error) {
	key := fmt.Sprintf("ggr:%s:%s", from, to)

	if val, err := s.cache.Get(ctx, key).Result(); err == nil {
		var data interface{}
		json.Unmarshal([]byte(val), &data)
		return data, nil
	}

	data, err := s.repo.GetGGR(ctx, from, to)
	if err != nil {
		return nil, err
	}

	bytes, _ := json.Marshal(data)
	s.cache.Set(ctx, key, bytes, time.Minute*5)

	return data, nil
}