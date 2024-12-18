package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	userAPI "github.com/8thgencore/microservice-auth/internal/delivery/user"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/service"
	serviceMocks "github.com/8thgencore/microservice-auth/internal/service/mocks"
	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	id              = "uuid"
	name            = "name"
	email           = "email"
	role            = userv1.Role_USER
	roleName        = "USER"
	password        = "password"
	passwordConfirm = "passwordConfirm"
	createdAt       = timestamppb.Now()
	updatedAt       = timestamppb.Now()
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *userv1.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		serviceErr = errors.New("service error")

		req = &userv1.CreateRequest{
			User: &userv1.UserCreate{
				Name:            name,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirm,
				Role:            role,
			},
		}

		userCreate = &model.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            roleName,
		}

		res = &userv1.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *userv1.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(minimock.AnyContext, userCreate).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.CreateMock.Expect(minimock.AnyContext, userCreate).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewImplementation(userServiceMock)

			res, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGet(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *userv1.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		serviceErr = errors.New("service error")

		req = &userv1.GetRequest{
			Id: id,
		}

		userInfo = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      roleName,
			CreatedAt: createdAt.AsTime(),
			UpdatedAt: sql.NullTime{
				Time:  updatedAt.AsTime(),
				Valid: true,
			},
		}

		res = &userv1.GetResponse{
			User: &userv1.User{
				Id:      id,
				Name:    name,
				Email:   email,
				Role:    role,
				Created: createdAt,
				Updated: updatedAt,
			},
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *userv1.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(minimock.AnyContext, id).Return(userInfo, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.GetMock.Expect(minimock.AnyContext, id).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewImplementation(userServiceMock)

			res, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *userv1.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = "uuid"
		name     = "name"
		email    = "email"
		role     = userv1.Role_USER
		roleName = "USER"

		serviceErr = errors.New("service error")

		req = &userv1.UpdateRequest{
			User: &userv1.UserUpdate{
				Id:    id,
				Name:  wrapperspb.String(name),
				Email: wrapperspb.String(email),
				Role:  role,
			},
		}

		userUpdate = &model.UserUpdate{
			ID:    id,
			Name:  &name,
			Email: &email,
			Role:  &roleName,
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *empty.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(minimock.AnyContext, userUpdate).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.Internal, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(minimock.AnyContext, userUpdate).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewImplementation(userServiceMock)

			res, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *userv1.DeleteRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id = "uuid"

		serviceErr = errors.New("service error")

		req = &userv1.DeleteRequest{
			Id: id,
		}

		res = &empty.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *empty.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(minimock.AnyContext, id).Return(nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.NotFound, serviceErr.Error()),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.DeleteMock.Expect(minimock.AnyContext, id).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := userAPI.NewImplementation(userServiceMock)

			res, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
