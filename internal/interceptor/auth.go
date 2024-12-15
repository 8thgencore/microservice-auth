package interceptor

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/delivery/user"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
	"github.com/8thgencore/microservice-auth/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth is a struct that handles authentication.
type Auth struct {
	TokenOperations tokens.TokenOperations
	TokenRepository repository.TokenRepository
}

// Map of endpoints that do not require authorization
var publicEndpoints = map[string]struct{}{
	"/auth_v1.AuthV1/Login":         {},
	"/auth_v1.AuthV1/RefreshTokens": {},
	"/auth_v1.AuthV1/Logout":        {},
}

// Map of endpoints that are only accessible by admins
var adminEndpoints = map[string]struct{}{
	"/user_v1.UserV1/Create":                 {},
	"/user_v1.UserV1/Get":                    {},
	"/user_v1.UserV1/Update":                 {},
	"/user_v1.UserV1/Delete":                 {},
	"/access_v1.AccessV1/AddRoleEndpoint":    {},
	"/access_v1.AccessV1/UpdateRoleEndpoint": {},
	"/access_v1.AccessV1/DeleteRoleEndpoint": {},
	"/access_v1.AccessV1/GetRoleEndpoints":   {},
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
		return nil, status.Errorf(codes.Unauthenticated, "failed to extract token: %v", err)
	}

	// Verify access token
	claims, err := c.TokenOperations.VerifyAccessToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "failed to verify token: %v", err)
	}

	version, err := c.TokenRepository.GetTokenVersion(ctx, claims.Subject)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	if claims.Version < version {
		return nil, status.Errorf(codes.Unauthenticated, "token is expired")
	}

	// Checking whether the current method is in the list of admin endpoints
	if _, exists := adminEndpoints[info.FullMethod]; exists {
		if claims.Role != userv1.Role_name[int32(userv1.Role_ADMIN)] {
			return nil, status.Errorf(codes.PermissionDenied, "access denied: insufficient permissions")
		}
	}

	// Create a new context with the user ID
	ctxWithUserID := context.WithValue(ctx, user.UserIDKey, claims.Subject)

	// Pass the updated context to the handler
	return handler(ctxWithUserID, req)
}
