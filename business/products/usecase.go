package products

import (
	"context"
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
