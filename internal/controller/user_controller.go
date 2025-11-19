package controller

import (
	"erp/internal/model"
	"erp/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService *service.UserService
}

// ListUsers handles GET /users
// It retrieves all users and renders the HTMX users list partial.
func (uc *UserController) ListUsers(c echo.Context) error {
	users, err := uc.UserService.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to load users")
	}

	// Preparar dados para o template
	dataMap := map[string]any{
		"Users": users,
	}

	return c.Render(http.StatusOK, "users_list", dataMap)
}

// NewUserForm handles GET /users/new
// It renders an empty HTMX form for creating a new user.
func (uc *UserController) NewUserForm(c echo.Context) error {

	return c.Render(http.StatusOK, "users_form", nil)
}

// CreateUser handles POST /users
// It binds form data, validates the User model, delegates creation to the service,
// and triggers HTMX redirect upon success.
func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(model.User)

	if err := c.Bind(user); err != nil {
		return c.Render(http.StatusOK, "user_error", err)
	}

	// Validation
	if err := user.IsValid(); err != nil {
		return c.Render(http.StatusOK, "user_error", err)
	}

	if err := uc.UserService.Create(user); err != nil {
		return c.Render(http.StatusOK, "user_error", err)
	}

	users, _ := uc.UserService.GetAll()

	// Preparar dados para o template
	dataMap := map[string]any{
		"Users": users,
	}

	// Rerender users_list with new usrs
	return c.Render(http.StatusOK, "users_list", dataMap)
}

// DeleteUser handles DELETE /users/:id
// It removes a user and re-renders the updated HTMX partial list.
func (uc *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := uc.UserService.Delete(id); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete user")
	}

	users, _ := uc.UserService.GetAll()

	// Preparar dados para o template
	dataMap := map[string]any{
		"Users": users,
	}

	return c.Render(http.StatusOK, "users_list", dataMap)
}
