package routes

import (
	"erp/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo, userController *controller.UserController) {
	e.GET("/", controller.HomePage)

	// Render list of users
	e.GET("/users", userController.ListUsers)

	// Show "new user" form
	e.GET("/users/new", userController.NewUserForm)

	// Create a new user
	e.POST("/users", userController.CreateUser)

	// Delete user by ID
	e.DELETE("/users/:id", userController.DeleteUser)
}
