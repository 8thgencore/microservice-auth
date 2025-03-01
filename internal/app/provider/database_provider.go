package provider

import (
	"context"

	"github.com/8thgencore/microservice-common/pkg/cache"
	redisClient "github.com/8thgencore/microservice-common/pkg/cache/redis"
	"github.com/8thgencore/microservice-common/pkg/closer"
	"github.com/8thgencore/microservice-common/pkg/db"
	"github.com/8thgencore/microservice-common/pkg/db/pg"
	"github.com/8thgencore/microservice-common/pkg/db/transaction"
	"github.com/8thgencore/microservice-common/pkg/logger/sl"
	"github.com/redis/go-redis/v9"
)

// DatabaseClient returns a database client.
// If the client has not been created yet, it creates a new one using the DSN from the configuration.
// It also checks if the database is reachable by pinging it.
// The client is closed when the application shuts down.
func (s *ServiceProvider) DatabaseClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		c, err := pg.New(ctx, s.Config.Database.DSN())
		if err != nil {
			s.logger.Error("failed to create db client: ", sl.Err(err))
		}

		err = c.DB().Ping(ctx)
		if err != nil {
			s.logger.Error("failed to ping database: ", sl.Err(err))
		}

		closer.Add(c.Close)

		s.dbClient = c
	}

	return s.dbClient
}

// TxManager returns a transaction manager.
// If the transaction manager has not been created yet, it creates a new one using the database client.
func (s *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DatabaseClient(ctx).DB())
	}
	return s.txManager
}

// CacheClient returns a cache client.
func (s *ServiceProvider) CacheClient(ctx context.Context) cache.Client {
	cfg := s.Config.Redis
	opt := &redis.Options{
		Addr:        cfg.Address(),
		Password:    cfg.Password,
		DB:          0, // use default DB
		DialTimeout: cfg.ConnectionTimeout,
		ReadTimeout: cfg.IdleTimeout,
		PoolSize:    cfg.MaxIdle,
	}

	if s.cache == nil {
		c := redisClient.NewClient(opt, s.logger)

		if err := c.Ping(ctx); err != nil {
			s.logger.Error("failed to connect to redis: ", sl.Err(err))
		}

		s.cache = c
	}

	return s.cache
}
