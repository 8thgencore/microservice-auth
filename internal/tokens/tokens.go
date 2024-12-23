package tokens

import "github.com/8thgencore/microservice-auth/internal/model"

// TokenOperations is the interface for token functions.
type TokenOperations interface {
	// GenerateAccessToken creates JWT access token for the user.
	GenerateAccessToken(user model.User) (string, error)
	// GenerateRefreshToken creates JWT refresh token with minimal claims (e.g., only username).
	GenerateRefreshToken(userID string) (string, error)
	// VerifyAccessToken checks the validity of an access token.
	VerifyAccessToken(tokenStr string) (*model.UserClaims, error)
	// VerifyRefreshToken checks the validity of a refresh token.
	VerifyRefreshToken(tokenStr string) (*model.RefreshClaims, error)
}
