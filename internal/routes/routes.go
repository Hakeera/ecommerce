package routes

import (
	"erp/internal/controller"

	"github.com/labstack/echo/v4"
)

func SetUpRoutes(e *echo.Echo) {
	e.GET("/", controller.HomePage)

}
