package users

import (
	"context"
	"errors"
	"final_project/app/middleware"
	"final_project/helpers"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JwtToken       middleware.ConfigJWT
}

func NewUserUseCase(repo Repository, timeOut time.Duration, token middleware.ConfigJWT) UseCase {
	return &UserUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
		JwtToken:       token,
	}
}
func (uc *UserUseCase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("name is empty")
	}

	if !helpers.CheckEmail(domain.Email) {
		return Domain{}, errors.New("invalid email type")
	}

	if !helpers.CheckPassword(domain.Password) {
		return Domain{}, errors.New("password must contain more than 6 character, contains uppercase, lowercase and numbers")
	}
	if domain.Phone_number == 0 {
		return Domain{}, errors.New("phone number is empty")
	}
	if domain.Gender == "" {
		return Domain{}, errors.New("gender is empty")
	}
	if domain.Dob.IsZero() {
		return Domain{}, errors.New("date of birtday is empty")
	}
	user, err := uc.Repo.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email is empty")
	}

	if password == "" {
		return Domain{}, errors.New("password is empty")
	}
	user, err := uc.Repo.Login(ctx, email, password)
	var fail error
	user.Token, fail = uc.JwtToken.GenerateToken(user.ID)
	if fail != nil {
		return Domain{}, fail
	}
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) Details(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.Details(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) UploadReview(ctx context.Context, domain Review_RatingDomain, id int) (Review_RatingDomain, error) {
	if domain.Rating > 5 || domain.Rating < 0 {
		return Review_RatingDomain{}, errors.New("invalid rating")
	}
	review, err := uc.Repo.UploadReview(ctx, domain, id)
	if err != nil {
		return Review_RatingDomain{}, err
	}
	return review, nil
}
