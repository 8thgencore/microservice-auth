package converter

import (
	"github.com/8thgencore/microservice-auth/internal/model"
	accessv1 "github.com/8thgencore/microservice-auth/pkg/access/v1"
)

// ToEndpointPermissionsMap converts slice of service layer structures to map.
func ToEndpointPermissionsMap(endpointPermissions []*model.EndpointPermissions) map[string][]string {
	res := make(map[string][]string)
	for _, e := range endpointPermissions {
		res[e.Endpoint] = e.Roles
	}
	return res
}

// ToEndpointPermissionsFromAPI converts structure of API layer to service layer model.
func ToEndpointPermissionsFromAPI(endpointPermissions *accessv1.EndpointPermissions) *model.EndpointPermissions {
	return &model.EndpointPermissions{
		Endpoint: endpointPermissions.Endpoint,
		Roles:    endpointPermissions.AllowedRoles,
	}
}

// ToEndpointPermissionsService converts service layer model to structure of API layer.
func ToEndpointPermissionsService(endpointPermissions *model.EndpointPermissions) *accessv1.EndpointPermissions {
	return &accessv1.EndpointPermissions{
		Endpoint:     endpointPermissions.Endpoint,
		AllowedRoles: endpointPermissions.Roles,
	}
}
