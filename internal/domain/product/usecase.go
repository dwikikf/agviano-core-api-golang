package domain

import "context"

type Usecase interface {
	GetAll(ctx context.Context, page, size int) ([]Product, int64, error)
	GetByID(ctx context.Context, id uint64) (*Product, error)
	Create(ctx context.Context, data *CreateProductInput) (*Product, error)
	Update(ctx context.Context, data *UpdateProductInput) (*Product, error)
	Delete(ctx context.Context, id uint64) error
}

type CreateProductInput struct {
	CategoryID  uint64
	Name        string
	Slug        string
	Description string
	Price       float64
	Stock       uint
	ImageURL    string
	IsActive    bool
}

type UpdateProductInput struct {
	ID          uint64
	CategoryID  uint64
	Name        string
	Slug        string
	Description string
	Price       float64
	Stock       uint
	ImageURL    string
	IsActive    bool
}
