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

	jwtConfig = config.JWTConfig{
		SecretKey:       secretKey,
		AccessTokenTTL:  30 * time.Minute,
		RefreshTokenTTL: 360 * time.Minute,
	}
)

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

		keyRepositoryErr  = fmt.Errorf("failed to generate token")
		userRepositoryErr = fmt.Errorf("user not found")
		wrongPasswordErr  = fmt.Errorf("wrong password")

		authInfo = &model.AuthInfo{
			ID:       userID,
			Username: username,
			Password: string(hashedPassword),
			Role:     role,
		}

		user = model.User{
			ID:   userID,
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
				mock.GenerateAccessTokenMock.Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).Return("", keyRepositoryErr)
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
				mock.GenerateAccessTokenMock.Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).Return(accessToken, nil)
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

		refreshClaims = &model.RefreshClaims{UserID: userID}

		user = model.User{
			Name: username,
			Role: role,
		}

		// userNotFound    = errors.New("user not found")
		generateAccessErr = fmt.Errorf("failed to generate access token")
		// generateRefreshErr   = fmt.Errorf("failed to generate refresh token")
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
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(nil, tokenInvalidErr)
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
			err:  generateAccessErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(&user, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateAccessTokenMock.Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).Return("", generateAccessErr)
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

		refreshClaims = &model.RefreshClaims{UserID: userID}

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
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(nil, tokenInvalidErr)
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
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateRefreshTokenMock.
					Expect(user.ID, []byte(secretKey), jwtConfig.RefreshTokenTTL).
					Return("", repositoryErr)
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
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := authService.NewService(userRepositoryMock, tokenOperationsMock, jwtConfig)

			res, err := srv.GetRefreshToken(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, res)
		})
	}
}