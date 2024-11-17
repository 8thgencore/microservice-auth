package auth

import (
	"github.com/8thgencore/microservice-auth/internal/service"
	authv1 "github.com/8thgencore/microservice-auth/pkg/pb/auth/v1"
)

// Implementation structure describes API layer.
type Implementation struct {
	authv1.UnimplementedAuthV1Server
	authService service.AuthService
}

// NewImplementation creates new object of API layer.
func NewImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
