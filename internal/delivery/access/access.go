package access

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	accessv1 "github.com/8thgencore/microservice-auth/pkg/access/v1"
)

// Check performs user authorization.
func (i *Implementation) Check(ctx context.Context, req *accessv1.CheckRequest) (*empty.Empty, error) {
	err := i.accessService.Check(ctx, req.GetEndpoint())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
