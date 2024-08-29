package app

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/config"
	"github.com/8thgencore/microservice_auth/internal/delivery/user"
	"github.com/8thgencore/microservice_auth/internal/repository"
	"github.com/8thgencore/microservice_auth/internal/service"
	"github.com/8thgencore/microservice_auth/pkg/closer"
	"github.com/8thgencore/microservice_auth/pkg/db"
	"github.com/8thgencore/microservice_auth/pkg/db/pg"
	"github.com/8thgencore/microservice_auth/pkg/db/transaction"
	"github.com/8thgencore/microservice_auth/pkg/logger"
	"go.uber.org/zap"

	logRepository "github.com/8thgencore/microservice_auth/internal/repository/log"
	userRepository "github.com/8thgencore/microservice_auth/internal/repository/user"
	userService "github.com/8thgencore/microservice_auth/internal/service/user"
)

type serviceProvider struct {
	config *config.Config

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	logRepository  repository.LogRepository

	userService service.UserService

	userImpl *user.UserImplementation
}

func newServiceProvider(config *config.Config) *serviceProvider {
	return &serviceProvider{
		config: config,
	}
}

func (s *serviceProvider) DatabaseClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		c, err := pg.New(ctx, s.config.Database.DSN())
		if err != nil {
			logger.Fatal("failed to create db client: ", zap.Error(err))
		}

		err = c.DB().Ping(ctx)
		if err != nil {
			logger.Fatal("failed to ping database: ", zap.Error(err))
		}

		closer.Add(c.Close)

		s.dbClient = c
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DatabaseClient(ctx).DB())
	}
	return s.txManager
}

// Repository
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.userRepository
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.logRepository
}

// Service
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx), s.LogRepository(ctx), s.TxManager(ctx))
	}
	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.UserImplementation {
	if s.userImpl == nil {
		s.userImpl = user.NewUserImplementation(s.UserService(ctx))
	}
	return s.userImpl
}
