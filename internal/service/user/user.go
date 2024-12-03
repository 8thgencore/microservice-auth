package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/8thgencore/microservice-auth/internal/model"
	"github.com/8thgencore/microservice-common/pkg/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
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
func (s *serv) Create(ctx context.Context, user *model.UserCreate) (string, error) {
	// Check if passwords match
	if user.Password != user.PasswordConfirm {
		return "", ErrPasswordsMismatch
	}

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Generate a UUIDv7 for the user
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	user.ID = uuidv7.String()

	// Create the user
	var id string
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
			return "", ErrUserNameExists
		}
		if errors.Is(err, ErrUserEmailExists) {
			return "", ErrUserEmailExists
		}

		return "", ErrUserCreate
	}

	return id, nil
}

// Get retrieves a user by their ID.
func (s *serv) Get(ctx context.Context, id string) (*model.User, error) {
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
	var currentUser *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		currentUser, errTx = s.userRepository.Get(ctx, user.ID)
		if errTx != nil {
			return errTx
		}

		currentUser.Version = currentUser.Version + 1
		convertedVersion, err := safeIntToInt32(currentUser.Version)
		if err != nil {
			logger.Error("version value out of range for int32: %d", zap.Int("method", currentUser.Version))
			return err
		}
		user.Version = sql.NullInt32{Int32: convertedVersion, Valid: true}

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

	if err := s.tokenRepository.SetTokenVersion(ctx, currentUser.ID, currentUser.Version); err != nil {
		return ErrUserUpdate
	}

	return nil
}

// Delete handles the deletion of a user.
func (s *serv) Delete(ctx context.Context, id string) error {
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

// Safe conversion function
func safeIntToInt32(value int) (int32, error) {
	if value > math.MaxInt32 || value < math.MinInt32 {
		return 0, fmt.Errorf("value out of range for int32: %d", value)
	}
	return int32(value), nil
}

// logUserAction is a helper function to log actions performed on a user.
func (s *serv) logUserAction(ctx context.Context, action string, userID string) error {
	// Generate a UUIDv7 for the user
	uuidv7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	return s.logRepository.Log(ctx, &model.Log{
		ID:   uuidv7.String(),
		Text: fmt.Sprintf("%s with id: %s", action, userID),
	})
}
