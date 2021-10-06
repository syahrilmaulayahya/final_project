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
type Product_description struct {
	ProductID   int `gorm:"primaryKey, unique"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Product_type struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Product struct {
	ID                  int    `gorm:"primaryKey"`
	Code                string `gorm:"unique"`
	Name                string `gorm:"index"`
	Price               float64
	Picture_url         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Product_typeID      int
	Product_type        Product_type
	Product_description Product_description
	Size                Size
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
type Payment_Method struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Shipment struct {
	ID             int `gorm:"primaryKey"`
	Name           string
	Shipment_Type  string
	Shipment_Price float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Transaction struct {
	ID               int `gorm:"primaryKey"`
	Status           string
	UserID           int
	Shopping_CartID  int
	Total_Qty        int
	Total_Price      float64
	Payment_MethodID int
	Payment_Method   Payment_Method
	ShipmentID       int
	Shipment         Shipment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
type Transaction_Detail struct {
	UserID         int
	StatusShipment string
	TransactionID  int `gorm:"primaryKey"`
	Transaction    Transaction
	ProductID      int `gorm:"primaryKey"`
	Product        Product
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (transaction_details *Transaction_Detail) ToDomain() transactions.Transaction_DetailDomain {
	return transactions.Transaction_DetailDomain{

		UserID:         transaction_details.UserID,
		StatusShipment: transaction_details.StatusShipment,
		TransactionID:  transaction_details.TransactionID,
		ProductID:      transaction_details.ProductID,
		CreatedAt:      transaction_details.CreatedAt,
		UpdatedAt:      transaction_details.UpdatedAt,
	}
}
func (payment_method *Payment_Method) ToDomain() transactions.Payment_MethodDomain {
	return transactions.Payment_MethodDomain{
		ID:        payment_method.ID,
		Name:      payment_method.Name,
		CreatedAt: payment_method.CreatedAt,
		UpdatedAt: payment_method.UpdatedAt,
	}
}
func (shopping_cart *Shopping_Cart) ToDomain() transactions.Shopping_CartDomain {
	return transactions.Shopping_CartDomain{
		ID:        shopping_cart.ID,
		UserID:    shopping_cart.UserID,
		ProductID: shopping_cart.ProductID,
		Product:   shopping_cart.Product.ToDomain(),
		SizeID:    shopping_cart.SizeID,
		Size:      shopping_cart.Size.ToDomain(),
		Quantity:  shopping_cart.Quantity,
		Price:     shopping_cart.Price,
		CreatedAt: shopping_cart.CreatedAt,
		UpdatedAt: shopping_cart.UpdatedAt,
	}
}
func (shipment *Shipment) ToDomain() transactions.ShipmentDomain {
	return transactions.ShipmentDomain{
		ID:             shipment.ID,
		Name:           shipment.Name,
		Shipment_Type:  shipment.Shipment_Type,
		Shipment_Price: shipment.Shipment_Price,
		UpdatedAt:      shipment.UpdatedAt,
		CreatedAt:      shipment.CreatedAt,
	}
}
func (checkout *Transaction) ToDomain() transactions.TransactionDomain {
	return transactions.TransactionDomain{
		ID:               checkout.ID,
		Status:           checkout.Status,
		UserID:           checkout.UserID,
		Shopping_CartID:  checkout.Shopping_CartID,
		Total_Qty:        checkout.Total_Qty,
		Total_Price:      checkout.Total_Price,
		Payment_MethodID: checkout.Payment_MethodID,
		Payment_Method:   checkout.Payment_Method.ToDomain(),
		ShipmentID:       checkout.ShipmentID,
		Shipment:         checkout.Shipment.ToDomain(),
		CreatedAt:        checkout.CreatedAt,
		UpdatedAt:        checkout.UpdatedAt,
	}
}
func ListSCToDomain(data []Shopping_Cart) (result []transactions.Shopping_CartDomain) {
	for _, SC := range data {
		result = append(result, SC.ToDomain())
	}
	return
}

func ListPMToDomain(data []Payment_Method) (result []transactions.Payment_MethodDomain) {
	for _, PM := range data {
		result = append(result, PM.ToDomain())
	}
	return
}

func ListShipmentToDomain(data []Shipment) (result []transactions.ShipmentDomain) {
	for _, Shipment := range data {
		result = append(result, Shipment.ToDomain())
	}
	return
}

func (product *Product) ToDomain() transactions.ProductDomain {
	return transactions.ProductDomain{
		ID:                  product.ID,
		Code:                product.Code,
		Name:                product.Name,
		Price:               product.Price,
		Picture_url:         product.Picture_url,
		CreatedAt:           product.CreatedAt,
		UpdatedAt:           product.UpdatedAt,
		Product_typeID:      product.Product_typeID,
		Product_type:        product.Product_type.ToDomain(),
		Product_description: product.Product_description.ToDomain(),
		Size:                product.Size.ToDomain(),
	}
}

func (product *Product_type) ToDomain() transactions.Product_typeDomain {
	return transactions.Product_typeDomain{
		ID:        product.ID,
		Name:      product.Name,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func (Product *Size) ToDomain() transactions.SizeDomain {
	return transactions.SizeDomain{
		ID:        Product.ID,
		ProductID: Product.ProductID,
		Type:      Product.Type,
		Size:      Product.Size,
		Stock:     Product.Stock,
		CreatedAt: Product.CreatedAt,
		UpdatedAt: Product.UpdatedAt,
	}
}
func (Product *Product_description) ToDomain() transactions.Product_descriptionDomain {
	return transactions.Product_descriptionDomain{
		ProductID:   Product.ProductID,
		Description: Product.Description,
		CreatedAt:   Product.CreatedAt,
		UpdatedAt:   Product.UpdatedAt,
	}
}
