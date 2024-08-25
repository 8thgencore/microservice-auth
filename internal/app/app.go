package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/8thgencore/microservice_auth/internal/config"
	"github.com/8thgencore/microservice_auth/internal/config/env"
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// App structure contains main application structures.
type App struct {
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "Path to config file")
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
	wg.Add(1) // gRPC, HTTP and Swagger servers

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
	// Parse the command-line flags from os.Args[1:].
	flag.Parse()

	err := config.Load(configPath)
	if err != nil {
		log.Fatal("failed to load config: ", error.Error(err))
	}
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
	cfg, err := env.NewGrpcConfig()
	if err != nil {
		log.Fatal("failed to get grpc config: ", error.Error(err))
	}

	lis, err := net.Listen(cfg.Transport(), cfg.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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

func (s *server) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Printf("Create user: %+v\n", req)
	return &pb.CreateUserResponse{Id: 1}, nil
}

func (s *server) Get(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	fmt.Printf("Get user: %d\n", req.GetId())
	return &pb.GetUserResponse{
		Id:        req.GetId(),
		Name:      "John Doe",
		Email:     "john.doe@example.com",
		Role:      pb.Role_USER,
		CreatedAt: &timestamp.Timestamp{Seconds: 1623855600},
		UpdatedAt: &timestamp.Timestamp{Seconds: 1623855600},
	}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.Empty, error) {
	fmt.Printf("Update user: %+v\n", req)
	return &pb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *pb.DeleteUserRequest) (*pb.Empty, error) {
	fmt.Printf("Delete user: %d\n", req.GetId())
	return &pb.Empty{}, nil
}
