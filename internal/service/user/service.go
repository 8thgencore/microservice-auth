package user

import (
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	"github.com/8thgencore/microservice-common/pkg/db"
)

type serv struct {
	userRepository  repository.UserRepository
	logRepository   repository.LogRepository
	tokenOperations tokens.TokenOperations
	txManager       db.TxManager
}

// NewService creates new object of service layer.
func NewService(
	userRepository repository.UserRepository,
	logRepository repository.LogRepository,
	tokenOperations tokens.TokenOperations,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository:  userRepository,
		logRepository:   logRepository,
		tokenOperations: tokenOperations,
		txManager:       txManager,
	}
}
