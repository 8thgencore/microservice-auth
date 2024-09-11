package access

import (
	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
)

type serv struct {
	accessRepository repository.AccessRepository
	tokenOperations  tokens.TokenOperations
	jwtConfig        config.JWTConfig
}

// NewService creates new object of service layer.
func NewService(
	accessRepository repository.AccessRepository,
	tokenOperations tokens.TokenOperations,
	jwtConfig config.JWTConfig,
) service.AccessService {
	return &serv{
		accessRepository: accessRepository,
		tokenOperations:  tokenOperations,
		jwtConfig:        jwtConfig,
	}
}
