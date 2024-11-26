package interceptor

import (
	"context"
	"errors"

	"github.com/8thgencore/microservice-auth/internal/tokens"
	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
	"github.com/8thgencore/microservice-auth/pkg/utils"
	"google.golang.org/grpc"
)

type Auth struct {
	TokenOperations tokens.TokenOperations
}

// Map of endpoints that do not require authorization
var publicEndpoints = map[string]struct{}{
	"/auth_v1.AuthV1/Login":         {},
	"/auth_v1.AuthV1/RefreshTokens": {},
	"/auth_v1.AuthV1/Logout":        {},
}

// Map of endpoints that are only accessible by admins
var adminEndpoints = map[string]struct{}{
	"/auth_v1.AuthV1/AdminEndpoint1": {},
	"/auth_v1.AuthV1/AdminEndpoint2": {},
}

// AuthInterceptor is used for authorization.
func (c *Auth) AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Checking whether the current method is in the list of public endpoints
	if _, exists := publicEndpoints[info.FullMethod]; exists {
		return handler(ctx, req)
	}

	// Extract token from context
	token, err := utils.ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	// Verify access token
	claims, err := c.TokenOperations.VerifyAccessToken(token)
	if err != nil {
		return nil, err
	}

	// Checking whether the current method is in the list of admin endpoints
	if _, exists := adminEndpoints[info.FullMethod]; exists {
		if claims.Role != userv1.Role_name[int32(userv1.Role_ADMIN)] {
			return nil, errors.New("access denied: insufficient permissions")
		}
	}

	return handler(ctx, req)
}
