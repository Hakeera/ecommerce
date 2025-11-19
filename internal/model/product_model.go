package model

import (
	"github.com/lib/pq"
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" form:"name" json:"name"`
	Category    string `gorm:"not null" form:"category" json:"category"`
	ImageURL    string `gorm:"not null" form:"image_url" json:"image_url"`
	Description string `gorm:"type:text" form:"description" json:"description"`

	Price int `gorm:"not null" form:"price" json:"price"`
	Stock int `gorm:"not null" form:"stock" json:"stock"`

	Colors pq.StringArray `gorm:"type:text[]" form:"colors" json:"colors"`
	Sizes  pq.StringArray `gorm:"type:text[]" form:"sizes" json:"sizes"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
