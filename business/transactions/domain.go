package transactions

import (
	"context"
	"time"
)

type Shopping_CartDomain struct {
	ID        int
	UserID    int
	ProductID int
	Product   ProductDomain
	SizeID    int
	Size      SizeDomain
	Quantity  int
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Payment_MethodDomain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type ShipmentDomain struct {
	ID             int
	Name           string
	Shipment_Type  string
	Shipment_Price float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type TransactionDomain struct {
	ID               int
	Status           string
	UserID           int
	Shopping_CartID  int
	Total_Qty        int
	Total_Price      float64
	Payment_MethodID int
	Payment_Method   Payment_MethodDomain
	ShipmentID       int
	Shipment         ShipmentDomain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
type Transaction_DetailDomain struct {
	UserID         int
	StatusShipment string
	TransactionID  int
	ProductID      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
type Product_descriptionDomain struct {
	ProductID   int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type Product_typeDomain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SizeDomain struct {
	ID        int
	ProductID int
	Type      string
	Size      string
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type ProductDomain struct {
	ID                  int
	Code                string
	Name                string
	Price               float64
	Picture_url         string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Product_typeID      int
	Product_type        Product_typeDomain
	Product_description Product_descriptionDomain
	Size                SizeDomain
}

type UseCase interface {
	Add(ctx context.Context, domain Shopping_CartDomain) (Shopping_CartDomain, error)
	DetailSC(ctx context.Context, id int) ([]Shopping_CartDomain, error)

	AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error)
	GetPM(ctx context.Context) ([]Payment_MethodDomain, error)

	AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error)
	GetShipment(ctx context.Context) ([]ShipmentDomain, error)

	Checkout(ctx context.Context, userid, shopping_cartid int) (TransactionDomain, error)
	ChoosePnS(ctx context.Context, domain TransactionDomain) (TransactionDomain, error)
	Pay(ctx context.Context, transactionid int, amount float64) (TransactionDomain, error)

	GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, Shopping_CartDomain, error)
	Delivered(ctx context.Context, userid, transid int) (Transaction_DetailDomain, error)
	Canceled(ctx context.Context, userid, transid int) (Transaction_DetailDomain, error)
}
type Repository interface {
	Add(ctx context.Context, domain Shopping_CartDomain) (Shopping_CartDomain, error)
	DetailSC(ctx context.Context, id int) ([]Shopping_CartDomain, error)

	AddPM(ctx context.Context, domain Payment_MethodDomain) (Payment_MethodDomain, error)
	GetPM(ctx context.Context) ([]Payment_MethodDomain, error)

	AddShipment(ctx context.Context, domain ShipmentDomain) (ShipmentDomain, error)
	GetShipment(ctx context.Context) ([]ShipmentDomain, error)
	Checkout(ctx context.Context, userid, shopping_cartid int) (TransactionDomain, error)
	ChoosePnS(ctx context.Context, domain TransactionDomain) (TransactionDomain, error)
	Pay(ctx context.Context, transactionid int, amount float64) (TransactionDomain, error)

	GetTransDetail(ctx context.Context, userid, transid int) (Transaction_DetailDomain, TransactionDomain, Shopping_CartDomain, error)
	Delivered(ctx context.Context, userid, transid int) (Transaction_DetailDomain, error)
	Canceled(ctx context.Context, userid, transid int) (Transaction_DetailDomain, error)
}
