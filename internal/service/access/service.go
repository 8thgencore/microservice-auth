package access

import (
	"sync"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
)

type accessService struct {
	accessRepository repository.AccessRepository
	tokenOperations  tokens.TokenOperations
	jwtConfig        config.JWTConfig
	accessibleRoles  map[string][]string
	rolesMutex       sync.RWMutex
}

// NewService creates new object of service layer.
func NewService(
	accessRepository repository.AccessRepository,
	tokenOperations tokens.TokenOperations,
	jwtConfig config.JWTConfig,
) service.AccessService {
	return &accessService{
		accessRepository: accessRepository,
		tokenOperations:  tokenOperations,
		jwtConfig:        jwtConfig,
	}
}
