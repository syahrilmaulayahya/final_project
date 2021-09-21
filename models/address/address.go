package address

import "time"

type Address struct {
	Id           int       `gorm:"primary key" json:"id"`
	UserID       int       `json:"userid"`
	Address      string    `json:"address"`
	Phone_number int       `json:"phone_number"`
	CreatedAt    time.Time `json:"createAt"`
	UpdatedAt    time.Time `json:"UpdatedAt"`
	DeletedAt    time.Time `json:"DeletedAt"`
}
