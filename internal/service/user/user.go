package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/8thgencore/microservice-auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// ErrUserNameExists - custom error for user name duplicate.
var ErrUserNameExists = errors.New("user with provided name already exists")

// ErrUserEmailExists - custom error for email duplicate.
var ErrUserEmailExists = errors.New("user with provided email already exists")

// ErrUserEmailExists - custom error if user not found.
var ErrUserNotFound = errors.New("user not found")

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

		return s.logUserAction(ctx, "Created user", id)
	})
	if err != nil {
		if errors.Is(err, ErrUserNameExists) {
			return 0, ErrUserNameExists
		}
		if errors.Is(err, ErrUserEmailExists) {
			return 0, ErrUserEmailExists
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

		return s.logUserAction(ctx, "Read user info", id)
	})
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, errors.New("failed to read user info")
	}
	return user, nil
}

func (s *serv) Update(ctx context.Context, user *model.UserUpdate) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		_, errTx := s.userRepository.Get(ctx, user.ID)
		if errTx != nil {
			return errTx
		}

		errTx = s.userRepository.Update(ctx, user)
		if errTx != nil {
			return errTx
		}

		return s.logUserAction(ctx, "Updated user", user.ID)
	})
	if err != nil {
		if errors.Is(err, ErrUserNameExists) {
			return ErrUserNameExists
		}
		if errors.Is(err, ErrUserEmailExists) {
			return ErrUserEmailExists
		}
		return errors.New("failed to update user info")
	}
	return nil
}

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		_, errTx := s.userRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		errTx = s.userRepository.Delete(ctx, id)
		if errTx != nil {
			return errTx
		}

		return s.logUserAction(ctx, "Deleted user", id)
	})
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

// logUserAction is a helper function to log actions performed on a user.
func (s *serv) logUserAction(ctx context.Context, action string, userID int64) error {
	return s.logRepository.Log(ctx, &model.Log{
		Text: fmt.Sprintf("%s with id: %d", action, userID),
	})
}
