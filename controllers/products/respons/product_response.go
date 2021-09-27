package respons

import (
	"final_project/business/products"
	"time"
)

type ProductResponse struct {
	ID                  int         `json:"id"`
	Code                string      `json:"code"`
	Name                string      `json:"name"`
	Total_Stock         int         `json:"total_stock"`
	Price               float64     `json:"price"`
	Picture_url         string      `json:"picture_url"`
	CreatedAt           time.Time   `json:"createdAt"`
	UpdatedAt           time.Time   `json:"updatedAt"`
	Product_typeID      int         `json:"product_typeid"`
	Product_type        interface{} `json:"product_type"`
	Product_description interface{} `json:"product_desription"`
	Review_Rating       interface{} `json:"review_rating"`
	Size                interface{} `json:"size"`
}

type Review_RatingResponse struct {
	ID        int       `json:"id"`
	Review    string    `json:"review"`
	Rating    float32   `json:"rating"`
	UserID    int       `json:"userid"`
	ProductID uint      `json:"productid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Product_descriptionResponse struct {
	ProductID   int       `json:"productid"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Product_typeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SizeResponse struct {
	ID        int       `json:"id"`
	ProductID int       `json:"productid"`
	Type      string    `json:"tipe"`
	Size      string    `json:"size"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type UploadProductResponse struct {
	ID             int       `json:"id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	Picture_url    string    `json:"picture_url"`
	Product_typeID int       `json:"product_typeid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func ProductFromDomain(domain products.ProductDomain) UploadProductResponse {
	return UploadProductResponse{
		ID:             domain.ID,
		Code:           domain.Code,
		Name:           domain.Name,
		Price:          domain.Price,
		Picture_url:    domain.Picture_url,
		Product_typeID: domain.Product_typeID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
func SizeFromDomain(domain products.SizeDomain) SizeResponse {
	return SizeResponse{
		ID:        domain.ID,
		ProductID: domain.ProductID,
		Type:      domain.Type,
		Size:      domain.Size,
		Stock:     domain.Stock,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomain(domain products.ProductDomain) ProductResponse {
	return ProductResponse{
		ID:                  domain.ID,
		Code:                domain.Code,
		Name:                domain.Name,
		Price:               domain.Price,
		Picture_url:         domain.Picture_url,
		CreatedAt:           domain.CreatedAt,
		UpdatedAt:           domain.UpdatedAt,
		Product_typeID:      domain.Product_typeID,
		Product_type:        domain.Product_type,
		Product_description: domain.Product_description,
		Review_Rating:       domain.Review_Rating,
		Size:                domain.Size,
	}
}
func TypeFromDomain(domain products.Product_typeDomain) Product_typeResponse {
	return Product_typeResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ListFromDomain(data []products.ProductDomain) (result []ProductResponse) {
	result = []ProductResponse{}
	for _, products := range data {
		result = append(result, FromDomain(products))
	}
	return
}
