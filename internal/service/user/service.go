package user

import (
	"context"
	"log/slog"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	"github.com/8thgencore/microservice-common/pkg/db"
)

type userService struct {
	logger          *slog.Logger
	userRepository  repository.UserRepository
	logRepository   repository.LogRepository
	tokenRepository repository.TokenRepository
	tokenOperations tokens.TokenOperations
	txManager       db.TxManager
	adminConfig     *config.AdminConfig
}

// NewService creates new object of service layer and ensures admin user exists.
func NewService(
	logger *slog.Logger,
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	tokenRepository repository.TokenRepository,
	tokenOperations tokens.TokenOperations,
	txManager db.TxManager,
	adminConfig *config.AdminConfig,
) service.UserService {
	s := &userService{
		logger:          logger,
		userRepository:  userRepository,
		logRepository:   logRepository,
		tokenRepository: tokenRepository,
		tokenOperations: tokenOperations,
		txManager:       txManager,
		adminConfig:     adminConfig,
	}

	// Ensure admin exists during service initialization
	if err := s.EnsureAdminExists(context.Background()); err != nil {
		// Log the error but don't fail the service initialization
		s.logger.Error("failed to ensure admin exists", slog.String("error", err.Error()))
	}

	return s
}

// newTestService creates service instance without admin check (for testing only)
func newTestService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	tokenRepository repository.TokenRepository,
	tokenOperations tokens.TokenOperations,
	txManager db.TxManager,
	adminConfig *config.AdminConfig,
) service.UserService {
	return &userService{
		userRepository:  userRepository,
		logRepository:   logRepository,
		tokenRepository: tokenRepository,
		tokenOperations: tokenOperations,
		txManager:       txManager,
		adminConfig:     adminConfig,
	}
}
