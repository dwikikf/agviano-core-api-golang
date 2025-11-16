package repository

import (
	"context"

	domainProd "github.com/dwikikf/agviano-core-api-golang/internal/domain/product"
	"github.com/dwikikf/agviano-core-api-golang/internal/errs"
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

	err := r.db.WithContext(ctx).Preload("Category").Find(&products).Error
	if err != nil {
		return nil, errs.TranslateError(err)
	}
	return products, nil
}

func (r *productGormRepository) FindByID(ctx context.Context, id uint64) (*domainProd.Product, error) {
	var prod domainProd.Product

	err := r.db.WithContext(ctx).Preload("Category").First(&prod, id).Error
	if err != nil {
		return nil, errs.TranslateError(err)
	}
	return &prod, nil
}

func (r *productGormRepository) Create(ctx context.Context, prod *domainProd.Product) (*domainProd.Product, error) {
	err := r.db.WithContext(ctx).Create(prod).Error
	if err != nil {
		return nil, errs.TranslateError(err)
	}
	return prod, nil
}

func (r *productGormRepository) Update(ctx context.Context, prod *domainProd.Product) (*domainProd.Product, error) {
	res := r.db.WithContext(ctx).Model(&domainProd.Product{}).Where("id = ?", prod.ID).Updates(prod)

	if res.Error != nil {
		return nil, errs.TranslateError(res.Error)
	}

	if res.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return prod, nil
}

func (r *productGormRepository) Delete(ctx context.Context, id uint64) error {
	return nil
}
