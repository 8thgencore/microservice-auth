package access

import (
	"context"
	"sync"

	"github.com/8thgencore/microservice-auth/internal/converter"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
)

type accessService struct {
	accessRepository repository.AccessRepository
	tokenOperations  tokens.TokenOperations
	accessibleRoles  map[string][]string
	rolesMutex       sync.RWMutex
}

// NewService creates new object of service layer.
func NewService(
	ctx context.Context,
	accessRepository repository.AccessRepository,
	tokenOperations tokens.TokenOperations,
) (service.AccessService, error) {
	endpointPermissions, err := accessRepository.GetRoleEndpoints(ctx)
	if err != nil {
		return nil, ErrFailedToReadAccessPolicy
	}
	accessibleRoles := converter.ToEndpointPermissionsMap(endpointPermissions)

	return &accessService{
		accessRepository: accessRepository,
		tokenOperations:  tokenOperations,
		accessibleRoles:  accessibleRoles,
	}, nil
}
