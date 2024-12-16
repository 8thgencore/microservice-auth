package user

import (
	"context"
	"errors"

	"github.com/8thgencore/microservice-auth/internal/converter"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/service/user"
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
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.NotFound, err.Error())
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
		return nil, status.Error(codes.Internal, err.Error())
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
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &empty.Empty{}, nil
}

// GetMe returns information about the currently authenticated user
func (impl *Implementation) GetMe(ctx context.Context, _ *empty.Empty) (*userv1.GetMeResponse, error) {
	// Get user ID from context (assuming it was set by auth middleware)
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	user, err := impl.userService.Get(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user info: %v", err)
	}

	return &userv1.GetMeResponse{
		User: converter.ToUserFromService(user),
	}, nil
}

// UpdateMe handles the current user's information update
func (impl *Implementation) UpdateMe(ctx context.Context, req *userv1.UpdateMeRequest) (*empty.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	// Get user ID from context
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	// Convert request to UserUpdate model
	update := &model.UserUpdate{
		ID: userID,
	}
	if req.Name != nil {
		name := req.Name.Value
		update.Name = &name
	}
	if req.Email != nil {
		email := req.Email.Value
		update.Email = &email
	}

	err := impl.userService.Update(ctx, update)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, user.ErrUserNotFound.Error())
		case errors.Is(err, user.ErrUserNameExists):
			return nil, status.Error(codes.AlreadyExists, user.ErrUserNameExists.Error())
		case errors.Is(err, user.ErrUserEmailExists):
			return nil, status.Error(codes.AlreadyExists, user.ErrUserEmailExists.Error())
		default:
			return nil, status.Error(codes.Internal, user.ErrUserUpdate.Error())
		}
	}

	return &empty.Empty{}, nil
}

// DeleteMe handles the deletion of the current user's account
func (impl *Implementation) DeleteMe(ctx context.Context, _ *empty.Empty) (*empty.Empty, error) {
	// Get user ID from context
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	err := impl.userService.Delete(ctx, userID)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, user.ErrUserNotFound.Error())
		default:
			return nil, status.Error(codes.Internal, user.ErrUserDelete.Error())
		}
	}

	return &empty.Empty{}, nil
}

// ChangePassword handles password change requests
func (impl *Implementation) ChangePassword(
	ctx context.Context, req *userv1.ChangePasswordRequest,
) (*empty.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	// Get user ID from context (assuming it was set by auth middleware)
	userID, ok := ctx.Value(UserIDKey).(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user not authenticated")
	}

	err := impl.userService.ChangePassword(ctx, userID, req.CurrentPassword, req.NewPassword)
	if err != nil {
		switch {
		case errors.Is(err, user.ErrInvalidCurrentPassword):
			return nil, status.Error(codes.InvalidArgument, user.ErrInvalidCurrentPassword.Error())
		case errors.Is(err, user.ErrUserNotFound):
			return nil, status.Error(codes.NotFound, user.ErrUserNotFound.Error())
		default:
			return nil, status.Error(codes.Internal, user.ErrUserChangePassword.Error())
		}
	}

	return &empty.Empty{}, nil
}
