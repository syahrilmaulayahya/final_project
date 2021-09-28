package respons

import (
	"final_project/business/transactions"
	"time"
)

type Shopping_CartResponse struct {
	ID        int `json:"id"`
	UserID    int `json:"userid"`
	ProductID int `json:"productid"`
	SizeID    int `json:"sizeid"`
	Quantity  int `json:"quantity"`
	Price     float64
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain transactions.Shopping_CartDomain) Shopping_CartResponse {
	return Shopping_CartResponse{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ProductID: domain.ProductID,
		SizeID:    domain.SizeID,
		Quantity:  domain.Quantity,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
