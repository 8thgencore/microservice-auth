package user

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/delivery/user/converter"
	user1 "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

func (impl *UserImplementation) Get(ctx context.Context, req *user1.GetRequest) (*user1.GetResponse, error) {
	user, err := impl.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &user1.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
