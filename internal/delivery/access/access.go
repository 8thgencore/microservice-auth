package access

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/8thgencore/microservice-auth/internal/converter"
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

// AddRoleEndpoint adds a new role-endpoint permission.
func (i *Implementation) AddRoleEndpoint(
	ctx context.Context,
	req *accessv1.AddRoleEndpointRequest,
) (*empty.Empty, error) {
	err := i.accessService.AddRoleEndpoint(ctx, req.GetEndpoint(), converter.ToRoleStrings(req.GetAllowedRoles()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// UpdateRoleEndpoint updates an existing role-endpoint permission.
func (i *Implementation) UpdateRoleEndpoint(
	ctx context.Context,
	req *accessv1.UpdateRoleEndpointRequest,
) (*empty.Empty, error) {
	err := i.accessService.UpdateRoleEndpoint(ctx, req.GetEndpoint(), converter.ToRoleStrings(req.GetAllowedRoles()))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// DeleteRoleEndpoint deletes an existing role-endpoint permission.
func (i *Implementation) DeleteRoleEndpoint(
	ctx context.Context,
	req *accessv1.DeleteRoleEndpointRequest,
) (*empty.Empty, error) {
	err := i.accessService.DeleteRoleEndpoint(ctx, req.GetEndpoint())
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// ListRoleEndpoints retrieves the list of role-endpoint permissions.
func (i *Implementation) ListRoleEndpoints(
	ctx context.Context,
	_ *empty.Empty,
) (*accessv1.ListRoleEndpointsResponse, error) {
	endpoints, err := i.accessService.ListRoleEndpoints(ctx)
	if err != nil {
		return nil, err
	}

	// Convert the service's response to the gRPC response format
	var endpointPermissions []*accessv1.EndpointPermissions
	for _, ep := range endpoints {
		endpointPermissions = append(endpointPermissions, converter.ToEndpointPermissionsService(ep))
	}

	return &accessv1.ListRoleEndpointsResponse{
		EndpointPermissions: endpointPermissions,
	}, nil
}
