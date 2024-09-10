package auth

import (
	"context"

	"github.com/8thgencore/microservice-auth/internal/converter"
	authv1 "github.com/8thgencore/microservice-auth/pkg/auth/v1"
)

// Login user and return refresh token.
func (i *Implementation) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	refreshToken, err := i.authService.Login(ctx, converter.ToUserLoginFromAPI(req.GetCreds()))
	if err != nil {
		return nil, err
	}

	return &authv1.LoginResponse{
		RefreshToken: refreshToken,
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
