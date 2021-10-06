package requests

import (
	"final_project/business/products"
)

type Product_Filter struct {
	Product_typeID int `json:"product_typeid"`
}

func (product_filter *Product_Filter) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		Product_typeID: product_filter.Product_typeID,
	}
}
