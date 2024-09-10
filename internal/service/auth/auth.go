package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/8thgencore/microservice-auth/internal/model"
)

func (s *serv) Login(ctx context.Context, creds *model.UserCreds) (*model.TokenPair, error) {
	// Get role and hashed password by username from storage
	authInfo, err := s.userRepository.GetAuthInfo(ctx, creds.Username)
	if err != nil {
		return &model.TokenPair{}, errors.New("user not found")
	}

	// Check password correctness
	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(creds.Password))
	if err != nil {
		return &model.TokenPair{}, errors.New("wrong password")
	}

	// Generate access token with detailed user info (name, role)
	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		Name: authInfo.Username,
		Role: authInfo.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.AccessTokenTTL,
	)
	if err != nil {
		return &model.TokenPair{}, errors.New("failed to generate access token")
	}

	// Generate refresh token with minimal information (e.g., user ID or username)
	refreshToken, err := s.tokenOperations.GenerateRefreshToken(
		authInfo.ID,
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return &model.TokenPair{}, errors.New("failed to generate refresh token")
	}

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	// Verify the refresh token to get the user identity (no sensitive data)
	claims, err := s.tokenOperations.VerifyRefreshToken(refreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.userRepository.Get(ctx, claims.UserID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate new access token with user details (name and role)
	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		Name: user.Name,
		Role: user.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.AccessTokenTTL,
	)
	if err != nil {
		return "", errors.New("failed to generate access token")
	}

	return accessToken, nil
}

func (s *serv) GetRefreshToken(_ context.Context, oldRefreshToken string) (string, error) {
	// Verify the old refresh token to get the user identity
	claims, err := s.tokenOperations.VerifyRefreshToken(oldRefreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// Generate a new refresh token with minimal info (e.g., user ID or username)
	refreshToken, err := s.tokenOperations.GenerateRefreshToken(
		claims.UserID,
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return "", errors.New("failed to generate refresh token")
	}

	return refreshToken, nil
}
