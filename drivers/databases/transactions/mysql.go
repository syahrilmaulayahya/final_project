package users

import (
	"context"
	"final_project/business/transactions"

	"gorm.io/gorm"
)

type MysqlTransactionRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) transactions.Repository {
	return &MysqlTransactionRepository{
		Conn: conn,
	}
}

func (rep MysqlTransactionRepository) Add(ctx context.Context, domain transactions.Shopping_CartDomain) (transactions.Shopping_CartDomain, error) {
	var shopping_cart Shopping_Cart
	var size Size
	var product Product
	productSearch := rep.Conn.Where("id = ?", domain.ProductID).First(&product)
	if productSearch.Error != nil {
		return transactions.Shopping_CartDomain{}, productSearch.Error
	}
	shopping_cart.UserID = domain.UserID
	shopping_cart.ProductID = domain.ProductID
	shopping_cart.SizeID = domain.SizeID
	shopping_cart.Quantity = domain.Quantity
	shopping_cart.Price = float64(domain.Quantity) * product.Price

	result := rep.Conn.Preload("Product").Preload("Size").Create(&shopping_cart)
	if result.Error != nil {
		return transactions.Shopping_CartDomain{}, result.Error
	}
	sizeSearch := rep.Conn.Where("id = ?", domain.SizeID).First(&size)
	if sizeSearch.Error != nil {
		return transactions.Shopping_CartDomain{}, sizeSearch.Error
	}
	stockUpdate := rep.Conn.Model(&size).Where("id = ?", domain.SizeID).Update("stock", size.Stock-domain.Quantity)
	if stockUpdate.Error != nil {
		return transactions.Shopping_CartDomain{}, stockUpdate.Error
	}

	return shopping_cart.ToDomain(), nil

}

func (rep MysqlTransactionRepository) DetailSC(ctx context.Context, id int) ([]transactions.Shopping_CartDomain, error) {
	var listSC []Shopping_Cart
	result := rep.Conn.Preload("Product").Preload("Size").Find(&listSC, "user_id = ?", id)
	if result.Error != nil {
		return []transactions.Shopping_CartDomain{}, result.Error
	}
	return ListSCToDomain(listSC), nil
}

func (rep MysqlTransactionRepository) AddPM(ctx context.Context, domain transactions.Payment_MethodDomain) (transactions.Payment_MethodDomain, error) {
	var payment_method Payment_Method
	payment_method.Name = domain.Name
	result := rep.Conn.Create(&payment_method)
	if result.Error != nil {
		return transactions.Payment_MethodDomain{}, result.Error
	}
	return payment_method.ToDomain(), nil
}

func (rep MysqlTransactionRepository) GetPM(ctx context.Context) ([]transactions.Payment_MethodDomain, error) {
	var listPayment []Payment_Method
	result := rep.Conn.Find(&listPayment)
	if result.Error != nil {
		return []transactions.Payment_MethodDomain{}, result.Error
	}
	return ListPMToDomain(listPayment), nil
}

func (rep MysqlTransactionRepository) AddShipment(ctx context.Context, domain transactions.ShipmentDomain) (transactions.ShipmentDomain, error) {
	var shipment Shipment
	shipment.Name = domain.Name
	shipment.Shipment_Type = domain.Shipment_Type
	shipment.Shipment_Price = domain.Shipment_Price
	result := rep.Conn.Create(&shipment)
	if result.Error != nil {
		return transactions.ShipmentDomain{}, result.Error
	}
	return shipment.ToDomain(), nil
}

func (rep MysqlTransactionRepository) GetShipment(ctx context.Context) ([]transactions.ShipmentDomain, error) {
	var listshipment []Shipment
	result := rep.Conn.Find(&listshipment)
	if result.Error != nil {
		return []transactions.ShipmentDomain{}, result.Error
	}
	return ListShipmentToDomain(listshipment), nil
}
