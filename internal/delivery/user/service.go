package user

import (
	"github.com/8thgencore/microservice-auth/internal/service"
	userv1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
)

// UserImplementation structure describes API layer.
type UserImplementation struct {
	userv1.UnimplementedUserV1Server
	userService service.UserService
}

// NewUserImplementation creates new object of API layer.
func NewUserImplementation(userService service.UserService) *UserImplementation {
	return &UserImplementation{
		userService: userService,
	}
}
