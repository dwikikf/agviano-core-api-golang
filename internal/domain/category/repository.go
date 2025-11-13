package domain

import (
	"context"
)

type Repository interface {
	FindAll(ctx context.Context) ([]Category, error)
	FindByID(ctx context.Context, id uint64) (*Category, error)
	Create(ctx context.Context, category *Category) (*Category, error)
	Update(ctx context.Context, category *Category) (*Category, error)
	// Delete(ctx context.Context, id uint64) error
}
