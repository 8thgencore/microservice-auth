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
	// ErrMetadataNotProvided occurs when metadata is not passed in the request.
	ErrMetadataNotProvided = errors.New("metadata is not provided")
	// ErrAuthHeaderNotProvided occurs when the authorization header is missing from the request.
	ErrAuthHeaderNotProvided = errors.New("authorization header is not provided")
	// ErrInvalidAuthHeaderFormat occurs when the authorization header has an incorrect format.
	ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")
	// ErrFailedToReadAccessPolicy occurs when the access policy could not be read.
	ErrFailedToReadAccessPolicy = errors.New("failed to read access policy")
	// ErrEndpointNotFound occurs when the specified endpoint is not found.
	ErrEndpointNotFound = errors.New("failed to find endpoint")
	// ErrInvalidAccessToken occurs when the access token is invalid.
	ErrInvalidAccessToken = errors.New("access token is invalid")
	// ErrAccessDenied occurs when access to the requested resource is denied.
	ErrAccessDenied = errors.New("access denied")
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
