package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {

	// Executa o template base
	return c.Render(http.StatusOK, "base", nil)
}
