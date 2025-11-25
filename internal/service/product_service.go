package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
	"errors"
	"fmt"
)

// ProductService contains business rules and application-level logic.
type ProductService struct {
	ProductRepository *repository.ProductRepository
}

// NewProductService creates a new ProductService instance.
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: repo,
	}
}

// GetProductsByCategory fetches products filtered by category.
func (s *ProductService) GetProductsByCategory(category string) ([]model.Product, error) {
	if category == "" {
		return nil, errors.New("categoria não pode ser vazia")
	}

	products, err := s.ProductRepository.FindByCategory(category)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar produtos por categoria: %w", err)
	}

	return products, nil
}
