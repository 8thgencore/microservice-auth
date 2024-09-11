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
	ErrAccessTokenGen  = errors.New("failed to generate access token")
	ErrRefreshTokenGen = errors.New("failed to generate refresh token")
	ErrInvalidRefresh  = errors.New("invalid refresh token")
)

func (s *serv) Login(ctx context.Context, creds *model.UserCreds) (*model.TokenPair, error) {
	authInfo, err := s.userRepository.GetAuthInfo(ctx, creds.Username)
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(creds.Password))
	if err != nil {
		return nil, ErrWrongPassword
	}

	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		Name: authInfo.Username,
		Role: authInfo.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.AccessTokenTTL,
	)
	if err != nil {
		return nil, ErrAccessTokenGen
	}

	refreshToken, err := s.tokenOperations.GenerateRefreshToken(
		authInfo.ID,
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return nil, ErrRefreshTokenGen
	}

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := s.tokenOperations.VerifyRefreshToken(refreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", ErrInvalidRefresh
	}

	user, err := s.userRepository.Get(ctx, claims.UserID)
	if err != nil {
		return "", ErrUserNotFound
	}

	accessToken, err := s.tokenOperations.GenerateAccessToken(model.User{
		Name: user.Name,
		Role: user.Role,
	},
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.AccessTokenTTL,
	)
	if err != nil {
		return "", ErrAccessTokenGen
	}

	return accessToken, nil
}

func (s *serv) GetRefreshToken(_ context.Context, oldRefreshToken string) (string, error) {
	claims, err := s.tokenOperations.VerifyRefreshToken(oldRefreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return "", ErrInvalidRefresh
	}

	refreshToken, err := s.tokenOperations.GenerateRefreshToken(
		claims.UserID,
		[]byte(s.jwtConfig.SecretKey),
		s.jwtConfig.RefreshTokenTTL,
	)
	if err != nil {
		return "", ErrRefreshTokenGen
	}

	return refreshToken, nil
}
