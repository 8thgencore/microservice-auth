package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/8thgencore/microservice-auth/internal/model"
)

// Errors
var (
	ErrUserNotFound    = errors.New("user not found")
	ErrWrongPassword   = errors.New("wrong password")
	ErrTokenGeneration = errors.New("failed to generate token")
	ErrInvalidRefresh  = errors.New("invalid refresh token")
	ErrLogoutFailed    = errors.New("failed to logout")
)

// Login checks the user's credentials and returns a token pair if they are valid
func (s *authService) Login(ctx context.Context, creds *model.UserCreds) (*model.TokenPair, error) {
	authInfo, err := s.userRepository.GetAuthInfo(ctx, creds.Username)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(creds.Password))
	if err != nil {
		return nil, ErrWrongPassword
	}

	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		ID:   authInfo.ID,
		Name: authInfo.Username,
		Role: authInfo.Role,
	},
	)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	refreshToken, err := s.tokenOperations.GenerateRefreshToken(authInfo.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// GetAccessToken generates a new access token for a user given a valid refresh token
func (s *authService) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	if err := s.validateRefreshToken(refreshToken); err != nil {
		return "", err
	}

	claims, err := s.tokenOperations.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", ErrInvalidRefresh
	}

	user, err := s.userRepository.Get(ctx, claims.UserID)
	if err != nil {
		return "", ErrUserNotFound
	}

	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		ID:   user.ID,
		Name: user.Name,
		Role: user.Role,
	},
	)
	if err != nil {
		return "", ErrTokenGeneration
	}

	return accessToken, nil
}

// GetRefreshToken generates a new refresh token for a user given a valid old refresh token
func (s *authService) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	if err := s.validateRefreshToken(oldRefreshToken); err != nil {
		return "", err
	}

	claims, err := s.tokenOperations.VerifyRefreshToken(oldRefreshToken)
	if err != nil {
		return "", ErrInvalidRefresh
	}

	refreshToken, err := s.tokenOperations.GenerateRefreshToken(claims.UserID)
	if err != nil {
		return "", ErrTokenGeneration
	}

	if err = s.tokenRepository.AddRevokedToken(ctx, oldRefreshToken); err != nil {
		return "", ErrTokenGeneration
	}

	return refreshToken, nil
}

// Logout invalidates the refresh token by adding it to the list of revoked tokens
func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	_, err := s.tokenOperations.VerifyRefreshToken(refreshToken)
	if err != nil {
		return ErrInvalidRefresh
	}

	if err = s.tokenRepository.AddRevokedToken(ctx, refreshToken); err != nil {
		return ErrLogoutFailed
	}

	return nil
}

// validateRefreshToken checks if a refresh token is valid and not revoked
func (s *authService) validateRefreshToken(refreshToken string) error {
	revoked, err := s.tokenRepository.IsTokenRevoked(context.Background(), refreshToken)
	if err != nil {
		return ErrInvalidRefresh
	}
	if revoked {
		return ErrInvalidRefresh
	}

	return nil
}
