package users

import (
	"context"
	"errors"
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

func (rep MysqlTransactionRepository) Checkout(ctx context.Context, userid, shopping_cartid int) (transactions.TransactionDomain, error) {
	var checkout Transaction
	var shopping_cart Shopping_Cart
	var shipment Shipment

	shipmentChoose := rep.Conn.First(&shipment, "id = ?", 1)
	if shipmentChoose.Error != nil {
		return transactions.TransactionDomain{}, shipmentChoose.Error
	}

	sc := rep.Conn.First(&shopping_cart, "id = ?", shopping_cartid)
	if sc.Error != nil {
		return transactions.TransactionDomain{}, sc.Error
	}
	if shopping_cart.UserID != userid {
		return transactions.TransactionDomain{}, errors.New("shopping cart not found")
	}
	checkout.Status = "unpaid"
	checkout.UserID = userid
	checkout.Shopping_CartID = shopping_cartid
	checkout.Total_Qty = shopping_cart.Quantity
	checkout.Total_Price = shopping_cart.Price + shipment.Shipment_Price
	checkout.Payment_MethodID = 1
	checkout.ShipmentID = 1
	result := rep.Conn.Create(&checkout)
	if result.Error != nil {
		return transactions.TransactionDomain{}, result.Error
	}
	return checkout.ToDomain(), nil
}

func (rep MysqlTransactionRepository) ChoosePnS(ctx context.Context, domain transactions.TransactionDomain) (transactions.TransactionDomain, error) {
	var pns Transaction
	var shipment Shipment
	var shopping_cart Shopping_Cart

	chooseShipment := rep.Conn.First(&shipment, "id = ?", domain.ShipmentID)
	if chooseShipment.Error != nil {
		return transactions.TransactionDomain{}, chooseShipment.Error
	}
	result := rep.Conn.First(&pns, "id = ?", domain.ID).Table("transactions").Where("id= ?", domain.ID).Updates(map[string]interface{}{"payment_method_id": domain.Payment_MethodID, "shipment_id": domain.ShipmentID})
	if result.Error != nil {
		return transactions.TransactionDomain{}, result.Error
	}
	chooseSC := rep.Conn.First(&shopping_cart, "id = ?", pns.Payment_MethodID)
	if chooseSC.Error != nil {
		return transactions.TransactionDomain{}, chooseSC.Error
	}
	resultFinal := rep.Conn.Preload("Shipment").Preload("Payment_Method").First(&pns, "id = ?", domain.ID).Table("transactions").Where("id= ?", domain.ID).Updates(map[string]interface{}{"total_price": shopping_cart.Price + shipment.Shipment_Price})
	if resultFinal.Error != nil {
		return transactions.TransactionDomain{}, result.Error
	}
	return pns.ToDomain(), nil
}

func (rep MysqlTransactionRepository) Pay(ctx context.Context, transactionid int, amount float64) (transactions.TransactionDomain, error) {
	var transaction Transaction
	chooseTransaction := rep.Conn.First(&transaction, "id = ?", transactionid)
	if chooseTransaction.Error != nil {
		return transactions.TransactionDomain{}, chooseTransaction.Error
	}
	if transaction.Total_Price == amount {
		updateStatus := rep.Conn.Preload("Shipment").Preload("Payment_Method").First(&transaction, "id = ?", transactionid).Table("transactions").Where("id = ?", transactionid).Updates(map[string]interface{}{"status": "paid"})
		if updateStatus.Error != nil {
			return transactions.TransactionDomain{}, updateStatus.Error
		}
	} else {
		return transactions.TransactionDomain{}, errors.New("invalid payment, please enter number same as total price")
	}
	return transaction.ToDomain(), nil

}
