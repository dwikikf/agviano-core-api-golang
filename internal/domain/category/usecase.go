package domain

import "context"

type Usecase interface {
	GetAll(ctx context.Context, page, size int) ([]Category, int64, error)
	GetByID(ctx context.Context, id uint64) (*Category, error)
	Create(ctx context.Context, data *CreateCatData) (*Category, error)
	Update(ctx context.Context, data *UpdateCatData) (*Category, error)
	// Delete(ctx context.Context, id uint64) error
}

type CreateCatData struct {
	ID   uint64
	Name string
	Slug string
}

type UpdateCatData struct {
	ID   uint64
	Name string
	Slug string
}
