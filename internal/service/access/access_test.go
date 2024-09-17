package access

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	"github.com/8thgencore/microservice-auth/internal/config"
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository"
	repositoryMocks "github.com/8thgencore/microservice-auth/internal/repository/mocks"
	"github.com/8thgencore/microservice-auth/internal/tokens"
	tokenMocks "github.com/8thgencore/microservice-auth/internal/tokens/mocks"
)

var (
	mdNoAuthHeader = metadata.New(map[string]string{"header": "access_token"})
	mdNoAuthPrefix = metadata.New(map[string]string{"Authorization": "access_token"})
	md             = metadata.New(map[string]string{"Authorization": "Bearer access_token"})

	ctxNoMd         = context.Background()
	ctx             = metadata.NewIncomingContext(ctxNoMd, md)
	ctxNoAuthHeader = metadata.NewIncomingContext(ctxNoMd, mdNoAuthHeader)
	ctxNoAuthPrefix = metadata.NewIncomingContext(ctxNoMd, mdNoAuthPrefix)

	username  = "username"
	roleUser  = "USER"
	roleAdmin = "ADMIN"

	jwtConfig = config.JWTConfig{
		SecretKey:       "secret",
		AccessTokenTTL:  time.Duration(30 * time.Minute),
		RefreshTokenTTL: time.Duration(360 * time.Minute),
	}

	token = "access_token"

	secretKeyBytes = []byte(jwtConfig.SecretKey)

	claimsAdmin = &model.UserClaims{
		Username: username,
		Role:     roleAdmin,
	}

	claimsUser = &model.UserClaims{
		Username: username,
		Role:     roleUser,
	}
)

