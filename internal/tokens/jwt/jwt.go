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
	user model.User,
	secretKey []byte,
	duration time.Duration,
) (string, error) {
	claims := model.UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(user.ID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		Username: user.Name,
		Role:     user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("could not sign access token: %w", err)
	}

	return signedToken, nil
}

// GenerateRefreshToken creates JWT refresh token with minimal claims.
func (t *tokenOperations) GenerateRefreshToken(
	userID string,
	secretKey []byte,
	duration time.Duration,
) (string, error) {
	claims := model.RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   fmt.Sprint(userID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("could not sign refresh token: %w", err)
	}

	return signedToken, nil
}

// VerifyAccessToken checks the validity of an access token.
func (t *tokenOperations) VerifyAccessToken(tokenStr string, secretKey []byte) (*model.UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid access token: token is not valid")
	}

	claims, ok := token.Claims.(*model.UserClaims)
	if !ok {
		return nil, errors.New("invalid access token claims")
	}

	return claims, nil
}

// VerifyRefreshToken checks the validity of a refresh token.
func (t *tokenOperations) VerifyRefreshToken(tokenStr string, secretKey []byte) (*model.RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&model.RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token: token is not valid")
	}

	claims, ok := token.Claims.(*model.RefreshClaims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	return claims, nil
}
