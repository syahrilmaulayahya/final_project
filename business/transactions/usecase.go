package transactions

import (
	"context"
	"errors"
	"final_project/app/middleware"
	"time"
)

type TransactionUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JwtToken       middleware.ConfigJWT
}

func NewTransactionUseCase(repo Repository, timeOut time.Duration, token middleware.ConfigJWT) UseCase {
	return &TransactionUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
		JwtToken:       token,
	}
}
func (uc *TransactionUseCase) Add(ctx context.Context, domain Shopping_CartDomain) (Shopping_CartDomain, error) {
	if domain.ProductID == 0 {
		return Shopping_CartDomain{}, errors.New("product id is empty")
	}
	if domain.SizeID == 0 {
		return Shopping_CartDomain{}, errors.New("size id is empty")
	}
	if domain.Quantity < 0 {
		return Shopping_CartDomain{}, errors.New("invalid quantity")
	}
	transaction, err := uc.Repo.Add(ctx, domain)
	if err != nil {
		return Shopping_CartDomain{}, err
	}
	return transaction, nil

}
