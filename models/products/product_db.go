package products

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                  int                 `gorm:"primaryKey" json:"id"`
	Code                string              `json:"code"`
	Name                string              `json:"name"`
	Stock               int                 `json:"stock"`
	Price               float64             `json:"price"`
	Picture_url         string              `json:"picture_url"`
	CreatedAt           time.Time           `json:"createdAt"`
	UpdatedAt           time.Time           `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt      `gorm:"index" json:"deletedAt"`
	Product_typeID      int                 `json:"product_typeid"`
	Product_type        Product_type        `json:"product_type"`
	Product_description Product_description `json:"product_desription"`
	Review_Rating       []Review_Rating     `json:"review_rating"`
	Size                []Size              `json:"size"`
}

type Review_Rating struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Review    string         `json:"review"`
	Rating    float32        `json:"rating"`
	ProductID uint           `json:"productid"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Product_description struct {
	ProductID   int            `gorm:"primaryKey" json:"productid"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Product_type struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Size struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	ProductID int            `json:"productid"`
	Type      string         `json:"tipe"`
	Size      string         `json:"size"`
	Stock     uint           `json:"stock"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
