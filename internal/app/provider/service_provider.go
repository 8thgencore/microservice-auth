package provider

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/delivery/user"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-common/pkg/db"

	logRepository "github.com/8thgencore/microservice-auth/internal/repository/log"
	userRepository "github.com/8thgencore/microservice-auth/internal/repository/user"
	userService "github.com/8thgencore/microservice-auth/internal/service/user"
)

type ServiceProvider struct {
	Config *config.Config

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	logRepository  repository.LogRepository

	userService service.UserService

	userImpl *user.UserImplementation
}

func NewServiceProvider(config *config.Config) *ServiceProvider {
	return &ServiceProvider{
		Config: config,
	}
}

// Repository
func (s *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.userRepository
}

func (s *ServiceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.logRepository
}

// Service
func (s *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx), s.LogRepository(ctx), s.TxManager(ctx))
	}
	return s.userService
}

func (s *ServiceProvider) UserImpl(ctx context.Context) *user.UserImplementation {
	if s.userImpl == nil {
		s.userImpl = user.NewUserImplementation(s.UserService(ctx))
	}
	return s.userImpl
}
