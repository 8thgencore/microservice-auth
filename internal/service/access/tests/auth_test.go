package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	repositoryMocks "github.com/8thgencore/microservice-auth/internal/repository/mocks"
	accessService "github.com/8thgencore/microservice-auth/internal/service/access"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	tokenMocks "github.com/8thgencore/microservice-auth/internal/tokens/mocks"
)

var jwtConfig = config.JWTConfig{
	SecretKey:       "secret",
	AccessTokenTTL:  time.Duration(30 * time.Minute),
	RefreshTokenTTL: time.Duration(360 * time.Minute),
}

func TestCheck(t *testing.T) {
	t.Parallel()

	type accessRepositoryMockFunc func(mc *minimock.Controller) repository.AccessRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	type args struct {
		ctx context.Context
		req string
	}

	var (
		mdNoAuthHeader = metadata.New(map[string]string{"header": "access_token"})
		mdNoAuthPrefix = metadata.New(map[string]string{"Authorization": "access_token"})
		md             = metadata.New(map[string]string{"Authorization": "Bearer access_token"})

		ctxNoMd         = context.Background()
		ctx             = metadata.NewIncomingContext(ctxNoMd, md)
		ctxNoAuthHeader = metadata.NewIncomingContext(ctxNoMd, mdNoAuthHeader)
		ctxNoAuthPrefix = metadata.NewIncomingContext(ctxNoMd, mdNoAuthPrefix)

		mc = minimock.NewController(t)

		endpointCreate      = "/chat_v1.ChatV1/Create"
		endpointDelete      = "/chat_v1.ChatV1/Delete"
		endpointSendMessage = "/chat_v1.ChatV1/SendMessage"
		endpointNotExists   = "/chat_v1.ChatV1/NotExists"

		username  = "username"
		roleUser  = "USER"
		roleAdmin = "ADMIN"

		secretKeyBytes = []byte("secret")

		accessToken = "access_token"

		endpointPermissions = []*model.EndpointPermissions{
			{
				Endpoint: endpointCreate,
				Roles:    []string{roleAdmin},
			},
			{
				Endpoint: endpointDelete,
				Roles:    []string{roleAdmin},
			},
			{
				Endpoint: endpointSendMessage,
				Roles:    []string{roleAdmin, roleUser},
			},
		}

		claimsAdmin = &model.UserClaims{
			Username: username,
			Role:     roleAdmin,
		}

		claimsUser = &model.UserClaims{
			Username: username,
			Role:     roleUser,
		}

		noMdErr         = fmt.Errorf("metadata is not provided")
		noAuthHeaderErr = fmt.Errorf("authorization header is not provided")
		noAuthPrefixErr = fmt.Errorf("invalid authorization header format")
		noEndpointErr   = fmt.Errorf("failed to find endpoint")
		tokenInvalidErr = fmt.Errorf("access token is invalid")
		accessDeniedErr = fmt.Errorf("access denied")

		req = endpointCreate
	)

	tests := []struct {
		name                 string
		args                 args
		err                  error
		accessRepositoryMock accessRepositoryMockFunc
		tokenOperationsMock  tokenOperationsMockFunc
	}{
		{
			name: "metadata not provided error case",
			args: args{
				ctx: ctxNoMd,
				req: req,
			},
			err: noMdErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "authorization header not provided error case",
			args: args{
				ctx: ctxNoAuthHeader,
				req: req,
			},
			err: noAuthHeaderErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "authorization header format error case",
			args: args{
				ctx: ctxNoAuthPrefix,
				req: req,
			},
			err: noAuthPrefixErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				return mock
			},
		},
		{
			name: "endpoint not found error case",
			args: args{
				ctx: ctx,
				req: endpointNotExists,
			},
			err: noEndpointErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(minimock.AnyContext).Return(endpointPermissions, nil)
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
			err: tokenInvalidErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(accessToken, secretKeyBytes).Return(nil, tokenInvalidErr)
				return mock
			},
		},
		{
			name: "access denied error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: accessDeniedErr,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(accessToken, secretKeyBytes).Return(claimsUser, nil)
				return mock
			},
		},
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(accessToken, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv := accessService.NewService(accessRepositoryMock, tokenOperationsMock, jwtConfig)

			err := srv.Check(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
