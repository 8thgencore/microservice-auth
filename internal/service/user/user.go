package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/8thgencore/microservice-auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// ErrUserExists - custom error for user name duplicate.
var ErrUserExists = errors.New("user with provided name or email already exists")

func (s *serv) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	if user.Password != user.PasswordConfirm {
		return 0, errors.New("passwords don't match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to process password")
	}
	user.Password = string(hashedPassword)

	var id int64
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRepository.Create(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, &model.Log{
			Text: fmt.Sprintf("Created user with id: %d", id),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, ErrUserExists) {
			return 0, ErrUserExists
		}
		return 0, errors.New("failed to create user")
	}

	return id, nil
}

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var user *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		user, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, &model.Log{
			Text: fmt.Sprintf("Read info about user with id: %d", id),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return nil, errors.New("failed to read user info")
	}
	return user, nil
}

func (s *serv) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		_, errTx = s.userRepository.Get(ctx, user.ID)
		if errTx != nil {
			return errTx
		}

		errTx = s.userRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, &model.Log{
			Text: fmt.Sprintf("Updated user with id: %d", user.ID),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		if errors.Is(err, ErrUserExists) {
			return ErrUserExists
		}
		return errors.New("failed to update user info")
	}
	return nil
}

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		_, errTx = s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.userRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.logRepository.Log(ctx, &model.Log{
			Text: fmt.Sprintf("Deleted user with id: %d", id),
		})
		if errTx != nil {
			return errTx
		}

		return nil
	})
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
