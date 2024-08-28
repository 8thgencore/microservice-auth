package user

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Delete is used for deleting user.
func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}