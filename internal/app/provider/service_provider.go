package provider

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/delivery/access"
	"github.com/8thgencore/microservice-auth/internal/delivery/auth"
	"github.com/8thgencore/microservice-auth/internal/delivery/user"
	"github.com/8thgencore/microservice-auth/internal/interceptor"
	"github.com/8thgencore/microservice-auth/internal/repository"
	"github.com/8thgencore/microservice-auth/internal/service"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	"github.com/8thgencore/microservice-auth/internal/tokens/jwt"
	"github.com/8thgencore/microservice-common/pkg/cache"
	"github.com/8thgencore/microservice-common/pkg/db"
	"github.com/8thgencore/microservice-common/pkg/logger"
	"go.uber.org/zap"

	accessRepository "github.com/8thgencore/microservice-auth/internal/repository/access"
	logRepository "github.com/8thgencore/microservice-auth/internal/repository/log"
	tokenRepository "github.com/8thgencore/microservice-auth/internal/repository/token"
	userRepository "github.com/8thgencore/microservice-auth/internal/repository/user"
	accessService "github.com/8thgencore/microservice-auth/internal/service/access"
	authService "github.com/8thgencore/microservice-auth/internal/service/auth"
	userService "github.com/8thgencore/microservice-auth/internal/service/user"
)

// ServiceProvider is a struct that provides access to various services and repositories.
type ServiceProvider struct {
	Config *config.Config

	dbClient  db.Client
	txManager db.TxManager

	authInterceptor *interceptor.Auth

	cache cache.Client

	userRepository   repository.UserRepository
	accessRepository repository.AccessRepository
	logRepository    repository.LogRepository
	tokenRepository  repository.TokenRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation

	tokenOperations tokens.TokenOperations
}

// NewServiceProvider creates a new instance of ServiceProvider with the given configuration.
func NewServiceProvider(config *config.Config) *ServiceProvider {
	return &ServiceProvider{
		Config: config,
	}
}

// UserRepository returns a user repository.
func (s *ServiceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.userRepository
}

// AccessRepository returns a access repository.
func (s *ServiceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.accessRepository
}

// LogRepository returns a log repository.
func (s *ServiceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DatabaseClient(ctx))
	}
	return s.logRepository
}

// TokenRepository returns a Token repository.
func (s *ServiceProvider) TokenRepository(ctx context.Context) repository.TokenRepository {
	if s.tokenRepository == nil {
		s.tokenRepository = tokenRepository.NewRepository(
			s.CacheClient(ctx),
			s.Config.JWT.AccessTokenTTL,
			s.Config.JWT.RefreshTokenTTL,
		)
	}

	return s.tokenRepository
}

// UserService returns a user service.
func (s *ServiceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.LogRepository(ctx),
			s.TokenRepository(ctx),
			s.TokenOperations(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

// AuthService returns a auth service.
func (s *ServiceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserRepository(ctx),
			s.TokenRepository(ctx),
			s.TokenOperations(ctx),
		)
	}

	return s.authService
}

// AccessService returns an access service.
func (s *ServiceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		var err error
		s.accessService, err = accessService.NewService(
			ctx,
			s.AccessRepository(ctx),
			s.TokenOperations(ctx),
		)
		if err != nil {
			logger.Fatal("failed to run access service: ", zap.Error(err))
		}
	}

	return s.accessService
}

// UserImpl returns a user implementation.
func (s *ServiceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}
	return s.userImpl
}

// AuthImpl returns a auth implementation.
func (s *ServiceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}
	return s.authImpl
}

// AccessImpl returns a access implementation.
func (s *ServiceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}
	return s.accessImpl
}

// TokenOperations returns a token operation service.
func (s *ServiceProvider) TokenOperations(ctx context.Context) tokens.TokenOperations {
	if s.tokenOperations == nil {
		s.tokenOperations = jwt.NewTokenOperations(
			[]byte(s.Config.JWT.SecretKey),
			s.Config.JWT.AccessTokenTTL,
			s.Config.JWT.RefreshTokenTTL,
			s.TokenRepository(ctx),
		)
	}

	return s.tokenOperations
}

// AuthInterceptorFactory returns an instance of interceptor.Auth.
func (s *ServiceProvider) AuthInterceptorFactory(ctx context.Context) *interceptor.Auth {
	if s.authInterceptor == nil {
		s.authInterceptor = &interceptor.Auth{
			TokenOperations: s.TokenOperations(ctx),
			TokenRepository: s.TokenRepository(ctx),
		}
	}

	return s.authInterceptor
}
