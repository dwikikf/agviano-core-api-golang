package domain

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Product, error)
	FindByID(ctx context.Context, id uint64) (*Product, error)
	Create(ctx context.Context, product *Product) (*Product, error)
	Update(ctx context.Context, product *Product) (*Product, error)
	Delete(ctx context.Context, id uint64) error
}
