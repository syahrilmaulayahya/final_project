package requests

import "final_project/business/transactions"

type Shopping_CartAdd struct {
	UserID    int `json:"id"`
	ProductID int `json:"productid"`
	SizeID    int `json:"sizeid"`
	Quantity  int `json:"quantity"`
}

func (shopping_cartadd *Shopping_CartAdd) ToDomain() transactions.Shopping_CartDomain {
	return transactions.Shopping_CartDomain{
		UserID:    shopping_cartadd.UserID,
		ProductID: shopping_cartadd.ProductID,
		SizeID:    shopping_cartadd.SizeID,
		Quantity:  shopping_cartadd.Quantity,
	}
}
