package users

import (
	"final_project/business/transactions"
	"time"
)

type Shopping_Cart struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	ProductID int
	SizeID    int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (shopping_cart *Shopping_Cart) ToDomain() transactions.Shopping_CartDomain {
	return transactions.Shopping_CartDomain{
		ID:        shopping_cart.ID,
		UserID:    shopping_cart.UserID,
		ProductID: shopping_cart.ProductID,
		SizeID:    shopping_cart.SizeID,
		Quantity:  shopping_cart.Quantity,
		CreatedAt: shopping_cart.CreatedAt,
		UpdatedAt: shopping_cart.UpdatedAt,
	}
}

func FromDomain(domain transactions.Shopping_CartDomain) Shopping_Cart {
	return Shopping_Cart{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ProductID: domain.ProductID,
		SizeID:    domain.SizeID,
		Quantity:  domain.Quantity,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
