package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/8thgencore/microservice-auth/internal/model"
	userv1 "github.com/8thgencore/microservice-auth/pkg/user/v1"
)

// ToUserFromService converts service layer model to structure of API layer.
func ToUserFromService(user *model.User) *userv1.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &userv1.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      userv1.Role(userv1.Role_value[user.Role]),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserCreateFromAPI converts structure of API layer to service layer model.
func ToUserCreateFromAPI(user *userv1.UserCreate) *model.UserCreate {
	return &model.UserCreate{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            userv1.Role_name[int32(user.Role)],
	}
}

// ToUserUpdateFromAPI converts structure of API layer to service layer model.
func ToUserUpdateFromAPI(user *userv1.UserUpdate) *model.UserUpdate {
	var (
		name  sql.NullString
		email sql.NullString
		role  sql.NullString
	)

	if user.Name != nil {
		name = sql.NullString{
			String: user.Name.GetValue(),
			Valid:  true,
		}
	}
	if user.Email != nil {
		email = sql.NullString{
			String: user.Email.GetValue(),
			Valid:  true,
		}
	}

	if user.Role != 0 {
		role = sql.NullString{
			String: userv1.Role_name[int32(user.Role)],
			Valid:  true,
		}
	}

	return &model.UserUpdate{
		ID:    user.Id,
		Name:  name,
		Email: email,
		Role:  role,
	}
}

// ToRoleStrings converts a list of Role enum values to a list of role strings.
func ToRoleStrings(roles []userv1.Role) []string {
	var roleStrings []string
	for _, role := range roles {
		roleStrings = append(roleStrings, userv1.Role_name[int32(role)])
	}

	return roleStrings
}

// ToRoleEnumsAPI converts a list of role strings to a list of Role enum values.
func ToRoleEnumsAPI(roleStrings []string) []userv1.Role {
	var roles []userv1.Role
	for _, roleStr := range roleStrings {
		if val, ok := userv1.Role_value[roleStr]; ok {
			roles = append(roles, userv1.Role(val))
		}
	}
	return roles
}
