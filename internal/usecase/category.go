package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
)

type categoryUsecase struct {
	repo domainCat.Repository
}

func NewCategoryService(repo domainCat.Repository) domainCat.Usecase {
	return &categoryUsecase{repo}
}

func (s *categoryUsecase) GetAll(ctx context.Context) ([]domainCat.Category, error) {
	categories, err := s.repo.FindAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("categoryUsecase: Failed to get all categories : %w", err)
	}

	return categories, nil
}

func (s *categoryUsecase) GetByID(ctx context.Context, id uint64) (*domainCat.Category, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domainCat.ErrNotFound) {
			return nil, domainCat.ErrNotFound
		}

		return nil, fmt.Errorf("categoryUsecase: Failed to get category by id = %d : %w", id, err)
	}
	return category, nil
}

func (s *categoryUsecase) Create(ctx context.Context, data *domainCat.CreateCatData) (*domainCat.Category, error) {
	if data.Name == "" {
		return nil, domainCat.ErrNameEmpty
	}

	newSlug := generateSlug(data.Name)

	createdCat, err := s.repo.Create(ctx, &domainCat.Category{
		Name: data.Name,
		Slug: newSlug,
	})

	if err != nil {
		return nil, fmt.Errorf("categoryUsecase: Failed to create new category : %w", err)
	}

	return createdCat, nil
}

func (s *categoryUsecase) Update(ctx context.Context, data *domainCat.UpdateCatData) (*domainCat.Category, error) {
	if data.Name == "" {
		return nil, domainCat.ErrNameEmpty
	}

	newSlug := generateSlug(data.Name)

	category := &domainCat.Category{
		ID:   data.ID,
		Name: data.Name,
		Slug: newSlug,
	}

	updatedCat, err := s.repo.Update(ctx, category)
	if err != nil {
		if errors.Is(err, domainCat.ErrNotFound) {
			return nil, domainCat.ErrNotFound
		}
		return nil, fmt.Errorf("categoryUsecase: Failed to update category id = %d : %w", data.ID, err)
	}

	return updatedCat, nil
}

// func (s *categoryUsecase) Delete(ctx context.Context, id uint64) error {
// 	if err := s.repo.Delete(ctx, id); err != nil {
// 		if errors.Is(err, domainCat.ErrNotFound) {
// 			return domainCat.ErrNotFound
// 		}
// 		return fmt.Errorf("categoryUsecase: Failed to delete category id = %d : %w", id, err)
// 	}
// 	return nil
// }

func generateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
