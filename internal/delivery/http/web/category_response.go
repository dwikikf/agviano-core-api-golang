package web

import (
	"time"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
)

type CategoryResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToCategoryResponse(cat domainCat.Category) *CategoryResponse {
	return &CategoryResponse{
		ID:        cat.ID,
		Name:      cat.Name,
		Slug:      cat.Slug,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	}
}
