package admins

import (
	"context"
	"errors"
	"final_project/app/middleware"
	"final_project/helpers"
	"time"
)

type AdminUseCase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JwtToken       middleware.ConfigJWT
}

func NewAdminUseCase(repo Repository, timeOut time.Duration, token middleware.ConfigJWT) UseCase {
	return &AdminUseCase{
		Repo:           repo,
		ContextTimeout: timeOut,
		JwtToken:       token,
	}
}

func (uc *AdminUseCase) Register(ctx context.Context, domain AdminDomain) (AdminDomain, error) {
	if domain.Name == "" {
		return AdminDomain{}, errors.New("name is empty")
	}
	if !helpers.CheckEmail(domain.Email) {
		return AdminDomain{}, errors.New("invalid email type")
	}
	if !helpers.CheckPassword(domain.Password) {
		return AdminDomain{}, errors.New("password must contain more than 6 character, contains uppercase, lowercase and numbers")
	}
	admin, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return AdminDomain{}, err
	}
	return admin, nil
}

func (uc *AdminUseCase) Login(ctx context.Context, email, password string) (AdminDomain, error) {
	if email == "" {
		return AdminDomain{}, errors.New("email is empty")
	}

	if password == "" {
		return AdminDomain{}, errors.New("password is empty")
	}

	admin, err := uc.Repo.Login(ctx, email, password)
	var fail error
	admin.Token, fail = uc.JwtToken.GenerateToken(admin.ID)
	if fail != nil {
		return AdminDomain{}, fail
	}
	if err != nil {
		return AdminDomain{}, err
	}
	return admin, nil
}
