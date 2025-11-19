package service

import (
	"erp/internal/model"
	"erp/internal/repository"
	"errors"
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

	return s.ProductRepository.FindByCategory(category)
}
