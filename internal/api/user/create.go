package user

import (
	"context"
	"github.com/8thgencore/microservice_auth/internal/api/user/converter"

	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Create is used for creating new user.
func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := i.userService.Create(ctx, converter.ToUserCreateFromDesc(req.GetUser()))
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}
