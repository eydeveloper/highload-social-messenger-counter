package service

import (
	"github.com/eydeveloper/highload-social-messenger-counter/internal/entity"
	"github.com/redis/go-redis/v9"
)

type Counter interface {
	Increment(message entity.Message) error
}

type Service struct {
	Counter
}

func NewService(redisClient *redis.Client) *Service {
	return &Service{
		Counter: NewCounterService(redisClient),
	}
}
