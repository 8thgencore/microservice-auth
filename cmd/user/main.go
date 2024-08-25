package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	pb.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
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
