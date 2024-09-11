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
	ErrLogoutFailed    = errors.New("failed to logout")
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
		ID:   authInfo.ID,
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
	revoked, err := s.tokenRepository.IsTokenRevoked(ctx, refreshToken)
	if err != nil {
		return "", ErrRefreshTokenGen
	}
	if revoked {
		return "", ErrInvalidRefresh
	}

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

func (s *serv) GetRefreshToken(ctx context.Context, oldRefreshToken string) (string, error) {
	revoked, err := s.tokenRepository.IsTokenRevoked(ctx, oldRefreshToken)
	if err != nil {
		return "", ErrInvalidRefresh
	}
	if revoked {
		return "", ErrInvalidRefresh
	}

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

	if err = s.tokenRepository.AddRevokedToken(ctx, oldRefreshToken); err != nil {
		return "", ErrRefreshTokenGen
	}

	return refreshToken, nil
}

// Logout invalidates the refresh token.
func (s *serv) Logout(ctx context.Context, refreshToken string) error {
	// Verify the refresh token
	_, err := s.tokenOperations.VerifyRefreshToken(refreshToken, []byte(s.jwtConfig.SecretKey))
	if err != nil {
		return ErrInvalidRefresh
	}

	if err = s.tokenRepository.AddRevokedToken(ctx, refreshToken); err != nil {
		return ErrLogoutFailed
	}

	return nil
}
