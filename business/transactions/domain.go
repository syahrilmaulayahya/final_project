package transactions

import (
	"context"
	"time"
)

type Shopping_CartDomain struct {
	ID        int
	UserID    int
	ProductID int
	SizeID    int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UseCase interface {
	Add(ctx context.Context, domain Shopping_CartDomain) (Shopping_CartDomain, error)
}
type Repository interface {
	Add(ctx context.Context, domain Shopping_CartDomain) (Shopping_CartDomain, error)
}
