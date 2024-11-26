package utils

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authMetadataHeader = "authorization"
	authPrefix         = "Bearer "
)

var (
	// ErrMetadataNotProvided occurs when metadata is not passed in the request.
	ErrMetadataNotProvided = errors.New("metadata is not provided")
	// ErrAuthHeaderNotProvided occurs when the authorization header is missing from the request.
	ErrAuthHeaderNotProvided = errors.New("authorization header is not provided")
	// ErrInvalidAuthHeaderFormat occurs when the authorization header has an incorrect format.
	ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")
)

func ExtractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", ErrMetadataNotProvided
	}

	authHeader, ok := md[authMetadataHeader]
	if !ok || len(authHeader) == 0 {
		return "", ErrAuthHeaderNotProvided
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return "", ErrInvalidAuthHeaderFormat
	}

	return strings.TrimPrefix(authHeader[0], authPrefix), nil
}
