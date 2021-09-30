package requests

import "final_project/business/transactions"

type Shopping_CartCheckout struct {
	ID     int `json:"id"`
	UserID int `json:"userid"`
}

func (shopping_cartCheckout *Shopping_CartCheckout) ToDomain() transactions.Shopping_CartDomain {
	return transactions.Shopping_CartDomain{
		ID:     shopping_cartCheckout.ID,
		UserID: shopping_cartCheckout.UserID,
	}
}
