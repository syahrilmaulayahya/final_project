package users

import (
	"time"

	"gorm.io/gorm"
)

type UserResponse struct {
	ID           int            `json:"id"`
	Username     string         `json:"username"`
	Picture_url  string         `json:"picture_url"`
	Phone_number int            `json:"phone_number"`
	Email        string         `json:"email"`
	Token        string         `json:"token"`
	Name         string         `json:"name"`
	Gender       string         `json:"gender"`
	Dob          string         `json:"dob"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
