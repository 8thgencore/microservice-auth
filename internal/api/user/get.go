package user

import (
	"context"

	"github.com/8thgencore/microservice_auth/internal/api/user/converter"
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Get is used to obtain user info.
func (i *Implementation) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	user, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
