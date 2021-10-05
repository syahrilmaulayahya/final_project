package product

import (
	"final_project/business/products"
	"time"
)

type Product struct {
	ID                  int    `gorm:"primaryKey"`
	Code                string `gorm:"unique"`
	Name                string `gorm:"index"`
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
	UserID    int
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product_description struct {
	ProductID   int `gorm:"primaryKey, unique"`
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

func (product *Product_type) ToDomain() products.Product_typeDomain {
	return products.Product_typeDomain{
		ID:        product.ID,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
func (product *Review_Rating) ToDomain() products.Review_RatingDomain {
	return products.Review_RatingDomain{
		ID:        product.ID,
		Review:    product.Review,
		Rating:    product.Rating,
		UserID:    product.UserID,
		ProductID: product.ProductID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
func (Product *Product_description) ToDomain() products.Product_descriptionDomain {
	return products.Product_descriptionDomain{
		ProductID:   Product.ProductID,
		Description: Product.Description,
		CreatedAt:   Product.CreatedAt,
		UpdatedAt:   Product.UpdatedAt,
	}
}
func (Product *Size) ToDomain() products.SizeDomain {
	return products.SizeDomain{
		ID:        Product.ID,
		ProductID: Product.ProductID,
		Type:      Product.Type,
		Size:      Product.Size,
		Stock:     Product.Stock,
		CreatedAt: Product.CreatedAt,
		UpdatedAt: Product.UpdatedAt,
	}
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
		Product_type:        product.Product_type.ToDomain(),
		Product_description: product.Product_description.ToDomain(),
		Review_Rating:       reviewToListDomain(product.Review_Rating),
		Size:                sizeToListDomain(product.Size),
	}
}

func ToListDomain(data []Product) (result []products.ProductDomain) {
	result = []products.ProductDomain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return
}

func sizeToListDomain(data []Size) (result []products.SizeDomain) {
	result = []products.SizeDomain{}
	for _, product := range data {
		result = append(result, product.ToDomain())
	}
	return
}

func reviewToListDomain(data []Review_Rating) (result []products.Review_RatingDomain) {
	result = []products.Review_RatingDomain{}
	for _, product := range data {
		result = append(result, product.ToDomain())
	}
	return
}
