package users

import (
	"context"
	"errors"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeOut time.Duration) UseCase {
	return &UserUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
	}
}
func (uc *UserUseCase) Register(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	if domain.Phone_number == 0 {
		return Domain{}, errors.New("phone number empty")
	}
	if domain.Gender == "" {
		return Domain{}, errors.New("gender empty")
	}
	if domain.Dob.IsZero() {
		return Domain{}, errors.New("date of birtday empty")
	}
	user, err := uc.Repo.Register(ctx, domain)
	// user.Name = domain.Name
	// user.Email = domain.Password
	// user.Password = domain.Password
	// user.Phone_number = domain.Phone_number
	// user.Gender = domain.Gender
	// user.Dob = domain.Dob
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUseCase) Login(ctx context.Context, email, password string) (Domain, error) {
	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if password == "" {
		return Domain{}, errors.New("password empty")
	}
	user, err := uc.Repo.Login(ctx, email, password)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *UserUseCase) Get(ctx context.Context) (Domain, error) {
	user, err := uc.Repo.Get(ctx)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
