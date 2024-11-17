package user

import (
	"context"
	"fmt"

	"github.com/8thgencore/microservice-auth/internal/converter"
	"github.com/golang/protobuf/ptypes/empty"

	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
)

// Create is used for creating new user.
func (impl *Implementation) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	if req == nil || req.GetUser() == nil {
		return nil, fmt.Errorf("invalid request: user data is nil")
	}

	id, err := impl.userService.Create(ctx, converter.ToUserCreateFromAPI(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &userv1.CreateResponse{
		Id: id,
	}, nil
}

// Get is used for getting user info.
func (impl *Implementation) Get(ctx context.Context, req *userv1.GetRequest) (*userv1.GetResponse, error) {
	user, err := impl.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &userv1.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}

// Update is used for updating user info.
func (impl *Implementation) Update(ctx context.Context, req *userv1.UpdateRequest) (*empty.Empty, error) {
	if req == nil || req.GetUser() == nil {
		return nil, fmt.Errorf("invalid request: user data is nil")
	}

	err := impl.userService.Update(ctx, converter.ToUserUpdateFromAPI(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// Delete is used for deleting user.
func (impl *Implementation) Delete(ctx context.Context, req *userv1.DeleteRequest) (*empty.Empty, error) {
	err := impl.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
