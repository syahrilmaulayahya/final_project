package products

import (
	"context"
	"time"
)

type ProductDomain struct {
	ID                  int
	Code                string
	Name                string
	Total_Stock         int
	Price               float64
	Picture_url         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Product_typeID      int
	Product_type        interface{}
	Product_description interface{}
	Review_Rating       interface{}
	Size                interface{}
}
type Review_RatingDomain struct {
	ID        int
	Review    string
	Rating    float32
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product_descriptionDomain struct {
	ProductID   uint
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Product_typeDomain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SizeDomain struct {
	ID        int
	ProductID int
	Type      string
	Size      string
	Stock     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)
}
type Repository interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)
}
