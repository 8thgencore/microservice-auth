package converter

import (
	"github.com/8thgencore/microservice-auth/internal/model"
	authv1 "github.com/8thgencore/microservice-auth/pkg/auth/v1"
)

// ToUserLoginFromAPI converts structure of API layer to service layer model.
func ToUserLoginFromAPI(creds *authv1.Creds) *model.UserCreds {
	return &model.UserCreds{
		Username: creds.Username,
		Password: creds.Password,
	}
}
