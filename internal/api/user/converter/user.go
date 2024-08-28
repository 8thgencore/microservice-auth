package converter

import (
	"database/sql"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/8thgencore/microservice_auth/internal/model"
	pb "github.com/8thgencore/microservice_auth/pkg/user/v1"
)

// ToUserFromService converts service layer model to structure of API layer.
func ToUserFromService(user *model.User) *pb.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &pb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      pb.Role(pb.Role_value[user.Role]),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

// ToUserCreateFromDesc converts structure of API layer to service layer model.
func ToUserCreateFromDesc(user *pb.UserCreate) *model.UserCreate {
	return &model.UserCreate{
		Name:            user.Name,
		Email:           user.Email,
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
		Role:            pb.Role_name[int32(user.Role)],
	}
}

// ToUserUpdateFromDesc converts structure of API layer to service layer model.
func ToUserUpdateFromDesc(user *pb.UserUpdate) *model.UserUpdate {
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
			String: pb.Role_name[int32(user.Role)],
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
