package model

import (
	"errors"
	"strings"
	"time"
)

// User represents a system user entity.
// It stores authentication data, role information,
// and creation/update timestamps.

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Username     string    `gorm:"unique;not null" form:"username"`
	Email        string    `gorm:"unique;not null" form:"email"`
	PasswordHash string    `gorm:"not null" form:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName overrides the default GORM table name.
func (User) TableName() string {
	return "users"
}

// IsValid performs basic validation checks for the User entity.
// It ensures username length and email format are valid.
func (u *User) IsValid() error {
	if len(u.Username) < 3 {
		return errors.New("username must have at least 3 characters")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}
