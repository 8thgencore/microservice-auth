package token

import (
	"context"
	"time"

	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-common/pkg/cache"
)

type repo struct {
	redisClient     cache.Client
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewRepository creates a new instance of TokenRepository.
func NewRepository(
	redisClient cache.Client,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) repository.TokenRepository {
	return &repo{
		redisClient:     redisClient,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

// AddRevokedToken adds a revoked refresh token to Redis with a TTL (time-to-live).
func (r *repo) AddRevokedToken(ctx context.Context, refreshToken string) error {
	if err := r.redisClient.SetEx(ctx, refreshToken, true, r.refreshTokenTTL); err != nil {
		return err
	}

	return nil
}

// IsTokenRevoked checks if a refresh token is in the list of revoked tokens.
func (r *repo) IsTokenRevoked(ctx context.Context, refreshToken string) (bool, error) {
	_, err := r.redisClient.Get(ctx, refreshToken)
	if err != nil {
		if string(err.Error()) == "key not found" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// NeedUpdateAccessToken implements repository.TokenRepository.
func (r *repo) NeedUpdateAccessToken(ctx context.Context, userID string) (bool, error) {
	panic("unimplemented")
}

// UpdateUserVersion implements repository.TokenRepository.
func (r *repo) UpdateUserVersion(ctx context.Context, userID string, version int) error {
	panic("unimplemented")
}
