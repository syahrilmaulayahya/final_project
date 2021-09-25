package products

import (
	"context"
	"errors"
	"time"
)

type ProductUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewProductUseCase(repo Repository, timeOut time.Duration) UseCase {
	return &ProductUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
	}
}

func (uc *ProductUseCase) Get(ctx context.Context) ([]ProductDomain, error) {
	product, err := uc.Repo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (uc *ProductUseCase) UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error) {
	if domain.Name == "" {
		return Product_typeDomain{}, errors.New("product type name empty")
	}
	productType, err := uc.Repo.UploadType(ctx, domain)
	if err != nil {
		return Product_typeDomain{}, err
	}
	return productType, nil
}
