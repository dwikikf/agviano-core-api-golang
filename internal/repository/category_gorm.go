package repository

import (
	"context"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
	"github.com/dwikikf/agviano-core-api-golang/internal/errs"
	"gorm.io/gorm"
)

type categoryGormRepository struct {
	db *gorm.DB
}

func NewCategoryGormRepository(db *gorm.DB) domainCat.Repository {
	return &categoryGormRepository{db}
}

func (r *categoryGormRepository) FindAll(ctx context.Context, page, size int) ([]domainCat.Category, int64, error) {
	var categories []domainCat.Category
	var total int64

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	q := r.db.WithContext(ctx).Model(&domainCat.Category{})
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, errs.TranslateError(err)
	}

	offset := (page - 1) * size
	err := r.db.WithContext(ctx).Order("id desc").Limit(size).Offset(offset).Find(&categories).Error
	if err != nil {
		return nil, 0, errs.TranslateError(err)
	}

	return categories, total, nil
}

func (r *categoryGormRepository) FindByID(ctx context.Context, id uint64) (*domainCat.Category, error) {
	var cat domainCat.Category

	err := r.db.WithContext(ctx).First(&cat, id).Error
	if err != nil {
		return nil, errs.TranslateError(err)
	}

	return &cat, nil
}

func (r *categoryGormRepository) Create(ctx context.Context, cat *domainCat.Category) (*domainCat.Category, error) {
	err := r.db.WithContext(ctx).Create(cat).Error
	if err != nil {
		return nil, errs.TranslateError(err)
	}
	return cat, nil
}

func (r *categoryGormRepository) Update(ctx context.Context, cat *domainCat.Category) (*domainCat.Category, error) {
	res := r.db.WithContext(ctx).Model(&domainCat.Category{}).Where("id = ?", cat.ID).Updates(cat)

	if res.Error != nil {
		return nil, errs.TranslateError(res.Error)
	}

	if res.RowsAffected == 0 {
		return nil, errs.ErrNotFound
	}

	return cat, nil
}

// func (r *categoryGormRepository) Delete(ctx context.Context, id uint64) error {
//
// }
