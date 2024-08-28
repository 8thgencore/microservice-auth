package user

import (
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
	"github.com/8thgencore/microservice_auth/internal/service"
)

// Implementation structure describes API layer.
type Implementation struct {
	pb.UnimplementedUserV1Server
	userService service.UserService
}

// NewImplementation creates new object of API layer.
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
