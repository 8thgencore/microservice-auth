package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	authAPI "github.com/8thgencore/microservice-auth/internal/delivery/auth"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/service"
	serviceMocks "github.com/8thgencore/microservice-auth/internal/service/mocks"
	auth_v1 "github.com/8thgencore/microservice-auth/pkg/pb/auth/v1"
)

var (
	username     = "username"
	password     = "password"
	refreshToken = "refresh_token"
	accessToken  = "access_token"
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

		serviceErr = errors.New("service error")

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
			AccessToken:  accessToken,
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
				mock.LoginMock.Expect(minimock.AnyContext, creds).Return(
					&model.TokenPair{AccessToken: accessToken, RefreshToken: refreshToken}, nil,
				)
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
				mock.LoginMock.Expect(minimock.AnyContext, creds).Return(&model.TokenPair{}, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
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

func TestRefreshTokens(t *testing.T) {
	t.Parallel()

	type authServiceMockFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx context.Context
		req *auth_v1.RefreshTokensRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshToken    = "refresh_token"
		accessToken     = "access_token"
		oldRefreshToken = "old_refresh_token"

		serviceErr = errors.New("service error")

		req = &auth_v1.RefreshTokensRequest{
			RefreshToken: oldRefreshToken,
		}

		res = &auth_v1.RefreshTokensResponse{
			RefreshToken: refreshToken,
			AccessToken:  accessToken,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *auth_v1.RefreshTokensResponse
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
				mock.GetAccessTokenMock.Expect(minimock.AnyContext, oldRefreshToken).
					Return(accessToken, nil)
				mock.GetRefreshTokenMock.Expect(minimock.AnyContext, oldRefreshToken).
					Return(refreshToken, nil)

				return mock
			},
		},
		{
			name: "access token error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetAccessTokenMock.Expect(minimock.AnyContext, oldRefreshToken).
					Return("", serviceErr)
				return mock
			},
		},
		{
			name: "refresh token error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetAccessTokenMock.Expect(minimock.AnyContext, oldRefreshToken).
					Return(accessToken, nil)
				mock.GetRefreshTokenMock.Expect(minimock.AnyContext, oldRefreshToken).
					Return("", serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := authAPI.NewImplementation(authServiceMock)

			res, err := api.RefreshTokens(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
