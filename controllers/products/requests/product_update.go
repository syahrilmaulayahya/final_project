package requests

import "final_project/business/products"

type ProductUpdate struct {
	ID          int     `json:"id"`
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Picture_url string  `json:"picture_url"`
}

func (productUpdate ProductUpdate) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		ID:          productUpdate.ID,
		Code:        productUpdate.Code,
		Name:        productUpdate.Name,
		Price:       productUpdate.Price,
		Picture_url: productUpdate.Picture_url,
	}
}
