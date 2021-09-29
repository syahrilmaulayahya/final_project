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

func (uc *TransactionUseCase) DetailSC(ctx context.Context, id int) ([]Shopping_CartDomain, error) {
	transactions, err := uc.Repo.DetailSC(ctx, id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (uc *TransactionUseCase) AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error) {
	if domain.Name == "" {
		return Payment_MethodDomain{}, errors.New("payment method name is empty")
	}
	paymentMethod, err := uc.Repo.AddPM(ctx, domain)
	if err != nil {
		return Payment_MethodDomain{}, err
	}
	return paymentMethod, nil
}

func (uc *TransactionUseCase) GetPM(ctx context.Context) ([]Payment_MethodDomain, error) {
	payment_method, err := uc.Repo.GetPM(ctx)
	if err != nil {
		return nil, err
	}
	return payment_method, nil
}

func (uc *TransactionUseCase) AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error) {
	if domain.Name == "" {
		domain.Name = "J&T"
	}
	if domain.Shipment_Type == "" {
		domain.Shipment_Type = "Regular"
	}
	if domain.Shipment_Price == 0 {
		domain.Shipment_Price = 50000
	}
	shipment, err := uc.Repo.AddShipment(ctx, domain)
	if err != nil {
		return ShipmentDomain{}, err
	}
	return shipment, nil
}

// func (uc *TransactionUseCase) GetShipment(ctx context.Context) ([]ShipmentDomain, error) {
// 	shipment, err := uc.Repo.GetShipment(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return shipment, nil
// }
