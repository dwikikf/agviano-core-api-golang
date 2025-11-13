package repository

import (
	"context"
	"errors"
	"fmt"

	domainProd "github.com/dwikikf/agviano-core-api-golang/internal/domain/product"
	"gorm.io/gorm"
)

type productGormRepository struct {
	db *gorm.DB
}

func NewProductGormRepository(db *gorm.DB) domainProd.Repository {
	return &productGormRepository{db}
}

func (r *productGormRepository) FindAll(ctx context.Context) ([]domainProd.Product, error) {
	var products []domainProd.Product
	if err := r.db.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, fmt.Errorf("productGormRepository: Failed to find all : %w", err)
	}
	return products, nil
}

func (r *productGormRepository) FindByID(ctx context.Context, id uint64) (*domainProd.Product, error) {
	var prod domainProd.Product
	if err := r.db.WithContext(ctx).First(&prod, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainProd.ErrNotFound
		}
		return nil, fmt.Errorf("productGormRepository: Failed to find by id = %d : %w", id, err)
	}
	return &prod, nil
}

func (r *productGormRepository) Create(ctx context.Context, prod *domainProd.Product) (*domainProd.Product, error) {
	if err := r.db.WithContext(ctx).Create(prod).Error; err != nil {
		return nil, fmt.Errorf("productGormRepository: Failed to create new product : %w", err)
	}
	return prod, nil
}

func (r *productGormRepository) Update(ctx context.Context, prod *domainProd.Product) (*domainProd.Product, error) {
	result := r.db.WithContext(ctx).Model(&domainProd.Product{}).Where("id = ?", prod.ID).Updates(prod)

	if result.Error != nil {
		return nil, fmt.Errorf("productGormRepository: Failed to update product id = %d : %w", prod.ID, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, domainProd.ErrNotFound
	}

	return prod, nil
}

func (r *productGormRepository) Delete(ctx context.Context, id uint64) error {
	result := r.db.WithContext(ctx).Delete(&domainProd.Product{}, id)

	if result.Error != nil {
		return fmt.Errorf("productGormRepository: Failed to delete product id = %d : %w", id, result.Error)
	}

	if result.RowsAffected == 0 {
		return domainProd.ErrNotFound
	}

	return nil
}
