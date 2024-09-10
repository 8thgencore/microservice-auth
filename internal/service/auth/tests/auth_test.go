package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	repositoryMocks "github.com/8thgencore/microservice-auth/internal/repository/mocks"
	authService "github.com/8thgencore/microservice-auth/internal/service/auth"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	tokenMocks "github.com/8thgencore/microservice-auth/internal/tokens/mocks"
)

var secretKey string = "secret"

var jwtConfig = config.JWTConfig{
	SecretKey:       secretKey,
	AccessTokenTTL:  time.Duration(30 * time.Minute),
	RefreshTokenTTL: time.Duration(360 * time.Minute),
}

func TestLogin(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req *model.UserCreds
	}

	password := "password"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print("failed to process password")
		return
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		username               = "username"
		passwordWrong          = "passwordWrong"
		role                   = "USER"
		refreshTokenExpiration = 360 * time.Minute
		refreshToken           = "refresh_token"
		accessToken            = "access_token"

		keyRepositoryErr  = fmt.Errorf("failed to generate token")
		userRepositoryErr = fmt.Errorf("user not found")
		wrongPasswordErr  = fmt.Errorf("wrong password")

		authInfo = &model.AuthInfo{
			Username: username,
			Password: string(hashedPassword),
			Role:     role,
		}

		user = model.User{
			Name: username,
			Role: role,
		}

		req = &model.UserCreds{
			Username: username,
			Password: password,
		}

		reqWrongPass = &model.UserCreds{
			Username: username,
			Password: passwordWrong,
		}

		res = &model.TokenPair{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                *model.TokenPair
		err                 error
		userRepositoryMock  userRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "user repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &model.TokenPair{},
			err:  userRepositoryErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(nil, userRepositoryErr)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "wrong password error case",
			args: args{
				ctx: ctx,
				req: reqWrongPass,
			},
			want: &model.TokenPair{},
			err:  wrongPasswordErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "token generate error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: &model.TokenPair{},
			err:  keyRepositoryErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.GenerateMock.Expect(user, []byte(secretKey), refreshTokenExpiration).Return("", keyRepositoryErr)
				return mock
			},
		},
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.GenerateMock.Expect(user, []byte(secretKey), refreshTokenExpiration).Return(refreshToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(userRepositoryMock, tokenOperationsMock, jwtConfig)

			res, err := srv.Login(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetAccessToken(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshToken          = "refresh_token"
		accessToken           = "access_token"
		accessTokenExpiration = 30 * time.Minute

		username = "username"
		role     = "USER"

		claims = &model.UserClaims{
			Username: username,
			Role:     role,
		}

		user = model.User{
			Name: username,
			Role: role,
		}

		repositoryErr   = fmt.Errorf("failed to generate token")
		tokenInvalidErr = fmt.Errorf("invalid refresh token")

		req = refreshToken
		res = accessToken
	)

	tests := []struct {
		name                string
		args                args
		want                string
		err                 error
		userRepositoryMock  userRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  tokenInvalidErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(refreshToken, []byte(secretKey)).Return(nil, tokenInvalidErr)
				return mock
			},
		},
		{
			name: "token generate error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  repositoryErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(refreshToken, []byte(secretKey)).Return(claims, nil)
				mock.GenerateMock.Expect(user, []byte(secretKey), accessTokenExpiration).Return("", repositoryErr)
				return mock
			},
		},
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(refreshToken, []byte(secretKey)).Return(claims, nil)
				mock.GenerateMock.Expect(user, []byte(secretKey), accessTokenExpiration).Return(accessToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(userRepositoryMock, tokenOperationsMock, jwtConfig)

			res, err := srv.GetAccessToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetRefreshToken(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshTokenExpiration = 360 * time.Minute
		oldRefreshToken        = "old_refresh_token"
		refreshToken           = "refresh_token"

		username = "username"
		role     = "USER"

		claims = &model.UserClaims{
			Username: username,
			Role:     role,
		}

		user = model.User{
			Name: username,
			Role: role,
		}

		repositoryErr   = fmt.Errorf("failed to generate token")
		tokenInvalidErr = fmt.Errorf("invalid refresh token")

		req = oldRefreshToken
		res = refreshToken
	)

	tests := []struct {
		name                string
		args                args
		want                string
		err                 error
		userRepositoryMock  userRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  tokenInvalidErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(oldRefreshToken, []byte(secretKey)).Return(nil, tokenInvalidErr)
				return mock
			},
		},
		{
			name: "token generate error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  repositoryErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(oldRefreshToken, []byte(secretKey)).Return(claims, nil)
				mock.GenerateMock.Expect(user, []byte(secretKey), refreshTokenExpiration).Return("", repositoryErr)
				return mock
			},
		},
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyMock.Expect(oldRefreshToken, []byte(secretKey)).Return(claims, nil)
				mock.GenerateMock.Expect(user, []byte(secretKey), refreshTokenExpiration).Return(refreshToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(userRepositoryMock, tokenOperationsMock, jwtConfig)

			res, err := srv.GetRefreshToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
