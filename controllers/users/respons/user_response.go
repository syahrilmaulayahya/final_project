package respons

import (
	"final_project/business/users"
	"time"
)

type UserResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Token        string    `json:"token"`
	Phone_number int       `json:"phone_number"`
	Gender       string    `json:"gender"`
	Dob          string    `json:"dob"`
	Picture_url  string    `json:"picture_url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Token:        domain.Token,
		Phone_number: domain.Phone_number,
		Gender:       domain.Gender,
		Dob:          domain.Dob,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
