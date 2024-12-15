package converter

import (
	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-auth/internal/repository/user/dao"
)

// ToUserFromRepo converts repository layer model to structure of service layer.
func ToUserFromRepo(user *dao.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		Version:   user.Version,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ToAuthInfoFromRepo converts repository layer model to structure of service layer.
func ToAuthInfoFromRepo(authInfo *dao.AuthInfo) *model.AuthInfo {
	return &model.AuthInfo{
		ID:       authInfo.ID,
		Username: authInfo.Username,
		Role:     authInfo.Role,
		Version:  authInfo.Version,
		Password: authInfo.Password,
	}
}
