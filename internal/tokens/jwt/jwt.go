package jwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/tokens"
)

type tokenOperations struct{}

var _ tokens.TokenOperations = (*tokenOperations)(nil)

// NewTokenOperations creates a new object for using token functions.
func NewTokenOperations() tokens.TokenOperations {
	return &tokenOperations{}
}

// GenerateAccessToken creates JWT access token for the user.
func (t *tokenOperations) GenerateAccessToken(
	user model.User, secretKey []byte,
	duration time.Duration,
) (string, error) {
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		Username: user.Name,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// GenerateRefreshToken creates JWT refresh token with minimal claims.
func (t *tokenOperations) GenerateRefreshToken(
	userID int64,
	secretKey []byte,
	duration time.Duration,
) (string, error) {
	claims := model.RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// VerifyAccessToken checks the validity of an access token.
func (t *tokenOperations) VerifyAccessToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

// VerifyRefreshToken checks the validity of a refresh token.
func (t *tokenOperations) VerifyRefreshToken(tokenStr string, secretKey []byte) (*model.RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}
			return secretKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %s", err.Error())
	}

	claims, ok := token.Claims.(*model.RefreshClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token claims")
	}

	return claims, nil
}
