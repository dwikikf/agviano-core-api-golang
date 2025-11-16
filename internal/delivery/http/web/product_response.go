package web

import (
	"time"

	domainProd "github.com/dwikikf/agviano-core-api-golang/internal/domain/product"
)

type ProductResponse struct {
	ID          uint64           `json:"id"`
	CategoryID  uint64           `json:"category_id"`
	Category    CategoryResponse `json:"category"`
	Name        string           `json:"name"`
	Slug        string           `json:"slug"`
	Description string           `json:"description"`
	Price       float64          `json:"price"`
	Stock       uint             `json:"stock"`
	ImageURL    string           `json:"image_url"`
	IsActive    bool             `json:"is_active"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

func ToProductResponse(prod domainProd.Product) *ProductResponse {
	return &ProductResponse{
		ID:          prod.ID,
		CategoryID:  prod.CategoryID,
		Category:    *ToCategoryResponse(prod.Category),
		Name:        prod.Name,
		Slug:        prod.Slug,
		Description: prod.Description,
		Price:       prod.Price,
		Stock:       prod.Stock,
		ImageURL:    prod.ImageURL,
		IsActive:    prod.IsActive,
		CreatedAt:   prod.CreatedAt,
		UpdatedAt:   prod.UpdatedAt,
	}
}
