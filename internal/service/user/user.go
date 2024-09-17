package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/8thgencore/microservice-auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// Errors
var (
	ErrUserNameExists     = errors.New("user with provided name already exists")
	ErrUserEmailExists    = errors.New("user with provided email already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrPasswordsMismatch  = errors.New("passwords don't match")
	ErrPasswordProcessing = errors.New("failed to process password")
	ErrUserCreate         = errors.New("failed to create user")
	ErrUserRead           = errors.New("failed to read user info")
	ErrUserUpdate         = errors.New("failed to update user info")
	ErrUserDelete         = errors.New("failed to delete user")
)

// Create handles the creation of a new user.
func (s *serv) Create(ctx context.Context, user *model.UserCreate) (int64, error) {
	// Check if passwords match
	if user.Password != user.PasswordConfirm {
		return 0, ErrPasswordsMismatch
	}

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword

	// Create the user
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
		return 0, ErrUserCreate
	}

	return id, nil
}

// Get retrieves a user by their ID.
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
		return nil, ErrUserRead
	}

	return user, nil
}

// Update handles the updating of a user's information.
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
		if errors.Is(err, ErrUserNotFound) {
			return ErrUserNotFound
		}
		if errors.Is(err, ErrUserNameExists) {
			return ErrUserNameExists
		}
		if errors.Is(err, ErrUserEmailExists) {
			return ErrUserEmailExists
		}
		return ErrUserUpdate
	}

	return nil
}

// Delete handles the deletion of a user.
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
		if errors.Is(err, ErrUserNotFound) {
			return ErrUserNotFound
		}
		return ErrUserDelete
	}

	return nil
}

// Helper function for password hashing
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrPasswordProcessing
	}
	return string(hashedPassword), nil
}

// logUserAction is a helper function to log actions performed on a user.
func (s *serv) logUserAction(ctx context.Context, action string, userID int64) error {
	return s.logRepository.Log(ctx, &model.Log{
		Text: fmt.Sprintf("%s with id: %d", action, userID),
	})
}
