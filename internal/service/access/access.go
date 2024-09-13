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

var accessibleRoles map[string][]string

func (s *serv) Check(ctx context.Context, endpoint string) error {
	token, err := s.extractToken(ctx)
	if err != nil {
		return err
	}

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

	claims, err := s.tokenOperations.VerifyAccessToken(token, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return ErrInvalidAccessToken
	}

	if !slices.Contains(roles, claims.Role) {
		return ErrAccessDenied
	}

	return nil
}

// AddRoleEndpoint adds a new resource after verifying access permissions.
func (s *serv) AddRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	// Check access permissions for the "AddResource" endpoint.
	err := s.Check(ctx, "AddResource")
	if err != nil {
		return err
	}

	// Call the repository to add the resource.
	err = s.accessRepository.AddRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		return err
	}

	return nil
}

// UpdateRoleEndpoint edits an existing resource after verifying access permissions.
func (s *serv) UpdateRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	// Check access permissions for the "EditResource" endpoint.
	err := s.Check(ctx, "EditResource")
	if err != nil {
		return err
	}

	// Call the repository to edit the resource.
	err = s.accessRepository.UpdateRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRoleEndpoint deletes a resource after verifying access permissions.
func (s *serv) DeleteRoleEndpoint(ctx context.Context, endpoint string) error {
	// Check access permissions for the "DeleteResource" endpoint.
	err := s.Check(ctx, "DeleteResource")
	if err != nil {
		return err
	}

	// Call the repository to delete the resource.
	err = s.accessRepository.DeleteRoleEndpoint(ctx, endpoint)
	if err != nil {
		return err
	}

	return nil
}

// ListRoleEndpoints retrieves the list of resources after verifying access permissions.
func (s *serv) ListRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error) {
	// Check access permissions for the "GetResourceList" endpoint.
	err := s.Check(ctx, "GetResourceList")
	if err != nil {
		return nil, err
	}

	// Call the repository to get the list of resources.
	resources, err := s.accessRepository.GetRoleEndpoints(ctx)
	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (s *serv) extractToken(ctx context.Context) (string, error) {
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
