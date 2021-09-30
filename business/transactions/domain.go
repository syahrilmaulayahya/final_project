package transactions

import (
	"context"
	"time"
)

type Shopping_CartDomain struct {
	ID        int
	UserID    int
	ProductID int
	Product   interface{}
	SizeID    int
	Size      interface{}
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
	Payment_Method   interface{}
	ShipmentID       int
	Shipment         interface{}
	CreatedAt        time.Time
	UpdatedAt        time.Time
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
}
