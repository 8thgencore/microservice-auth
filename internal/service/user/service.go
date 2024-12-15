package user

import (
	"context"
	"log/slog"

	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	"github.com/8thgencore/microservice-common/pkg/db"
	"github.com/8thgencore/microservice-common/pkg/logger"
)

type serv struct {
	userRepository  repository.UserRepository
	logRepository   repository.LogRepository
	tokenRepository repository.TokenRepository
	tokenOperations tokens.TokenOperations
	txManager       db.TxManager
}

// NewService creates new object of service layer and ensures admin user exists.
func NewService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	tokenRepository repository.TokenRepository,
	tokenOperations tokens.TokenOperations,
	txManager db.TxManager,
) service.UserService {
	s := &serv{
		userRepository:  userRepository,
		logRepository:   logRepository,
		tokenRepository: tokenRepository,
		tokenOperations: tokenOperations,
		txManager:       txManager,
	}

	// Ensure admin exists during service initialization
	if err := s.EnsureAdminExists(context.Background()); err != nil {
		// Log the error but don't fail the service initialization
		logger.Error("failed to ensure admin exists", slog.String("error", err.Error()))
	}

	return s
}
