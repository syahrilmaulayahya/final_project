package product

import (
	"final_project/business/products"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                  int    `gorm:"primaryKey"`
	Code                string `gorm:"unique"`
	Name                string
	Total_Stock         int
	Price               float64
	Picture_url         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Product_typeID      int
	Product_type        Product_type
	Product_description Product_description
	Review_Rating       []Review_Rating
	Size                []Size
}

type Review_Rating struct {
	ID        int `gorm:"primaryKey"`
	Review    string
	Rating    float32
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product_description struct {
	ProductID   int `gorm:"primaryKey"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Product_type struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Size struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Type      string
	Size      string
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (product *Product) ToDomain() products.ProductDomain {
	return products.ProductDomain{
		ID:                  product.ID,
		Code:                product.Code,
		Name:                product.Name,
		Price:               product.Price,
		Picture_url:         product.Picture_url,
		CreatedAt:           product.CreatedAt,
		UpdatedAt:           product.UpdatedAt,
		Product_typeID:      product.Product_typeID,
		Product_type:        product.Product_type,
		Product_description: product.Product_description,
		Review_Rating:       product.Review_Rating,
		Size:                product.Size,
	}
}

func FromDomain(domain products.ProductDomain) Product {
	return Product{
		ID:             domain.ID,
		Code:           domain.Code,
		Name:           domain.Name,
		Price:          domain.Price,
		Picture_url:    domain.Picture_url,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		Product_typeID: domain.Product_typeID,
	}
}

func ToListDomain(data []Product) (result []products.ProductDomain) {
	result = []products.ProductDomain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}
