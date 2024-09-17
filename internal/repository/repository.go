package repository

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/model"
)

// UserRepository is the interface for user info repository communication.
type UserRepository interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
	GetAuthInfo(ctx context.Context, username string) (*model.AuthInfo, error)
}

// AccessRepository is the interface for access policies repository communication.
type AccessRepository interface {
	GetRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error)
	AddRoleEndpoint(ctx context.Context, endpoint string, allowedRoles []string) error
	UpdateRoleEndpoint(ctx context.Context, endpoint string, allowedRoles []string) error
	DeleteRoleEndpoint(ctx context.Context, endpoint string) error
}

// LogRepository is the interface for transaction log repository communication.
type LogRepository interface {
	Log(ctx context.Context, log *model.Log) error
}

// TokenRepository is the interface for revoked token repository communication.
type TokenRepository interface {
	AddRevokedToken(ctx context.Context, refreshToken string) error
	IsTokenRevoked(ctx context.Context, refreshToken string) (bool, error)
}
