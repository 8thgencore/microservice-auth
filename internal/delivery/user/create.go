package user

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/delivery/user/converter"

	userv1 "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Create is used for creating new user.
func (impl *UserImplementation) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	id, err := impl.userService.Create(ctx, converter.ToUserCreateFromApi(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &userv1.CreateResponse{
		Id: id,
	}, nil
}
