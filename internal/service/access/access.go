package access

import (
	"context"
	"errors"
	"slices"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/8thgencore/microservice-auth/internal/converter"
)

const (
	authMetadataHeader = "authorization"
	authPrefix         = "Bearer "
)

var (
	ErrMetadataNotProvided      = errors.New("metadata is not provided")
	ErrAuthHeaderNotProvided    = errors.New("authorization header is not provided")
	ErrInvalidAuthHeaderFormat  = errors.New("invalid authorization header format")
	ErrFailedToReadAccessPolicy = errors.New("failed to read access policy")
	ErrEndpointNotFound         = errors.New("failed to find endpoint")
	ErrInvalidAccessToken       = errors.New("access token is invalid")
	ErrAccessDenied             = errors.New("access denied")
)

var accessibleRoles map[string][]string

func (s *serv) Check(ctx context.Context, endpoint string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ErrMetadataNotProvided
	}

	authHeader, ok := md[authMetadataHeader]
	if !ok || len(authHeader) == 0 {
		return ErrAuthHeaderNotProvided
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return ErrInvalidAuthHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	if accessibleRoles == nil {
		endpointPermissions, errRepo := s.accessRepository.GetRoleEndpoints(ctx)
		if errRepo != nil {
			return ErrFailedToReadAccessPolicy
		}
		accessibleRoles = converter.ToEndpointPermissionsMap(endpointPermissions)
	}

	roles, ok := accessibleRoles[endpoint]
	if !ok {
		return ErrEndpointNotFound
	}

	claims, err := s.tokenOperations.VerifyAccessToken(accessToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return ErrInvalidAccessToken
	}

	if !slices.Contains(roles, claims.Role) {
		return ErrAccessDenied
	}

	return nil
}
