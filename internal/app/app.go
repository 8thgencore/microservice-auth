package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/8thgencore/microservice-auth/internal/app/provider"
	"github.com/8thgencore/microservice-auth/internal/config"
	"google.golang.org/grpc"
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
