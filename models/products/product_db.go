package products

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                  uint                `json:"id"`
	Code                string              `json:"code"`
	Name                string              `json:"name"`
	Stock               uint                `json:"stock"`
	Price               float64             `json:"price"`
	Picture_url         string              `json:"picture_url"`
	CreatedAt           time.Time           `json:"createdAt"`
	UpdatedAt           time.Time           `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt      `gorm:"index" json:"deletedAt"`
	Product_type        Product_type        `json:"product_type"`
	Product_description Product_description `json:"product_desription"`
	Review_Rating       []Review_Rating     `json:"review_rating"`
}

type Review_Rating struct {
	ID        uint           `json:"id"`
	Review    string         `json:"review"`
	Rating    float32        `json:"rating"`
	ProductID uint           `json:"productid"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Product_description struct {
	Description string         `json:"description"`
	ProductID   uint           `json:"productid"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Product_type struct {
	Id        uint           `json:"id"`
	ProductID uint           `json:"productid"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
