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
	Dob          time.Time `json:"dob"`
	Address      string    `json:"Address"`
	Picture_url  string    `json:"picture_url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
type UserNoTokenResponse struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone_number int       `json:"phone_number"`
	Gender       string    `json:"gender"`
	Dob          time.Time `json:"dob"`
	Address      string    `json:"Address"`
	Picture_url  string    `json:"picture_url"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NoTokenFromDomain(domain users.Domain) UserNoTokenResponse {
	return UserNoTokenResponse{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Phone_number: domain.Phone_number,
		Gender:       domain.Gender,
		Dob:          domain.Dob,
		Address:      domain.Address,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
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
		Address:      domain.Address,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

type Review_RatingResponse struct {
	ID        int       `json:"id"`
	Review    string    `json:"review"`
	Rating    float32   `json:"rating"`
	UserID    int       `json:"userid"`
	ProductID int       `json:"productid"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updateat"`
}

func ReviewFromDomain(domain users.Review_RatingDomain) Review_RatingResponse {
	return Review_RatingResponse{
		ID:        domain.ID,
		Review:    domain.Review,
		Rating:    domain.Rating,
		UserID:    domain.ID,
		ProductID: domain.ProductID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
