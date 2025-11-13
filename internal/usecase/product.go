package usecase

import (
	"context"
	"errors"
	"fmt"

	domainCat "github.com/dwikikf/agviano-core-api-golang/internal/domain/category"
	domainProd "github.com/dwikikf/agviano-core-api-golang/internal/domain/product"
)

type ProductUsecase struct {
	repo    domainProd.Repository
	catRepo domainCat.Repository
}

func NewProductService(repo domainProd.Repository, catRepo domainCat.Repository) domainProd.Usecase {
	return &ProductUsecase{repo, catRepo}
}

func (s *ProductUsecase) GetAll(ctx context.Context) ([]domainProd.Product, error) {
	products, err := s.repo.FindAll(ctx)
	if err != nil {
		if errors.Is(err, domainProd.ErrNotFound) {
			return nil, domainProd.ErrNotFound
		}
		return nil, fmt.Errorf("ProductUsecase: Failed to get all products : %w", err)
	}
	return products, nil
}

func (s *ProductUsecase) GetByID(ctx context.Context, id uint64) (*domainProd.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, domainProd.ErrNotFound) {
			return nil, domainProd.ErrNotFound
		}
		return nil, fmt.Errorf("ProductUsecase: Failed to get product by id = %d : %w", id, err)
	}

	return product, nil
}

func (s *ProductUsecase) Create(ctx context.Context, data *domainProd.CreateProductInput) (*domainProd.Product, error) {
	if data.Name == "" {
		return nil, domainProd.ErrNameEmpty
	}

	if data.Slug == "" {
		return nil, domainProd.ErrSlugEmpty
	}

	if data.Price < 0 {
		return nil, domainProd.ErrInvalidPrice
	}

	if data.CategoryID != 0 {
		_, err := s.catRepo.FindByID(ctx, data.CategoryID)

		if errors.Is(err, domainCat.ErrNotFound) {

			return nil, errors.New("referenced Category ID not found")
		}
		if err != nil {

			return nil, fmt.Errorf("ProductUsecase: failed to check Category dependency: %w", err)
		}
	}

	createdProduct, err := s.repo.Create(ctx, &domainProd.Product{
		CategoryID:  data.CategoryID,
		Name:        data.Name,
		Slug:        data.Slug,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		ImageURL:    data.ImageURL,
		IsActive:    data.IsActive,
	})

	if err != nil {
		return nil, fmt.Errorf("ProductUsecase: Failed to create new product : %w", err)
	}

	return createdProduct, nil
}

func (s *ProductUsecase) Update(ctx context.Context, data *domainProd.UpdateProductInput) (*domainProd.Product, error) {
	if data.Name == "" {
		return nil, domainProd.ErrNameEmpty
	}

	if data.Slug == "" {
		return nil, domainProd.ErrSlugEmpty
	}

	if data.Price < 0 {
		return nil, domainProd.ErrInvalidPrice
	}

	if data.CategoryID != 0 {
		_, err := s.catRepo.FindByID(ctx, data.CategoryID)

		if errors.Is(err, domainCat.ErrNotFound) {

			return nil, errors.New("referenced Category ID not found")
		}
		if err != nil {

			return nil, fmt.Errorf("ProductUsecase: failed to check Category dependency: %w", err)
		}
	}

	prod := &domainProd.Product{
		ID:          data.ID,
		CategoryID:  data.CategoryID,
		Name:        data.Name,
		Slug:        data.Slug,
		Description: data.Description,
		Price:       data.Price,
		Stock:       data.Stock,
		ImageURL:    data.ImageURL,
		IsActive:    data.IsActive,
	}

	updatedProd, err := s.repo.Update(ctx, prod)
	if err != nil {
		if errors.Is(err, domainProd.ErrNotFound) {
			return nil, domainProd.ErrNotFound
		}
		return nil, fmt.Errorf("ProductUsecase: Failed to update product id = %d : %w", data.ID, err)
	}

	return updatedProd, nil
}

func (s *ProductUsecase) Delete(ctx context.Context, id uint64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		if errors.Is(err, domainProd.ErrNotFound) {
			return domainProd.ErrNotFound
		}
		return fmt.Errorf("ProductUsecase: Failed to delete product id = %d : %w", id, err)
	}
	return nil
}
