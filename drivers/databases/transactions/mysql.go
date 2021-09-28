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
