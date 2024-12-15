package model

import (
	"database/sql"
	"time"
)

// UserRole type is the type for user role.
type UserRole string

// UserRole constants
const (
	UserRoleUser  UserRole = "USER"
	UserRoleAdmin UserRole = "ADMIN"
)

// User type is the main structure for user.
type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      string
	Version   int
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// UserCreate type is the structure for creating user.
type UserCreate struct {
	ID              string
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

// UserUpdate type is the structure for updating user info.
type UserUpdate struct {
	ID      string
	Name    sql.NullString
	Email   sql.NullString
	Role    sql.NullString
	Version sql.NullInt32
}
