package requests

import "final_project/business/products"

type StockUpload struct {
	ID    int `json:"id"`
	Stock int `json:"stock"`
}

func (stock *StockUpload) ToDomain() products.SizeDomain {
	return products.SizeDomain{
		ID:    stock.ID,
		Stock: stock.Stock,
	}
}
