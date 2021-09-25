package product

import (
	"context"
	"final_project/business/products"

	"gorm.io/gorm"
)

type MysqlProductRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) products.Repository {
	return &MysqlProductRepository{
		Conn: conn,
	}
}

func (rep MysqlProductRepository) Get(ctx context.Context) ([]products.ProductDomain, error) {
	var product []Product
	result := rep.Conn.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return ToListDomain(product), nil
}