func TestNewService(t *testing.T) {
	t.Parallel()

	var (
		mc = minimock.NewController(t)

		endpoint = getRoleEndpointsEndpoint

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpoint, Roles: []string{roleAdmin}},
		}
	)

	tests := []struct {
		name                 string
		expectedErr          error
		expectedRolesMap     map[string][]string
		accessRepositoryMock func(mc *minimock.Controller) repository.AccessRepository
		tokenOperationsMock  func(mc *minimock.Controller) tokens.TokenOperations
	}{
		{
			name:        "success case",
			expectedErr: nil,
			expectedRolesMap: map[string][]string{
				endpoint: {roleAdmin},
			},
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				return tokenMocks.NewTokenOperationsMock(mc)
			},
		},
		{
			name:             "error on access repository",
			expectedErr:      ErrFailedToReadAccessPolicy,
			expectedRolesMap: nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(nil, errors.New("some error"))
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				return tokenMocks.NewTokenOperationsMock(mc)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)

			srv, err := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)
			if tt.expectedErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expectedErr, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, srv)

				// Check that the role map is created correctly
				accessSrv, ok := srv.(*accessService)
				require.True(t, ok)
				require.Equal(t, tt.expectedRolesMap, accessSrv.accessibleRoles)
			}
		})
	}
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
		mc = minimock.NewController(t)

		endpointCreate      = "/chat_v1.ChatV1/Create"
		endpointDelete      = "/chat_v1.ChatV1/Delete"
		endpointSendMessage = "/chat_v1.ChatV1/SendMessage"
		endpointNotExists   = "/chat_v1.ChatV1/NotExists"

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpointCreate, Roles: []string{roleAdmin}},
			{Endpoint: endpointDelete, Roles: []string{roleAdmin}},
			{Endpoint: endpointSendMessage, Roles: []string{roleAdmin, roleUser}},
		}

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
			err: ErrMetadataNotProvided,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
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
			err: ErrAuthHeaderNotProvided,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
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
			err: ErrInvalidAuthHeaderFormat,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
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
			err: ErrEndpointNotFound,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
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
			err: ErrInvalidAccessToken,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.
					Expect(token, secretKeyBytes).
					Return(nil, ErrInvalidAccessToken)
				return mock
			},
		},
		{
			name: "access denied error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: ErrAccessDenied,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsUser, nil)
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
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)

			srv, err := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)
			require.NoError(t, err)

			err = srv.Check(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestGetRoleEndpoints(t *testing.T) {
	t.Parallel()

	type myKey string
	const myKeyValue myKey = "myKey"
	ctxSecond := context.WithValue(ctx, myKeyValue, "value")

	var (
		mc = minimock.NewController(t)

		endpoint = getRoleEndpointsEndpoint

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpoint, Roles: []string{roleAdmin}},
		}
	)

	tests := []struct {
		name                 string
		expectedErr          error
		expectedResult       []*model.EndpointPermissions
		accessRepositoryMock func(mc *minimock.Controller) repository.AccessRepository
		tokenOperationsMock  func(mc *minimock.Controller) tokens.TokenOperations
	}{
		{
			name:           "check endpoint error case",
			expectedErr:    ErrAccessDenied,
			expectedResult: nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsUser, nil)
				return mock
			},
		},
		{
			name:           "get role endpoint error case",
			expectedErr:    ErrEndpointNotFound,
			expectedResult: nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.When(ctx).Then(endpointPermissions, nil)
				mock.GetRoleEndpointsMock.When(ctxSecond).Then(nil, ErrEndpointNotFound)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
		{
			name:           "get role endpoints success case",
			expectedErr:    nil,
			expectedResult: endpointPermissions,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.When(ctx).Then(endpointPermissions, nil)
				mock.GetRoleEndpointsMock.When(ctxSecond).Then(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)

			srv, err := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)
			require.NoError(t, err)
			require.NotNil(t, srv)

			result, err := srv.GetRoleEndpoints(ctxSecond)
			require.Equal(t, tt.expectedErr, err)
			require.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestAddRoleEndpoint(t *testing.T) {
	t.Parallel()

	type accessRepositoryMockFunc func(mc *minimock.Controller) repository.AccessRepository
	type tokenOperationsMockFunc func(mc *minimock.Controller) tokens.TokenOperations

	var (
		mc    = minimock.NewController(t)
		roles = []string{roleAdmin}

		endpoint = addRoleEndpointEndpoint

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpoint, Roles: []string{roleAdmin}},
		}
	)

	tests := []struct {
		name                 string
		err                  error
		accessRepositoryMock accessRepositoryMockFunc
		tokenOperationsMock  tokenOperationsMockFunc
	}{
		{
			name: "check endpoint error case",
			err:  ErrAccessDenied,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsUser, nil)
				return mock
			},
		},
		{
			name: "add role endpoint error case",
			err:  ErrEndpointAlreadyExists,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.AddRoleEndpointMock.Expect(ctx, endpoint, roles).Return(ErrEndpointAlreadyExists)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
		{
			name: "add role endpoint success case",
			err:  nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.AddRoleEndpointMock.Expect(ctx, endpoint, roles).Return(nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv, _ := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)

			err := srv.AddRoleEndpoint(ctx, endpoint, roles)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestUpdateRoleEndpoint(t *testing.T) {
	t.Parallel()

	var (
		mc    = minimock.NewController(t)
		roles = []string{roleAdmin}

		endpoint = updateRoleEndpointEndpoint

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpoint, Roles: []string{roleAdmin}},
		}
	)

	tests := []struct {
		name                 string
		err                  error
		accessRepositoryMock func(mc *minimock.Controller) repository.AccessRepository
		tokenOperationsMock  func(mc *minimock.Controller) tokens.TokenOperations
	}{
		{
			name: "check endpoint error case",
			err:  ErrAccessDenied,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)

				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsUser, nil)
				return mock
			},
		},
		{
			name: "update role endpoint error case",
			err:  ErrEndpointNotFound,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.UpdateRoleEndpointMock.Expect(ctx, endpoint, roles).Return(ErrEndpointNotFound)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
		{
			name: "update role endpoint success case",
			err:  nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.UpdateRoleEndpointMock.Expect(ctx, endpoint, roles).Return(nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv, _ := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)

			err := srv.UpdateRoleEndpoint(ctx, endpoint, roles)
			require.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteRoleEndpoint(t *testing.T) {
	t.Parallel()

	var (
		mc = minimock.NewController(t)

		endpoint = deleteRoleEndpointEndpoint

		endpointPermissions = []*model.EndpointPermissions{
			{Endpoint: endpoint, Roles: []string{roleAdmin}},
		}
	)

	tests := []struct {
		name                 string
		err                  error
		accessRepositoryMock func(mc *minimock.Controller) repository.AccessRepository
		tokenOperationsMock  func(mc *minimock.Controller) tokens.TokenOperations
	}{
		{
			name: "check endpoint error case",
			err:  ErrAccessDenied,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsUser, nil)
				return mock
			},
		},
		{
			name: "delete role endpoint error case",
			err:  ErrEndpointNotFound,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.DeleteRoleEndpointMock.Expect(ctx, endpoint).Return(ErrEndpointNotFound)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
		{
			name: "delete role endpoint success case",
			err:  nil,
			accessRepositoryMock: func(mc *minimock.Controller) repository.AccessRepository {
				mock := repositoryMocks.NewAccessRepositoryMock(mc)
				mock.GetRoleEndpointsMock.Expect(ctx).Return(endpointPermissions, nil)
				mock.DeleteRoleEndpointMock.Expect(ctx, endpoint).Return(nil)
				return mock
			},
			tokenOperationsMock: func(mc *minimock.Controller) tokens.TokenOperations {
				mock := tokenMocks.NewTokenOperationsMock(mc)
				mock.VerifyAccessTokenMock.Expect(token, secretKeyBytes).Return(claimsAdmin, nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accessRepositoryMock := tt.accessRepositoryMock(mc)
			tokenOperationsMock := tt.tokenOperationsMock(mc)
			srv, _ := NewService(ctx, accessRepositoryMock, tokenOperationsMock, jwtConfig)

			err := srv.DeleteRoleEndpoint(ctx, endpoint)
			require.Equal(t, tt.err, err)
		})
	}
}
