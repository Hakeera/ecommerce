package service

import (
	"erp/internal/model"
	"erp/internal/repository"
	"errors"
	"fmt"
)

// UserService contains business rules and application-level logic.
// It orchestrates operations, validates inputs, and communicates with the repository layer.
type UserService struct {
	UserRepository *repository.UserRepository
}

// NewUserService returns a new instance of UserService.
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

// Create validates the User data and forwards the operation to the repository.
// This is where domain rules and constraints should be applied.
func (s *UserService) Create(user *model.User) error {
	fmt.Println("SERVICE CREATE USER")
	if err := user.IsValid(); err != nil {
		return err
	}

	if user.PasswordHash == "" {
		return errors.New("password cannot be empty")
	}

	return s.UserRepository.Create(user)
}

// GetAll retrieves all users from the repository.
func (s *UserService) GetAll() ([]model.User, error) {
	return s.UserRepository.FindAll()
}

// Delete removes a user by ID.
// Business logic or checks (e.g., preventing deletion of admin) can be placed here.
func (s *UserService) Delete(id string) error {
	return s.UserRepository.Delete(id)
}
