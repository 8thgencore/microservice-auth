package model

import (
	"database/sql"
	"time"
)

// User type is the main structure for user.
type User struct {
	ID        string
	Name      string
	Email     string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// UserCreate type is the structure for creating user.
type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
}

// UserUpdate type is the structure for updating user info.
type UserUpdate struct {
	ID    string
	Name  sql.NullString
	Email sql.NullString
	Role  sql.NullString
}
