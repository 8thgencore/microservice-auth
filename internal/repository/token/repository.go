package token

import (
	"context"
	"time"

	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-common/pkg/cache"
)

type repo struct {
	redisClient cache.Client
	ttl         time.Duration // Time-to-live for revoked tokens
}

// NewRepository creates a new instance of TokenRepository.
func NewRepository(redisClient cache.Client, ttl time.Duration) repository.TokenRepository {
	return &repo{
		redisClient: redisClient,
		ttl:         ttl,
	}
}

// AddRevokedToken adds a revoked refresh token to Redis with a TTL (time-to-live).
func (r *repo) AddRevokedToken(ctx context.Context, refreshToken string) error {
	if err := r.redisClient.Set(ctx, refreshToken, true); err != nil {
		return err
	}

	if err := r.redisClient.Expire(ctx, refreshToken, r.ttl); err != nil {
		return err
	}

	return nil
}

// IsTokenRevoked checks if a refresh token is in the list of revoked tokens.
func (r *repo) IsTokenRevoked(ctx context.Context, refreshToken string) (bool, error) {
	result, err := r.redisClient.Get(ctx, refreshToken)
	if result == nil {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return string(result.([]byte)) == "1", nil
}
