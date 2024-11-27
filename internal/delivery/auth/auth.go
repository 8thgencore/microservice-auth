package auth

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/converter"
	authv1 "github.com/8thgencore/microservice-auth/pkg/pb/auth/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Login user and return refresh token.
func (i *Implementation) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	tokenPair, err := i.authService.Login(ctx, converter.ToUserLoginFromAPI(req.GetCreds()))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &authv1.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// RefreshTokens returns access token for later operations.
func (i *Implementation) RefreshTokens(
	ctx context.Context,
	req *authv1.RefreshTokensRequest,
) (*authv1.RefreshTokensResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &authv1.RefreshTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Logout invalidates the refresh token.
func (i *Implementation) Logout(
	ctx context.Context,
	req *authv1.LogoutRequest,
) (*empty.Empty, error) {
	err := i.authService.Logout(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	return &empty.Empty{}, nil
}
