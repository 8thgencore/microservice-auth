package user

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	user1 "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// Delete is used for deleting user.
func (imlp *UserImplementation) Delete(ctx context.Context, req *user1.DeleteRequest) (*empty.Empty, error) {
	err := imlp.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
