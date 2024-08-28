package user

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/api/user/converter"
	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Update is used for updating user info.
func (i *Implementation) Update(ctx context.Context, req *pb.UpdateRequest) (*empty.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUserUpdateFromApi(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
