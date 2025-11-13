package repository

import (
	"context"
	"errors"
	"fmt"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
	"gorm.io/gorm"
)

type categoryGormRepository struct {
	db *gorm.DB
}

func NewCategoryGormRepository(db *gorm.DB) domainCat.Repository {
	return &categoryGormRepository{db}
}

func (r *categoryGormRepository) FindAll(ctx context.Context) ([]domainCat.Category, error) {
	var categories []domainCat.Category
	if err := r.db.WithContext(ctx).Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("categoryGormRepository: Failed to find all : %w", err)
	}
	return categories, nil
}

func (r *categoryGormRepository) FindByID(ctx context.Context, id uint64) (*domainCat.Category, error) {
	var cat domainCat.Category
	if err := r.db.WithContext(ctx).First(&cat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainCat.ErrNotFound
		}
		return nil, fmt.Errorf("categoryGormRepository: Failed to find by id = %d : %w", id, err)
	}
	return &cat, nil
}

func (r *categoryGormRepository) Create(ctx context.Context, cat *domainCat.Category) (*domainCat.Category, error) {
	if err := r.db.WithContext(ctx).Create(cat).Error; err != nil {
		return nil, fmt.Errorf("categoryGormRepository: Failed to create new category : %w", err)
	}
	return cat, nil
}

func (r *categoryGormRepository) Update(ctx context.Context, cat *domainCat.Category) (*domainCat.Category, error) {
	result := r.db.WithContext(ctx).Model(&domainCat.Category{}).Where("id = ?", cat.ID).Updates(cat)

	if result.Error != nil {
		return nil, fmt.Errorf("categoryGormRepository: Failed to update category id = %d : %w", cat.ID, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, domainCat.ErrNotFound
	}

	return cat, nil
}

// func (r *categoryGormRepository) Delete(ctx context.Context, id uint64) error {
// 	result := r.db.WithContext(ctx).Delete(&domainCat.Category{}, id)
// 	if result.Error != nil {
// 		return fmt.Errorf("categoryGormRepository: Failed to delete category id = %d : %w", id, result.Error)
// 	}

// 	if result.RowsAffected == 0 {
// 		return domainCat.ErrNotFound
// 	}

// 	return nil
// }
