package routes

import (
	"erp/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo, userController *controller.UserController, productController *controller.ProductController) {
	e.GET("/", controller.HomePage)

	// Users Routes
	e.GET("/users", userController.ListUsers)         // Render list of users
	e.GET("/users/new", userController.NewUserForm)   // Show "new user" form
	e.POST("/users", userController.CreateUser)       // Create a new user
	e.DELETE("/users/:id", userController.DeleteUser) // Delete user by ID

	// Product Routes
	e.GET("/products/page", productController.ProductPage)              // Render products page
	e.GET("/products/by-category", productController.ProductByCategory) // Filter Products by Categories
}
