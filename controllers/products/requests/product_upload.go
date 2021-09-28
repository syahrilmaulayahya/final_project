package requests

import "final_project/business/products"

type ProductUpload struct {
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Picture_url    string  `json:"picture_url"`
	Product_typeID int     `json:"product_typeid"`
}

func (product *ProductUpload) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		Code:           product.Code,
		Name:           product.Name,
		Price:          product.Price,
		Picture_url:    product.Picture_url,
		Product_typeID: product.Product_typeID,
	}
}
