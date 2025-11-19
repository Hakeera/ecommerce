package controller

import (
	"erp/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductService *service.ProductService
}

func (pc *ProductController) ProductPage(c echo.Context) error {
	categories := []string{"Trabalho", "Escolar", "Escritório", "Esporte"}

	dataMap := map[string]any{
		"Categories": categories,
	}

	return c.Render(http.StatusOK, "product_page", dataMap)
}

func (pc *ProductController) ProductByCategory(c echo.Context) error {
	category := c.QueryParam("category")
	if category == "" {
		return c.String(http.StatusBadRequest, "Categoria não informada")
	}

	products, err := pc.ProductService.GetProductsByCategory(category)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	dataMap := map[string]any{
		"Products": products,
	}

	return c.Render(http.StatusOK, "product_list", dataMap)
}
