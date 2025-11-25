package main

import (
	"ecommerce/dependency"
	"ecommerce/internal/routes"
	"log"

	"github.com/labstack/echo/v4"
)

// TODO: CRUD Produtos - SQL, Rotas, Model, Controller, Service, Repository
func main() {

	// Load Dependencies
	container := dependency.NewContainer()

	e := echo.New()
	e.Renderer = container.Renderer
	e.Static("/static", "view/static")

	routes.SetUpRoutes(e, container.UserController, container.ProductController)

	log.Println("🚀 Servidor iniciando na porta :8080")
	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
