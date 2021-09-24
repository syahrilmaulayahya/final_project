package users

import (
	"final_project/business/users"
	"time"
)

type User struct {
	ID int `gorm:"primaryKey" json:"id"`

	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Phone_number int
	Gender       string
	Dob          time.Time
	Picture_url  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		Phone_number: user.Phone_number,
		Gender:       user.Gender,
		Dob:          user.Dob,
		Picture_url:  user.Picture_url,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:           domain.ID,
		Name:         domain.Name,
		Email:        domain.Email,
		Password:     domain.Password,
		Phone_number: domain.Phone_number,
		Gender:       domain.Gender,
		Dob:          domain.Dob,
		Picture_url:  domain.Picture_url,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
