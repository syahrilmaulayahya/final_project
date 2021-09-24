package users

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	Name         string
	Email        string
	Password     string
	Token        string
	Phone_number int
	Gender       string
	Dob          string
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type UseCase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Get(ctx context.Context) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Get(ctx context.Context) (Domain, error)
}
