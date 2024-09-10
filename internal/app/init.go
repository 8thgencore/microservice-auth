package app

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/8thgencore/microservice-auth/internal/app/provider"
	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/interceptor"
	"github.com/8thgencore/microservice-auth/pkg/swagger"
	userv1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
	"github.com/8thgencore/microservice-common/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initSwaggerServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	a.cfg = cfg

	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(string(a.cfg.Env))

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = provider.NewServiceProvider(a.cfg)
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	var creds credentials.TransportCredentials
	var err error

	if a.cfg.TLS.Enable {
		creds, err = credentials.NewServerTLSFromFile(a.cfg.TLS.CertPath, a.cfg.TLS.KeyPath)
		if err != nil {
			return err
		}
	} else {
		creds = insecure.NewCredentials()
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			interceptor.LogInterceptor,
			interceptor.ValidateInterceptor,
		),
	)

	reflection.Register(a.grpcServer)

	userv1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	var creds credentials.TransportCredentials
	var err error

	if a.cfg.TLS.Enable {
		creds, err = credentials.NewClientTLSFromFile(a.cfg.TLS.CertPath, "")
		if err != nil {
			return err
		}
	} else {
		creds = insecure.NewCredentials()
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	mux := runtime.NewServeMux()

	if err := userv1.RegisterUserV1HandlerFromEndpoint(ctx, mux, a.cfg.GRPC.Address(), opts); err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:              a.cfg.HTTP.Address(),
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: 15 * time.Second,
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	mux := http.NewServeMux()

	// Helper function to serve embedded Swagger files
	serveFile := func(w http.ResponseWriter, fileName, contentType string) {
		file, err := swagger.SwaggerFiles.Open(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("Error closing file: %v", err)
			}
		}()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", contentType)
		if _, err := w.Write(content); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}

	// Serve Swagger UI HTML
	mux.HandleFunc("/docs", func(w http.ResponseWriter, _ *http.Request) {
		serveFile(w, "index.html", "text/html")
	})

	// Serve Swagger JSON
	mux.HandleFunc("/api.swagger.json", func(w http.ResponseWriter, _ *http.Request) {
		serveFile(w, "api.swagger.json", "application/json")
	})

	a.swaggerServer = &http.Server{
		Addr:              a.serviceProvider.Config.Swagger.Address(),
		Handler:           mux,
		ReadHeaderTimeout: 15 * time.Second,
	}

	return nil
}
