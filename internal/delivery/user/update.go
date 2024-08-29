package user

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/delivery/user/converter"
	"github.com/golang/protobuf/ptypes/empty"

	user1 "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Update is used for updating user info.
func (impl *UserImplementation) Update(ctx context.Context, req *user1.UpdateRequest) (*empty.Empty, error) {
	err := impl.userService.Update(ctx, converter.ToUserUpdateFromApi(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
