package service

import (
	"context"
	"github.com/eydeveloper/highload-social-messenger-counter/internal/entity"
	"github.com/redis/go-redis/v9"
)

type CounterService struct {
	redisClient *redis.Client
}

func NewCounterService(redisClient *redis.Client) *CounterService {
	return &CounterService{
		redisClient: redisClient,
	}
}

func (s *CounterService) Increment(message entity.Message) error {
	ctx := context.Background()
	err := s.redisClient.Incr(ctx, message.ReceiverId+":"+message.DialogId).Err()

	if err != nil {
		return err
	}

	return nil
}
