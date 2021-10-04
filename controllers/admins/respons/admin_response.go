package respons

import (
	"final_project/business/admins"
	"time"
)

type AdminResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AdminNoTokenResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func AdminFromDomain(domain admins.AdminDomain) AdminResponse {
	return AdminResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Token:     domain.Token,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func AdminNoTokenFromDomain(domain admins.AdminDomain) AdminNoTokenResponse {
	return AdminNoTokenResponse{
		ID:       domain.ID,
		Name:     domain.Name,
		Email:    domain.Email,
		Password: domain.Password,

		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
