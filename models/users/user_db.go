package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int            `gorm:"primaryKey" json:"id"`
	Username     string         `json:"username"`
	Picture_url  string         `json:"picture_url"`
	Phone_number int            `json:"phone_number"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"-"`
	Name         string         `json:"name"`
	Gender       string         `json:"gender"`
	Dob          string         `json:"dob"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Address      []Address      `json:"-"`
}

type Address struct {
	ID           int       `gorm:"primary key" json:"id"`
	UserID       int       `json:"userid"`
	Address      string    `json:"address"`
	Phone_number int       `json:"phone_number"`
	CreatedAt    time.Time `json:"createAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
	DeletedAt    time.Time `json:"DeletedAt"`
}
