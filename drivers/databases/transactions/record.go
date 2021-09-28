package users

import (
	"final_project/business/transactions"
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
type Product struct {
	ID             int    `gorm:"primaryKey"`
	Code           string `gorm:"unique"`
	Name           string `gorm:"index"`
	Total_Stock    int
	Price          float64
	Picture_url    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Product_typeID int
}
type Size struct {
	ID        int `gorm:"primaryKey"`
	ProductID int
	Type      string
	Size      string
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Shopping_Cart struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	ProductID int
	Product   Product
	SizeID    int
	Size      Size
	Quantity  int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (shopping_cart *Shopping_Cart) ToDomain() transactions.Shopping_CartDomain {
	return transactions.Shopping_CartDomain{
		ID:        shopping_cart.ID,
		UserID:    shopping_cart.UserID,
		ProductID: shopping_cart.ProductID,
		Product:   shopping_cart.Product,
		SizeID:    shopping_cart.SizeID,
		Size:      shopping_cart.Size,
		Quantity:  shopping_cart.Quantity,
		Price:     shopping_cart.Price,
		CreatedAt: shopping_cart.CreatedAt,
		UpdatedAt: shopping_cart.UpdatedAt,
	}
}
