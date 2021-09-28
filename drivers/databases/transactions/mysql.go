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
	shopping_cart.UserID = domain.UserID
	shopping_cart.ProductID = domain.ProductID
	shopping_cart.SizeID = domain.SizeID
	shopping_cart.Quantity = domain.Quantity

	result := rep.Conn.Create(&shopping_cart)
	if result.Error != nil {
		return transactions.Shopping_CartDomain{}, result.Error
	}
	return shopping_cart.ToDomain(), nil

}
