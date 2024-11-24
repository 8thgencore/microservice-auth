package dao

import (
	"database/sql"
	"time"
)

// User type is the main structure for user.
type User struct {
	ID        string       `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      string       `db:"role"`
	Version   int          `db:"version"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

// AuthInfo type is the structure for user authentication data from storage.
type AuthInfo struct {
	ID       string `db:"id"`
	Username string `db:"name"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
