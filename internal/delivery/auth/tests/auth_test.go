package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	authAPI "github.com/8thgencore/microservice-auth/internal/delivery/auth"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/service"
	serviceMocks "github.com/8thgencore/microservice-auth/internal/service/mocks"
	auth_v1 "github.com/8thgencore/microservice-auth/pkg/auth/v1"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *auth_v1.LoginRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		username     = "username"
		password     = "password"
		refreshToken = "refresh_token"

		serviceErr = fmt.Errorf("service error")

		req = &auth_v1.LoginRequest{
			Creds: &auth_v1.Creds{
				Username: username,
				Password: password,
			},
		}

		creds = &model.UserCreds{
			Username: username,
			Password: password,
		}

		res = &auth_v1.LoginResponse{
			RefreshToken: refreshToken,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *auth_v1.LoginResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(minimock.AnyContext, creds).Return(refreshToken, nil)
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
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.LoginMock.Expect(minimock.AnyContext, creds).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := authAPI.NewImplementation(authServiceMock)

			res, err := api.Login(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetAccessToken(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *auth_v1.GetAccessTokenRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshToken = "refresh_token"
		accessToken  = "access_token"

		serviceErr = fmt.Errorf("service error")

		req = &auth_v1.GetAccessTokenRequest{
			RefreshToken: refreshToken,
		}

		res = &auth_v1.GetAccessTokenResponse{
			AccessToken: accessToken,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *auth_v1.GetAccessTokenResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetAccessTokenMock.Expect(minimock.AnyContext, refreshToken).Return(accessToken, nil)
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
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetAccessTokenMock.Expect(minimock.AnyContext, refreshToken).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := authAPI.NewImplementation(authServiceMock)

			res, err := api.GetAccessToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetRefreshToken(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *auth_v1.GetRefreshTokenRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		oldRefreshToken = "old_refresh_token"
		refreshToken    = "refresh_token"

		serviceErr = fmt.Errorf("service error")

		req = &auth_v1.GetRefreshTokenRequest{
			OldRefreshToken: oldRefreshToken,
		}

		res = &auth_v1.GetRefreshTokenResponse{
			RefreshToken: refreshToken,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *auth_v1.GetRefreshTokenResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetRefreshTokenMock.Expect(minimock.AnyContext, oldRefreshToken).Return(refreshToken, nil)
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
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetRefreshTokenMock.Expect(minimock.AnyContext, oldRefreshToken).Return("", serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := authAPI.NewImplementation(authServiceMock)

			res, err := api.GetRefreshToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
