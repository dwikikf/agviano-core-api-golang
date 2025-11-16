package usecase

import (
	"context"

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
		return nil, err
	}
	return products, nil
}

func (s *ProductUsecase) GetByID(ctx context.Context, id uint64) (*domainProd.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductUsecase) Create(ctx context.Context, data *domainProd.CreateProductInput) (*domainProd.Product, error) {
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
		return nil, err
	}

	return createdProduct, nil
}

func (s *ProductUsecase) Update(ctx context.Context, data *domainProd.UpdateProductInput) (*domainProd.Product, error) {
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
		return nil, err
	}

	return updatedProd, nil
}

func (s *ProductUsecase) Delete(ctx context.Context, id uint64) error {
	return nil
}
