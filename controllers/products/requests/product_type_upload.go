package requests

import (
	"final_project/business/products"
)

type ProductTypeUpload struct {
	Name string `json:"name"`
}

func (product_type *ProductTypeUpload) ToDomain() products.Product_typeDomain {
	return products.Product_typeDomain{
		Name: product_type.Name,
	}
}
