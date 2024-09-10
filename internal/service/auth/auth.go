package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/8thgencore/microservice-auth/internal/model"
)

func (s *serv) Login(ctx context.Context, creds *model.UserCreds) (string, error) {
	// Get role and hashed password by username from storage
	authInfo, err := s.userRepository.GetAuthInfo(ctx, creds.Username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(creds.Password))
	if err != nil {
		return "", errors.New("wrong password")
	}

	refreshToken, err := s.tokenOperations.Generate(model.User{
		Name: authInfo.Username,
		Role: authInfo.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}

func (s *serv) GetAccessToken(_ context.Context, refreshToken string) (string, error) {
	claims, err := s.tokenOperations.Verify(refreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	accessToken, err := s.tokenOperations.Generate(model.User{
		Name: claims.Username,
		Role: claims.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.AccessTokenTTL,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return accessToken, nil
}

func (s *serv) GetRefreshToken(_ context.Context, oldRefreshToken string) (string, error) {
	claims, err := s.tokenOperations.Verify(oldRefreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	refreshToken, err := s.tokenOperations.Generate(model.User{
		Name: claims.Username,
		Role: claims.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return refreshToken, nil
}
