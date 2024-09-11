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

var (
	userID          int64 = 1
	username              = "username"
	passwordWrong         = "passwordWrong"
	role                  = "USER"
	refreshToken          = "refresh_token"
	oldRefreshToken       = "old_refresh_token"
	accessToken           = "access_token"

	secretKey = "secret"

	user = model.User{
		ID:   userID,
		Name: username,
		Role: role,
	}

	jwtConfig = config.JWTConfig{
		SecretKey:       secretKey,
		AccessTokenTTL:  30 * time.Minute,
		RefreshTokenTTL: 360 * time.Minute,
	}
)

func TestLogin(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenRepositoryMockFunc func(mc *minimock.Controller) repository.TokenRepository
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

		authInfo = &model.AuthInfo{
			ID:       userID,
			Username: username,
			Password: string(hashedPassword),
			Role:     role,
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
		tokenRepositoryMock tokenRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "user repository error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  authService.ErrUserNotFound,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(nil, authService.ErrUserNotFound)
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
			want: nil,
			err:  authService.ErrWrongPassword,
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
			want: nil,
			err:  authService.ErrAccessTokenGen,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return("", authService.ErrAccessTokenGen)
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
				// Expect Generate to be called for access token
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return(accessToken, nil)
				// Expect Generate to be called for refresh token
				mock.GenerateRefreshTokenMock.
					Expect(user.ID, []byte(secretKey), jwtConfig.RefreshTokenTTL).
					Return(refreshToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(
				userRepositoryMock,
				tokenRepositoryMock,
				tokenOperationsMock,
				jwtConfig,
			)

			res, err := srv.Login(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetAccessToken(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenRepositoryMockFunc func(mc *minimock.Controller) repository.TokenRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshClaims = &model.RefreshClaims{UserID: userID}

		req = refreshToken
		res = accessToken
	)

	tests := []struct {
		name                string
		args                args
		want                string
		err                 error
		userRepositoryMock  userRepositoryMockFunc
		tokenRepositoryMock tokenRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  authService.ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(nil, authService.ErrInvalidRefresh)
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
			err:  authService.ErrAccessTokenGen,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(&user, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return("", authService.ErrAccessTokenGen)
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
				mock.GetMock.Expect(ctx, userID).Return(&user, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateAccessTokenMock.Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).Return(accessToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(
				userRepositoryMock,
				tokenRepositoryMock,
				tokenOperationsMock,
				jwtConfig,
			)
			res, err := srv.GetAccessToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}

func TestGetRefreshToken(t *testing.T) {
	t.Parallel()

	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type tokenRepositoryMockFunc func(mc *minimock.Controller) repository.TokenRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		refreshClaims = &model.RefreshClaims{UserID: userID}

		req = oldRefreshToken
		res = refreshToken
	)

	tests := []struct {
		name                string
		args                args
		want                string
		err                 error
		userRepositoryMock  userRepositoryMockFunc
		tokenRepositoryMock tokenRepositoryMockFunc

		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  authService.ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, oldRefreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(nil, authService.ErrInvalidRefresh)
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
			err:  authService.ErrRefreshTokenGen,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, oldRefreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateRefreshTokenMock.
					Expect(user.ID, []byte(secretKey), jwtConfig.RefreshTokenTTL).
					Return("", authService.ErrRefreshTokenGen)
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
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, oldRefreshToken).Return(false, nil)
				mock.AddRevokedTokenMock.Expect(ctx, oldRefreshToken).Return(nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateRefreshTokenMock.
					Expect(user.ID, []byte(secretKey), jwtConfig.RefreshTokenTTL).
					Return(refreshToken, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(
				userRepositoryMock,
				tokenRepositoryMock,
				tokenOperationsMock,
				jwtConfig,
			)
			res, err := srv.GetRefreshToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}
