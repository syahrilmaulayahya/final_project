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
	Dob          time.Time
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
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
type UseCase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Details(ctx context.Context, id int) (Domain, error)
	UploadReview(ctx context.Context, domain Review_RatingDomain, id int) (Review_RatingDomain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email, password string) (Domain, error)
	Details(ctx context.Context, id int) (Domain, error)
	UploadReview(ctx context.Context, domain Review_RatingDomain, id int) (Review_RatingDomain, error)
}
