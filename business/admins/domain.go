package admins

import (
	"context"
	"time"
)

type AdminDomain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type UseCase interface {
	Register(ctx context.Context, domain AdminDomain) (AdminDomain, error)
	Login(ctx context.Context, email, password string) (AdminDomain, error)
}
type Repository interface {
	Register(ctx context.Context, domain AdminDomain) (AdminDomain, error)
	Login(ctx context.Context, email, password string) (AdminDomain, error)
}
