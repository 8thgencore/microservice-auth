package access

import (
	"context"
	"errors"
	"slices"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/8thgencore/microservice-auth/internal/converter"
	"github.com/8thgencore/microservice-auth/internal/model"
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

var (
	// ErrEndpointAlreadyExists occurs when trying to add an endpoint that already exists.
	ErrEndpointAlreadyExists = errors.New("endpoint already exists")
	// ErrEndpointDoesNotExist occurs when trying to access or modify an endpoint that does not exist.
	ErrEndpointDoesNotExist = errors.New("endpoint does not exist")
)

func (s *accessService) Check(ctx context.Context, endpoint string) error {
	token, err := s.extractToken(ctx)
	if err != nil {
		return err
	}

	s.rolesMutex.RLock()
	// Check if the accessibleRoles map is nil and initialize it if necessary
	if s.accessibleRoles == nil {
		s.rolesMutex.RUnlock() // Unlock read lock before acquiring write lock
		s.rolesMutex.Lock()    // Lock for write to initialize map
		defer s.rolesMutex.Unlock()

		// Double-check in case it was initialized by another goroutine
		if s.accessibleRoles == nil {
			endpointPermissions, errRepo := s.accessRepository.GetRoleEndpoints(ctx)
			if errRepo != nil {
				return ErrFailedToReadAccessPolicy
			}
			s.accessibleRoles = converter.ToEndpointPermissionsMap(endpointPermissions)
		}
	} else {
		defer s.rolesMutex.RUnlock()
	}

	roles, ok := s.accessibleRoles[endpoint]
	if !ok {
		return ErrEndpointNotFound
	}

	claims, err := s.tokenOperations.VerifyAccessToken(token, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return ErrInvalidAccessToken
	}

	if !slices.Contains(roles, claims.Role) {
		return ErrAccessDenied
	}

	return nil
}

// GetRoleEndpoints retrieves the list of resources after verifying access permissions.
func (s *accessService) GetRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error) {
	err := s.Check(ctx, "/access_v1.AccessV1/GetRoleEndpoints")
	if err != nil {
		return nil, err
	}

	resources, err := s.accessRepository.GetRoleEndpoints(ctx)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

// AddRoleEndpoint adds a new resource after verifying access permissions.
func (s *accessService) AddRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	err := s.Check(ctx, "/access_v1.AccessV1/AddRoleEndpoint")
	if err != nil {
		return err
	}

	err = s.accessRepository.AddRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		return err
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	s.accessibleRoles[endpoint] = roles
	return nil
}

// UpdateRoleEndpoint edits an existing resource after verifying access permissions.
func (s *accessService) UpdateRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	err := s.Check(ctx, "/access_v1.AccessV1/UpdateRoleEndpoint")
	if err != nil {
		return err
	}

	err = s.accessRepository.UpdateRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		return err
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	s.accessibleRoles[endpoint] = roles
	return nil
}

// DeleteRoleEndpoint deletes a resource after verifying access permissions.
func (s *accessService) DeleteRoleEndpoint(ctx context.Context, endpoint string) error {
	err := s.Check(ctx, "/access_v1.AccessV1/DeleteRoleEndpoint")
	if err != nil {
		return err
	}

	err = s.accessRepository.DeleteRoleEndpoint(ctx, endpoint)
	if err != nil {
		return err
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	delete(s.accessibleRoles, endpoint)
	return nil
}

func (s *accessService) extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrMetadataNotProvided
	}

	authHeader, ok := md[authMetadataHeader]
	if !ok || len(authHeader) == 0 {
		return "", ErrAuthHeaderNotProvided
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return "", ErrInvalidAuthHeaderFormat
	}

	return strings.TrimPrefix(authHeader[0], authPrefix), nil
}
