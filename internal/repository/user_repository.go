package repository

import (
	"erp/internal/model"

	"gorm.io/gorm"
)

// UserRepository defines all database operations related to the User entity.
// It isolates persistence logic from the service layer, making the code easier
// to maintain, test, and replace (e.g., switch DB engine).
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository returns a new instance of UserRepository with a database connection.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// Create inserts a new User record into the database.
func (r *UserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

// FindAll returns all users stored in the database.
func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	return users, err
}

// Delete removes a user by ID.
// If no record is found, GORM returns an error.
func (r *UserRepository) Delete(id string) error {
	return r.DB.Delete(&model.User{}, id).Error
}
