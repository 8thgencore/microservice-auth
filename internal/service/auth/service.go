package auth

import (
	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
)

type authService struct {
	userRepository  repository.UserRepository
	tokenRepository repository.TokenRepository
	tokenOperations tokens.TokenOperations
	jwtConfig       config.JWTConfig
}

// NewService creates new object of service layer.
func NewService(
	userRepository repository.UserRepository,
	tokenRepository repository.TokenRepository,
	tokenOperations tokens.TokenOperations,
	jwtConfig config.JWTConfig,
) service.AuthService {
	return &authService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		tokenOperations: tokenOperations,
		jwtConfig:       jwtConfig,
	}
}
