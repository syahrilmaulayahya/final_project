package products

import (
	"context"
	"time"
)

type ProductDomain struct {
	ID                  int
	Code                string
	Name                string
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
	ProductID   int
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
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
	UploadSize(ctx context.Context, sizedomain SizeDomain) (SizeDomain, error)
	UpdateStock(ctx context.Context, stock, id int) (SizeDomain, error)
	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)
}
type Repository interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
	UploadSize(ctx context.Context, sizedomain SizeDomain) (SizeDomain, error)
	UpdateStock(ctx context.Context, stock, id int) (SizeDomain, error)
	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)
}
