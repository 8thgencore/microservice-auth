package app

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/8thgencore/microservice_auth/internal/config"
	"github.com/8thgencore/microservice_auth/pkg/logger"
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// App structure contains main application structures.
type App struct {
	cfg *config.Config
}

// NewApp creates new App object.
func NewApp(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Run executes the application.
func (a *App) Run() error {

	wg := sync.WaitGroup{}
	wg.Add(1) // gRPC servers

	go func() {
		defer wg.Done()

		err := a.runGrpcServer()
		if err != nil {
			log.Fatal("failed to run gRPC server: ", error.Error(err))
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initGrpcServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
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

// gRPC

type server struct {
	pb.UnimplementedUserV1Server
}

func (a *App) initGrpcServer(ctx context.Context) error {
	return nil
}

func (a *App) runGrpcServer() error {
	lis, err := net.Listen(a.cfg.GRPC.Transport, a.cfg.GRPC.Address())
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
