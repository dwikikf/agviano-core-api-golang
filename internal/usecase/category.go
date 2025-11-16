package usecase

import (
	"context"
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
		return nil, err
	}

	return categories, nil
}

func (s *categoryUsecase) GetByID(ctx context.Context, id uint64) (*domainCat.Category, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryUsecase) Create(ctx context.Context, data *domainCat.CreateCatData) (*domainCat.Category, error) {
	created, err := s.repo.Create(ctx, &domainCat.Category{
		Name: data.Name,
		Slug: generateSlug(data.Name),
	})

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *categoryUsecase) Update(ctx context.Context, data *domainCat.UpdateCatData) (*domainCat.Category, error) {
	newSlug := generateSlug(data.Name)

	category := &domainCat.Category{
		ID:   data.ID,
		Name: data.Name,
		Slug: newSlug,
	}

	updatedCat, err := s.repo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	return updatedCat, nil
}

// func (s *categoryUsecase) Delete(ctx context.Context, id uint64) error {
//
// 	return nil
// }

func generateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
