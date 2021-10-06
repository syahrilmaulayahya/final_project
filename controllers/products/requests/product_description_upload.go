package requests

import "final_project/business/products"

type Product_descriptionUpload struct {
	ProductID   int    `json:"productid"`
	Description string `json:"description"`
}

func (product_description *Product_descriptionUpload) ToDomain() products.Product_descriptionDomain {
	return products.Product_descriptionDomain{
		ProductID:   product_description.ProductID,
		Description: product_description.Description,
	}
}
