package auth

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/converter"
	authv1 "github.com/8thgencore/microservice-auth/pkg/auth/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

// Login user and return refresh token.
func (i *Implementation) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	tokenPair, err := i.authService.Login(ctx, converter.ToUserLoginFromAPI(req.GetCreds()))
	if err != nil {
		return nil, err
	}

	return &authv1.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

// GetAccessToken returns access token for later operations.
func (i *Implementation) GetAccessToken(
	ctx context.Context,
	req *authv1.GetAccessTokenRequest,
) (*authv1.GetAccessTokenResponse, error) {
	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authv1.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}

// GetRefreshToken updates refresh token.
func (i *Implementation) GetRefreshToken(
	ctx context.Context,
	req *authv1.GetRefreshTokenRequest,
) (*authv1.GetRefreshTokenResponse, error) {
	refreshToken, err := i.authService.GetRefreshToken(ctx, req.GetOldRefreshToken())
	if err != nil {
		return nil, err
	}

	return &authv1.GetRefreshTokenResponse{
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
		return nil, err
	}

	return &empty.Empty{}, nil
}
