package product

import (
	"final_project/business/products"
	"time"
)

type Product struct {
	ID          int    `gorm:"primaryKey"`
	Code        string `gorm:"unique"`
	Name        string
	Total_Stock int
	Price       float64
	Picture_url string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (product *Product) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		ID:          product.ID,
		Code:        product.Code,
		Name:        product.Name,
		Total_Stock: product.Total_Stock,
		Price:       product.Price,
		Picture_url: product.Picture_url,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func FromDomain(domain products.ProductDomain) Product {
	return Product{
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

func ToListDomain(data []Product) (result []products.ProductDomain) {
	result = []products.ProductDomain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
