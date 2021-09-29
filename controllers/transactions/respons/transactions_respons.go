package respons

import (
	"final_project/business/transactions"
	"time"
)

type Shopping_CartResponse struct {
	ID        int `json:"id"`
	UserID    int `json:"userid"`
	ProductID int `json:"productid"`
	SizeID    int `json:"sizeid"`
	Quantity  int `json:"quantity"`
	Price     float64
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Payment_MethodRespons struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type ShipmentRespons struct {
	ID             int
	Name           string
	Shipment_Type  string
	Shipment_Price float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FromDomain(domain transactions.Shopping_CartDomain) Shopping_CartResponse {
	return Shopping_CartResponse{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ProductID: domain.ProductID,
		SizeID:    domain.SizeID,
		Quantity:  domain.Quantity,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func PMFromDomain(domain transactions.Payment_MethodDomain) Payment_MethodRespons {
	return Payment_MethodRespons{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func ShipmentFromDomain(domain transactions.ShipmentDomain) ShipmentRespons {
	return ShipmentRespons{
		ID:             domain.ID,
		Name:           domain.Name,
		Shipment_Type:  domain.Shipment_Type,
		Shipment_Price: domain.Shipment_Price,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
func ListFromDomain(data []transactions.Shopping_CartDomain) (result []Shopping_CartResponse) {
	result = []Shopping_CartResponse{}
	for _, shopping_cart := range data {

		result = append(result, FromDomain(shopping_cart))
	}
	return
}

func ListPMFromDomain(data []transactions.Payment_MethodDomain) (result []Payment_MethodRespons) {
	result = []Payment_MethodRespons{}
	for _, payment_method := range data {
		result = append(result, PMFromDomain(payment_method))
	}
	return
}

func ListShipmentFromDomain(data []transactions.ShipmentDomain) (result []ShipmentRespons) {
	result = []ShipmentRespons{}
	for _, shipment := range data {
		result = append(result, ShipmentFromDomain(shipment))
	}
	return
}
