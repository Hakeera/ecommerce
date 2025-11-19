package repository

import (
	"erp/internal/model"

	"gorm.io/gorm"
)

// ProductRepository defines all database operations related to the Product entity.
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository returns a new instance of ProductRepository.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// FindByCategory retrieves products filtered by category.
func (r *ProductRepository) FindByCategory(category string) ([]model.Product, error) {
	var products []model.Product

	err := r.DB.
		Where("category = ?", category).
		Find(&products).Error

	return products, err
}
