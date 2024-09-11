package access

import (
	"github.com/8thgencore/microservice-auth/internal/service"
	desc "github.com/8thgencore/microservice-auth/pkg/access/v1"
)

// Implementation structure describes API layer.
type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

// NewImplementation creates new object of API layer.
func NewImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
