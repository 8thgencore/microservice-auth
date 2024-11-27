package user

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/converter"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
)

// Create is used for creating new user.
func (impl *Implementation) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	if req == nil || req.GetUser() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: user data is nil")
	}

	id, err := impl.userService.Create(ctx, converter.ToUserCreateFromAPI(req.GetUser()))
	if err != nil {
		return nil,  status.Errorf(codes.Internal, err.Error())
	}

	return &userv1.CreateResponse{
		Id: id,
	}, nil
}

// Get is used for getting user info.
func (impl *Implementation) Get(ctx context.Context, req *userv1.GetRequest) (*userv1.GetResponse, error) {
	if req.GetId() == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request: user ID is required")
    }

	user, err := impl.userService.Get(ctx, req.GetId())
	if err != nil {
	     return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &userv1.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}

// Update is used for updating user info.
func (impl *Implementation) Update(ctx context.Context, req *userv1.UpdateRequest) (*empty.Empty, error) {
	if req == nil || req.GetUser() == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: user data is nil")
	}

	err := impl.userService.Update(ctx, converter.ToUserUpdateFromAPI(req.GetUser()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}

// Delete is used for deleting user.
func (impl *Implementation) Delete(ctx context.Context, req *userv1.DeleteRequest) (*empty.Empty, error) {
    if req.GetId() == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request: user ID is required")
    }

    err := impl.userService.Delete(ctx, req.GetId())
    if err != nil {
        return nil, status.Errorf(codes.NotFound, err.Error())
    }


	return &empty.Empty{}, nil
}
