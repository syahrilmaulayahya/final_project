package products

import (
	"context"
	"time"
)

type ProductDomain struct {
	ID          int
	Code        string
	Name        string
	Total_Stock int
	Price       float64
	Picture_url string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UseCase interface {
	Get(ctx context.Context) ([]ProductDomain, error)
}
type Repository interface {
	Get(ctx context.Context) ([]ProductDomain, error)
}
