package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/8thgencore/microservice-auth/internal/app/provider"
	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-common/pkg/closer"
	"github.com/8thgencore/microservice-common/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// App structure contains main application structures.
type App struct {
	cfg *config.Config

	serviceProvider  *provider.ServiceProvider
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
}

// NewApp creates new App object.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}
	return a, nil
}

// Run executes the application.
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(3) // gRPC, HTTP and Swagger servers

	go func() {
		defer wg.Done()

		if err := a.runGrpcServer(); err != nil {
			log.Fatal("failed to run gRPC server: ", error.Error(err))
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runHTTPServer(); err != nil {
			logger.Fatal("failed to run HTTP server: ", zap.Error(err))
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runSwaggerServer(); err != nil {
			logger.Fatal("failed to run Swagger server: ", zap.Error(err))
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			logger.Fatal("failed to run Prometheus server: ", zap.Error(err))
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runGrpcServer() error {
	cfg := a.serviceProvider.Config.GRPC

	logger.Info("gRPC server running on ", zap.String("address", cfg.Address()))

	lis, err := net.Listen(cfg.Transport, cfg.Address())
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (a *App) runHTTPServer() error {
	logger.Info("HTTP server running on ", zap.String("address", a.serviceProvider.Config.HTTP.Address()))

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	logger.Info("Swagger server running on ", zap.String("address", a.serviceProvider.Config.Swagger.Address()+"/docs"))

	if err := a.swaggerServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	logger.Info("Prometheus server running on ", zap.String("address", a.serviceProvider.Config.Prometheus.Address()))

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
