package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type AuthRepository struct {
	redisClient *redis.Client
}

func NewAuthRepository(redisClient *redis.Client) *AuthRepository {
	return &AuthRepository{
		redisClient: redisClient,
	}
}

func (r AuthRepository) WriteToken(ctx context.Context, id string, token string) error {
	return r.redisClient.Set(ctx, id, token, 0).Err()
}

func (r AuthRepository) GetToken(ctx context.Context, id string) (string, error) {
	token, err := r.redisClient.Get(ctx, id).Result()
	return token, err
}
