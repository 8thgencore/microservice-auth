package model

import jwt "github.com/golang-jwt/jwt/v5"

// UserClaims is custom wrapper for jwt claims.
type UserClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

// RefreshClaims - a data structure containing the minimum data for the refresh token.
type RefreshClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}
