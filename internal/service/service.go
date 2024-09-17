package service

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/model"
)

// UserService is the interface for service communication.
type UserService interface {
	Create(ctx context.Context, user *model.UserCreate) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}

// AuthService is the interface for service communication.
type AuthService interface {
	Login(ctx context.Context, creds *model.UserCreds) (*model.TokenPair, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error)
	Logout(ctx context.Context, refreshToken string) error
}

// AccessService is the interface for service communication.
type AccessService interface {
	Check(ctx context.Context, endpoint string) error
	GetRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error)
	AddRoleEndpoint(ctx context.Context, endpoint string, roles []string) error
	UpdateRoleEndpoint(ctx context.Context, endpoint string, roles []string) error
	DeleteRoleEndpoint(ctx context.Context, endpoint string) error
}
