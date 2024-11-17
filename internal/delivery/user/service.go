package user

import (
	"github.com/8thgencore/microservice-auth/internal/service"
	userv1 "github.com/8thgencore/microservice-auth/pkg/pb/user/v1"
)

// Implementation structure describes API layer.
type Implementation struct {
	userv1.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation creates new object of API layer.
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
