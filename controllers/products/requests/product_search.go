package requests

import "final_project/business/products"

type Product_Search struct {
	Name string `json:"name"`
}

func (product_search *Product_Search) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		Name: product_search.Name,
	}
}
