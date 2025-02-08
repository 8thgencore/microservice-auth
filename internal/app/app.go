package app

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync"

	"github.com/8thgencore/microservice-auth/internal/app/provider"
	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-common/pkg/closer"
	"github.com/8thgencore/microservice-common/pkg/logger/sl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// App structure contains main application structures.
type App struct {
	cfg    *config.Config
	logger *slog.Logger

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
			a.logger.Error("failed to run gRPC server: ", sl.Err(err))
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runHTTPServer(); err != nil {
			a.logger.Error("failed to run HTTP server: ", sl.Err(err))
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runSwaggerServer(); err != nil {
			a.logger.Error("failed to run Swagger server: ", sl.Err(err))
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			a.logger.Error("failed to run Prometheus server: ", sl.Err(err))
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) runGrpcServer() error {
	cfg := a.serviceProvider.Config.GRPC

	a.logger.Info("gRPC server running on ", slog.String("address", cfg.Address()))

	grpclog.SetLoggerV2(grpclog.NewLoggerV2(io.Discard, io.Discard, io.Discard))

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
	a.logger.Info("HTTP server running on ", slog.String("address", a.serviceProvider.Config.HTTP.Address()))

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	a.logger.Info("Swagger server running on ", slog.String("address", a.serviceProvider.Config.Swagger.Address()+"/docs"))

	if err := a.swaggerServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	a.logger.Info("Prometheus server running on ", slog.String("address", a.serviceProvider.Config.Prometheus.Address()))

	if err := a.prometheusServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
