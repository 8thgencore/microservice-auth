package token

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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

// SetTokenVersion sets the token version.
func (r *repo) SetTokenVersion(ctx context.Context, userID string, version int) error {
	key := "token_version:" + userID
	if err := r.redisClient.SetEx(ctx, key, version, r.accessTokenTTL); err != nil {
		return fmt.Errorf("could not set user version: %w", err)
	}

	return nil
}

// GetTokenVersion gets the current token version from the cache.
func (r *repo) GetTokenVersion(ctx context.Context, userID string) (int, error) {
	key := "token_version:" + userID
	rawVersion, err := r.redisClient.Get(ctx, key)
	if err != nil {
		if strings.Contains(err.Error(), "key not found") {
			return 0, nil
		}
		return 0, fmt.Errorf("could not get user version: %w", err)
	}

	version, err := strconv.Atoi(rawVersion)
	if err != nil {
		return 0, fmt.Errorf("could not parse user version: %w", err)
	}

	return version, nil
}
