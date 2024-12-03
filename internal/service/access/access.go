package access

import (
	"context"
	"errors"
	"slices"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/pkg/utils"
)

const (
	// Constants for service endpoints
	getRoleEndpointsEndpoint   = "/access_v1.AccessV1/GetRoleEndpoints"
	addRoleEndpointEndpoint    = "/access_v1.AccessV1/AddRoleEndpoint"
	updateRoleEndpointEndpoint = "/access_v1.AccessV1/UpdateRoleEndpoint"
	deleteRoleEndpointEndpoint = "/access_v1.AccessV1/DeleteRoleEndpoint"
)

var (
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
	// ErrFailedToGetEndpoint occurs when there is a problem retrieving an endpoint.
	ErrFailedToGetEndpoint = errors.New("failed to get endpoint")
	// ErrFailedToAddEndpoint occurs when there is a problem adding a new endpoint.
	ErrFailedToAddEndpoint = errors.New("failed to add endpoint")
	// ErrFailedToDeleteEndpoint occurs when there is a problem deleting an endpoint.
	ErrFailedToDeleteEndpoint = errors.New("failed to delete endpoint")
	// ErrFailedToUpdateEndpoint occurs when there is a problem updating an existing endpoint.
	ErrFailedToUpdateEndpoint = errors.New("failed to update endpoint")
)

func (s *accessService) Check(ctx context.Context, endpoint string) error {
	token, err := utils.ExtractToken(ctx)
	if err != nil {
		return err
	}

	claims, err := s.tokenOperations.VerifyAccessToken(token)
	if err != nil {
		return ErrInvalidAccessToken
	}

	s.rolesMutex.RLock()
	roles, ok := s.accessibleRoles[endpoint]
	s.rolesMutex.RUnlock()

	if !ok {
		return ErrEndpointNotFound
	}

	if !slices.Contains(roles, claims.Role) {
		return ErrAccessDenied
	}

	return nil
}

// GetRoleEndpoints retrieves the list of resources after verifying access permissions.
func (s *accessService) GetRoleEndpoints(ctx context.Context) ([]*model.EndpointPermissions, error) {
	err := s.Check(ctx, getRoleEndpointsEndpoint)
	if err != nil {
		return nil, err
	}

	resources, err := s.accessRepository.GetRoleEndpoints(ctx)
	if err != nil {
		return nil, ErrFailedToGetEndpoint
	}

	return resources, nil
}

// AddRoleEndpoint adds a new resource after verifying access permissions.
func (s *accessService) AddRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	err := s.Check(ctx, addRoleEndpointEndpoint)
	if err != nil {
		return err
	}

	err = s.accessRepository.AddRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		if errors.Is(err, ErrEndpointAlreadyExists) {
			return ErrEndpointAlreadyExists
		}
		return ErrFailedToAddEndpoint
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	s.accessibleRoles[endpoint] = roles

	return nil
}

// UpdateRoleEndpoint edits an existing resource after verifying access permissions.
func (s *accessService) UpdateRoleEndpoint(ctx context.Context, endpoint string, roles []string) error {
	err := s.Check(ctx, updateRoleEndpointEndpoint)
	if err != nil {
		return err
	}

	err = s.accessRepository.UpdateRoleEndpoint(ctx, endpoint, roles)
	if err != nil {
		return ErrFailedToUpdateEndpoint
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	s.accessibleRoles[endpoint] = roles

	return nil
}

// DeleteRoleEndpoint deletes a resource after verifying access permissions.
func (s *accessService) DeleteRoleEndpoint(ctx context.Context, endpoint string) error {
	err := s.Check(ctx, deleteRoleEndpointEndpoint)
	if err != nil {
		return err
	}

	err = s.accessRepository.DeleteRoleEndpoint(ctx, endpoint)
	if err != nil {
		return ErrFailedToDeleteEndpoint
	}

	s.rolesMutex.Lock()
	defer s.rolesMutex.Unlock()

	delete(s.accessibleRoles, endpoint)

	return nil
}
