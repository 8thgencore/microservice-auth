package auth

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	repositoryMocks "github.com/8thgencore/microservice-auth/internal/repository/mocks"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	tokenMocks "github.com/8thgencore/microservice-auth/internal/tokens/mocks"
)

var (
	userID          = int64(1)
	username        = "username"
	password        = "password"
	passwordWrong   = "passwordWrong"
	role            = "USER"
	refreshToken    = "refresh_token"
	oldRefreshToken = "old_refresh_token"
	accessToken     = "access_token"

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	require.NoError(t, err)

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
			err:  ErrUserNotFound,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(nil, ErrUserNotFound)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
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
			err:  ErrWrongPassword,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
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
			err:  ErrTokenGeneration,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return("", ErrTokenGeneration)
				return mock
			},
		},
		{
			name: "refresh token generate error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  ErrTokenGeneration,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetAuthInfoMock.Expect(minimock.AnyContext, username).Return(authInfo, nil)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return(accessToken, nil)
				mock.GenerateRefreshTokenMock.
					Expect(user.ID, []byte(secretKey), jwtConfig.RefreshTokenTTL).
					Return("", ErrTokenGeneration)
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
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := NewService(
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
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.
					Expect(refreshToken, []byte(secretKey)).
					Return(refreshClaims, nil)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return(accessToken, nil)
				return mock
			},
		},
		{
			name: "refresh token revoked case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(true, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "error checking revoked token case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(false, errors.New("db error"))
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.
					Expect(refreshToken, []byte(secretKey)).
					Return(nil, ErrInvalidRefresh)
				return mock
			},
		},
		{
			name: "get user error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrUserNotFound,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, ErrUserNotFound)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.
					Expect(refreshToken, []byte(secretKey)).
					Return(refreshClaims, nil)
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
			err:  ErrTokenGeneration,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(&user, nil)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, refreshToken).Return(false, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, []byte(secretKey)).Return(refreshClaims, nil)
				mock.GenerateAccessTokenMock.
					Expect(user, []byte(secretKey), jwtConfig.AccessTokenTTL).
					Return("", ErrTokenGeneration)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := NewService(
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
		{
			name: "error checking revoked token case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrInvalidRefresh,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, oldRefreshToken).Return(false, errors.New("db error"))
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "token verify error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrInvalidRefresh,
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
				mock.VerifyRefreshTokenMock.Expect(oldRefreshToken, []byte(secretKey)).Return(nil, ErrInvalidRefresh)
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
			err:  ErrTokenGeneration,
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
					Return("", ErrTokenGeneration)
				return mock
			},
		},
		{
			name: "add revoked token error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: "",
			err:  ErrTokenGeneration,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repositoryMocks.NewUserRepositoryMock(mc)
				return mock
			},
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.IsTokenRevokedMock.Expect(ctx, oldRefreshToken).Return(false, nil)
				mock.AddRevokedTokenMock.Expect(ctx, oldRefreshToken).Return(ErrTokenGeneration)
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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepositoryMock := tt.userRepositoryMock(mc)
			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := NewService(
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

func TestLogout(t *testing.T) {
	t.Parallel()

	type tokenRepositoryMockFunc func(mc *minimock.Controller) repository.TokenRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx          context.Context
		refreshToken string
	}

	var (
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		refreshToken = "validRefreshToken"
		secretKey    = []byte("secret")
	)

	tests := []struct {
		name                string
		args                args
		err                 error
		tokenRepositoryMock tokenRepositoryMockFunc
		tokenOperationsMock tokenOperationsMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			err: nil,
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.AddRevokedTokenMock.Expect(ctx, refreshToken).Return(nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, secretKey).Return(&model.RefreshClaims{}, nil)
				return mock
			},
		},
		{
			name: "invalid refresh token case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			err: ErrInvalidRefresh,
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, secretKey).Return(nil, ErrInvalidRefresh)
				return mock
			},
		},
		{
			name: "token repository failure case",
			args: args{
				ctx:          ctx,
				refreshToken: refreshToken,
			},
			err: ErrLogoutFailed,
			tokenRepositoryMock: func(mc *minimock.Controller) repository.TokenRepository {
				mock := repositoryMocks.NewTokenRepositoryMock(mc)
				mock.AddRevokedTokenMock.Expect(ctx, refreshToken).Return(errors.New("db error"))
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyRefreshTokenMock.Expect(refreshToken, secretKey).Return(&model.RefreshClaims{}, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tokenRepositoryMock := tt.tokenRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := NewService(
				nil, // userRepository не используется в этом методе
				tokenRepositoryMock,
				tokenOperationsMock,
				jwtConfig,
			)

			err := srv.Logout(tt.args.ctx, tt.args.refreshToken)
			require.Equal(t, tt.err, err)
		})
	}
}