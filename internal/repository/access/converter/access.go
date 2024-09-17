package converter

import (
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository/access/dao"
)

// ToEndpointPermissionsFromRepo converts repository layer model to structure of service layer.
func ToEndpointPermissionsFromRepo(endpointPermissions []*dao.EndpointPermissions) []*model.EndpointPermissions {
	var res []*model.EndpointPermissions
	for _, e := range endpointPermissions {
		res = append(res, &model.EndpointPermissions{
			Endpoint: e.Endpoint,
			Roles:    e.Roles,
		})
	}

	return res
}
