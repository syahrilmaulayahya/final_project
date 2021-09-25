package respons

import (
	"final_project/business/products"
	"time"
)

type ProductResponse struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Total_Stock int       `json:"total_stock"`
	Price       float64   `json:"price"`
	Picture_url string    `json:"picture_url"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func FromDomain(domain products.ProductDomain) ProductResponse {
	return ProductResponse{
		ID:          domain.ID,
		Code:        domain.Code,
		Name:        domain.Name,
		Total_Stock: domain.Total_Stock,
		Price:       domain.Price,
		Picture_url: domain.Picture_url,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ListFromDomain(data []products.ProductDomain) (result []ProductResponse) {
	result = []ProductResponse{}
	for _, products := range data {
		result = append(result, FromDomain(products))
	}
	return
}
