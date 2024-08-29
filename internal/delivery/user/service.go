package user

import (
	"github.com/8thgencore/microservice_auth/internal/service"
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// UserImplementation structure describes API layer.
type UserImplementation struct {
	pb.UnimplementedUserV1Server
	userService service.UserService
}

// NewUserImplementation creates new object of API layer.
func NewUserImplementation(userService service.UserService) *UserImplementation {
	return &UserImplementation{
		userService: userService,
	}
}
