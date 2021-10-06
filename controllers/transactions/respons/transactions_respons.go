package respons

import (
	"final_project/business/transactions"
	"time"
)

type Shopping_CartResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userid"`
	ProductID int       `json:"productid"`
	SizeID    int       `json:"sizeid"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Shopping_CartDetailResponse struct {
	ID        int             `json:"id"`
	UserID    int             `json:"userid"`
	ProductID int             `json:"productid"`
	Product   ProductResponse `json:"product"`
	SizeID    int             `json:"sizeid"`
	Size      SizeResponse    `json:"size"`
	Quantity  int             `json:"quantity"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}
type Payment_MethodRespons struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type ShipmentRespons struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Shipment_Type  string    `json:"shipment_type"`
	Shipment_Price float64   `json:"shipment_price"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type TransactionRespons struct {
	ID               int                   `json:"id"`
	Status           string                `json:"status"`
	UserID           int                   `json:"userid"`
	ShoppinCartID    int                   `json:"shopping_cartID"`
	Total_Qty        int                   `json:"total_qty"`
	Total_Price      float64               `json:"total_price"`
	Payment_MethodID int                   `json:"payment_methodId"`
	Payment_Method   Payment_MethodRespons `json:"payment_method"`
	ShipmentID       int                   `json:"shipmentid"`
	Shipment         ShipmentRespons       `json:"shipment"`
	CreatedAt        time.Time             `json:"createdAt"`
	UpdatedAt        time.Time             `json:"updatedAt"`
}
type Transaction_DetailRespons struct {
	UserID         int       `json:"userid"`
	StatusShipment string    `json:"statusShipment"`
	TransactionID  int       `json:"transactionid"`
	ProductID      int       `json:"productid"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type ProductResponse struct {
	ID             int       `json:"id"`
	Code           string    `json:"code"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	Picture_url    string    `json:"picture_url"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Product_typeID int       `json:"product_typeid"`
}
type SizeResponse struct {
	ID        int       `json:"id"`
	ProductID int       `json:"productid"`
	Type      string    `json:"tipe"`
	Size      string    `json:"size"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Product_typeResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type Product_descriptionResponse struct {
	ProductID   int       `json:"productid"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func DetailFromDomain(domain transactions.Transaction_DetailDomain) Transaction_DetailRespons {
	return Transaction_DetailRespons{

		UserID:         domain.UserID,
		StatusShipment: domain.StatusShipment,
		TransactionID:  domain.TransactionID,
		ProductID:      domain.ProductID,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
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

func ShoppingCartFromDomain(domain transactions.Shopping_CartDomain) Shopping_CartDetailResponse {
	return Shopping_CartDetailResponse{
		ID:        domain.ID,
		UserID:    domain.UserID,
		ProductID: domain.ProductID,
		Product:   ProductFromDomain(domain.Product),
		SizeID:    domain.SizeID,
		Size:      SizeFromDomain(domain.Size),
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
func TransactionFromDomain(domain transactions.TransactionDomain) TransactionRespons {
	return TransactionRespons{
		ID:               domain.ID,
		Status:           domain.Status,
		UserID:           domain.UserID,
		ShoppinCartID:    domain.Shopping_CartID,
		Total_Qty:        domain.Total_Qty,
		Total_Price:      domain.Total_Price,
		Payment_MethodID: domain.Payment_MethodID,
		Payment_Method:   PMFromDomain(domain.Payment_Method),
		ShipmentID:       domain.ShipmentID,
		Shipment:         ShipmentFromDomain(domain.Shipment),
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
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
func ListSCromDomain(data []transactions.Shopping_CartDomain) (result []Shopping_CartDetailResponse) {
	result = []Shopping_CartDetailResponse{}
	for _, payment_method := range data {
		result = append(result, ShoppingCartFromDomain(payment_method))
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

func ProductFromDomain(domain transactions.ProductDomain) ProductResponse {
	return ProductResponse{
		ID:             domain.ID,
		Code:           domain.Code,
		Name:           domain.Name,
		Price:          domain.Price,
		Picture_url:    domain.Picture_url,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
		Product_typeID: domain.Product_typeID,
	}
}
func TypeFromDomain(domain transactions.Product_typeDomain) Product_typeResponse {
	return Product_typeResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
func DescriptionFromDomain(domain transactions.Product_descriptionDomain) Product_descriptionResponse {
	return Product_descriptionResponse{
		ProductID:   domain.ProductID,
		Description: domain.Description,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
func SizeFromDomain(domain transactions.SizeDomain) SizeResponse {
	return SizeResponse{
		ID:        domain.ID,
		ProductID: domain.ProductID,
		Type:      domain.Type,
		Size:      domain.Size,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
