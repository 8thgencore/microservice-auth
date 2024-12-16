package converter

import (
	"database/sql"

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

// ToUserUpdateDAO converts service model to DAO model
func ToUserUpdateDAO(user *model.UserUpdate) *dao.UserUpdate {
	update := &dao.UserUpdate{
		ID: user.ID,
	}

	if user.Name != nil {
		update.Name = sql.NullString{String: *user.Name, Valid: true}
	}
	if user.Email != nil {
		update.Email = sql.NullString{String: *user.Email, Valid: true}
	}
	if user.Role != nil {
		update.Role = sql.NullString{String: *user.Role, Valid: true}
	}
	if user.Version != nil {
		update.Version = sql.NullInt32{Int32: *user.Version, Valid: true}
	}

	return update
}
