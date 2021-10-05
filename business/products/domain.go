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
	Product_type        Product_typeDomain
	Product_description Product_descriptionDomain
	Review_Rating       []Review_RatingDomain
	Size                []SizeDomain
}
type Review_RatingDomain struct {
	ID        int
	Review    string
	Rating    float32
	UserID    int
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
	Details(ctx context.Context, id int) (ProductDomain, error)
	Search(ctx context.Context, words string) ([]ProductDomain, error)
	FilterByType(ctx context.Context, typeid int) ([]ProductDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)

	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)

	UploadSize(ctx context.Context, sizedomain SizeDomain) (SizeDomain, error)
	UpdateSize(ctx context.Context, sizedomain SizeDomain, id int) (SizeDomain, error)
	UpdateStock(ctx context.Context, stock, id int) (SizeDomain, error)

	UploadDescription(ctx context.Context, domain Product_descriptionDomain) (Product_descriptionDomain, error)
	UpdateDescription(ctx context.Context, domain Product_descriptionDomain, id int) (Product_descriptionDomain, error)
}
type Repository interface {
	Get(ctx context.Context) ([]ProductDomain, error)
	Details(ctx context.Context, id int) (ProductDomain, error)
	Search(ctx context.Context, words string) ([]ProductDomain, error)
	FilterByType(ctx context.Context, typeid int) ([]ProductDomain, error)
	UploadProduct(ctx context.Context, productdomain ProductDomain) (ProductDomain, error)
	UpdateProduct(ctx context.Context, domain ProductDomain, id int) (ProductDomain, error)

	UploadType(ctx context.Context, domain Product_typeDomain) (Product_typeDomain, error)

	UploadSize(ctx context.Context, sizedomain SizeDomain) (SizeDomain, error)
	UpdateSize(ctx context.Context, sizedomain SizeDomain, id int) (SizeDomain, error)
	UpdateStock(ctx context.Context, stock, id int) (SizeDomain, error)

	UploadDescription(ctx context.Context, domain Product_descriptionDomain) (Product_descriptionDomain, error)
	UpdateDescription(ctx context.Context, domain Product_descriptionDomain, id int) (Product_descriptionDomain, error)
}
