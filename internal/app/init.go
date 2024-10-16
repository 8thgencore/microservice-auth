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
	"github.com/8thgencore/microservice-auth/internal/metrics"
	"github.com/8thgencore/microservice-auth/internal/tracing"
	accessv1 "github.com/8thgencore/microservice-auth/pkg/access/v1"
	authv1 "github.com/8thgencore/microservice-auth/pkg/auth/v1"
	"github.com/8thgencore/microservice-auth/pkg/swagger"
	userv1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
	"github.com/8thgencore/microservice-common/pkg/closer"
	"github.com/8thgencore/microservice-common/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
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
		a.initPrometheusServer,
		a.initTracing,
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
			interceptor.MetricsInterceptor,
			interceptor.TracingInterceptor,
		),
	)

	reflection.Register(a.grpcServer)

	userv1.RegisterUserV1Server(a.grpcServer, a.serviceProvider.UserImpl(ctx))
	authv1.RegisterAuthV1Server(a.grpcServer, a.serviceProvider.AuthImpl(ctx))
	accessv1.RegisterAccessV1Server(a.grpcServer, a.serviceProvider.AccessImpl(ctx))

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

func (a *App) initPrometheusServer(ctx context.Context) error {
	err := metrics.Init(ctx)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              a.serviceProvider.Config.Prometheus.Address(),
		Handler:           mux,
		ReadHeaderTimeout: 15 * time.Second,
	}

	return nil
}

func (a *App) initTracing(ctx context.Context) error {
	cfg := a.serviceProvider.Config.Tracing

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceName(cfg.ServiceName),
		),
	)
	if err != nil {
		return err
	}

	conn, err := grpc.NewClient(
		cfg.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return err
	}

	// Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return err
	}

	// Register the trace exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter.
	closer.Add(func() error {
		return tracerProvider.Shutdown(ctx)
	})

	err = tracing.InitGlobalTracer(cfg.ServiceName)
	if err != nil {
		return err
	}

	return nil
}
